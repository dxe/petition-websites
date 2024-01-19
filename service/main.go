package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/dxe/helptheducks.com/service/config"
	"github.com/dxe/helptheducks.com/service/mailer"
	"github.com/dxe/helptheducks.com/service/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strings"
	"time"
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
		AllowedOrigins:   []string{"http://localhost:5173", "https://helptheducks.com"},
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

	go worker()

	fmt.Printf("Listening on port %v\n", config.Port)
	http.ListenAndServe("localhost:"+config.Port, r)
}

func worker() {
	var err error
	mailClient, err = mailer.CreateClient()
	if err != nil {
		fmt.Printf("Could not create mail client: %v\n", err)
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
		fmt.Printf("Error getting messages to process: %v\n", err)
		return
	}

	var success, fail []int

	for _, message := range messages {
		fmt.Printf("Processing message id: %v\n", message.ID)
		fromEmail := strings.Join(strings.Split(strings.ToLower(message.Name), " "), ".") + "@mail.helptheducks.com"
		err := mailer.Send(mailClient, mailer.SendOptions{
			From:    fmt.Sprintf("%s <%s>", message.Name, fromEmail),
			ReplyTo: message.Email,
			// TODO: change this after confirming everything is working okay.
			To:      "jake@directactioneverywhere.com",
			Subject: "Prosecute Reichardt Duck Farm for Animal Abuse",
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
