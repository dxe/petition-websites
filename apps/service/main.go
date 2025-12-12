package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/dxe/service/config"
	"github.com/dxe/service/data"
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
	mailClient *ses.SES
)

func main() {
	db := getDb()
	defer db.Close()
	s := NewServer(db)

	go worker(db)
	s.runServer()
}

type server struct {
	db *sqlx.DB
}

func NewServer(db *sqlx.DB) *server {
	return &server{db: db}
}

func (s *server) runServer() {
	r = chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{
			"http://localhost:5173", 
			"http://localhost:3000", 
			"http://localhost:3002",
			"http://localhost:3003",
			"https://www.helptheducks.com", 
			"https://www.helpthechickens.com", 
			"https://righttorescue.com", 
			"https://www.freezoe.org", 
			"https://factoryfarmwatch.org",
			"https://helptheducks.dxe.io",
			},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers.
	}))

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/message", func(r chi.Router) {
		r.Post("/create", s.createMessageHandler)
	})

	r.Get("/tally", s.getTallyHandler)

	r.Get("/assemblyMembers", s.getAssemblyMembersHandler)

	fmt.Printf("Listening on port %v\n", config.Port)
	http.ListenAndServe(":"+config.Port, r)
}

func worker(db *sqlx.DB) {
	var err error
	mailClient, err = mailer.CreateClient()
	if err != nil {
		fmt.Printf("Could not create mail client: %v\n", err)
		return
	}
	for {
		processNewMessages(db)
		time.Sleep(60 * time.Second)
	}
}

func processNewMessages(db *sqlx.DB) {
	messages, err := model.GetMessagesToProcess(db)
	if err != nil {
		fmt.Printf("Error getting messages to process: %v\n", err)
		return
	}

	var success, fail []int

	for _, message := range messages {
		fmt.Printf("Processing message id: %v\n", message.ID)

		campaignName := message.Campaign.String
		settings, ok := config.CampaignEmailSettings[campaignName]
		if !ok {
			testCampaign := strings.TrimPrefix(campaignName, "test:")
			if testCampaign != campaignName {
				settings, ok = config.CampaignEmailSettings[testCampaign]
				if ok {
					settings.To = config.StaticRecipientList("tech@directactioneverywhere.com")
				}
			}
		}
		if !ok {
			settings = config.CampaignEmailSettings["test"]
		}

		normalizedName, err := removeAccents(message.Name)
		if err != nil {
			fmt.Printf("Error normalizing name: %v\n", err)
			fail = append(fail, message.ID)
			continue
		}
		fromEmail := strings.Join(strings.Split(strings.ToLower(strings.Trim(normalizedName, " ")), " "), ".") + "@" + settings.FromDomain
		err = mailer.Send(mailClient, mailer.SendOptions{
			From:    fmt.Sprintf("%s <%s>", message.Name, fromEmail),
			ReplyTo: message.Email,
			To:      settings.To(data.Municipality(message.City.String), data.Zip(message.Zip.String)),
			Subject: settings.Subject,
			Body:    message.Message,
		})
		if err != nil {
			fmt.Printf("Error sending email: %v", err)
			fail = append(fail, message.ID)
		} else {
			success = append(success, message.ID)
		}
	}

	err = model.UpdateMessageStatus(db, success, "SENT")
	if err != nil {
		fmt.Printf("Error updating message status: %v\n", err)
	}

	err = model.UpdateMessageStatus(db, fail, "FAILED")
	if err != nil {
		fmt.Printf("Error updating message status: %v\n", err)
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
