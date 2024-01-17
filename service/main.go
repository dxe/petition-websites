package main

import (
	"encoding/json"
	"fmt"
	"github.com/dxe/helptheducks.com/service/config"
	"github.com/dxe/helptheducks.com/service/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://helptheducks.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers.
	}))

	db := mustGetOrCreateDb()
	defer db.Close()

	r.Route("/message", func(r chi.Router) {
		r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
			// TODO: add captcha.
			var message model.Message
			err := json.NewDecoder(r.Body).Decode(&message)
			if err != nil {
				http.Error(w, fmt.Sprintf("error parsing request body: %q", err), http.StatusBadRequest)
				return
			}
			model.InsertMessage(r.Context(), db, message)
			w.Write([]byte("ok"))
		})
	})

	// TODO: add a worker for sending the messages via SES.

	fmt.Printf("Listening on port %v", config.Port)
	http.ListenAndServe("localhost:"+config.Port, r)
}
