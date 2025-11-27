package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/dxe/service/config"
	"github.com/dxe/service/mailer"
	"github.com/dxe/service/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	r          *chi.Mux
	db         *sqlx.DB
	mailClient *ses.SES
)

func main() {
	r = chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://www.helptheducks.com", "https://www.helpthechickens.com", "https://righttorescue.com", "https://www.freezoe.org"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers.
	}))

	db = getDb()
	defer db.Close()

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/message", func(r chi.Router) {
		r.Post("/create", createMessageHandler)
	})

	r.Get("/tally", getTallyHandler)

	go worker()

	slog.Info("Listening on port", "port", config.Port)
	http.ListenAndServe(":"+config.Port, r)
}

func worker() {
	var err error
	mailClient, err = mailer.CreateClient()
	if err != nil {
		slog.Error("Could not create mail client", "error", err)
		return
	}
	for {
		processNewMessages()
		time.Sleep(60 * time.Second)
	}
}

func processNewMessages() {
	messages, err := model.GetMessagesToProcess(db)
	if err != nil {
		slog.Error("Error getting messages to process", "error", err)
		return
	}

	var success, fail []int

	for _, message := range messages {
		slog.Info("Processing message", "id", message.ID)

		campaignName := message.Campaign.String
		settings, ok := config.EmailSettings[campaignName]
		if !ok {
			testCampaign := strings.TrimPrefix(campaignName, "test:")
			if testCampaign != campaignName {
				settings, ok = config.EmailSettings[testCampaign]
				if ok {
					settings.To = []string{"tech@directactioneverywhere.com"}
				}
			}
		}
		if !ok {
			settings = config.EmailSettings["test"]
		}

		normalizedName, err := removeAccents(message.Name)
		if err != nil {
			slog.Error("Error normalizing name", "error", err)
			fail = append(fail, message.ID)
			continue
		}
		fromEmail := strings.Join(strings.Split(strings.ToLower(strings.Trim(normalizedName, " ")), " "), ".") + "@" + settings.FromDomain
		err = mailer.Send(mailClient, mailer.SendOptions{
			From:    fmt.Sprintf("%s <%s>", message.Name, fromEmail),
			ReplyTo: message.Email,
			To:      settings.To,
			Subject: settings.Subject,
			Body:    message.Message,
		})
		if err != nil {
			slog.Error("Error sending email", "error", err)
			fail = append(fail, message.ID)
		} else {
			success = append(success, message.ID)
		}
	}

	err = model.UpdateMessageStatus(db, success, "SENT")
	if err != nil {
		slog.Error("Error updating message status", "error", err)
	}

	err = model.UpdateMessageStatus(db, fail, "FAILED")
	if err != nil {
		slog.Error("Error updating message status", "error", err)
	}
}

func removeAccents(s string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, err := transform.String(t, s)
	if err != nil {
		return "", err
	}
	return output, nil
}
