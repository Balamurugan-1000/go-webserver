package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found")
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		ExposedHeaders:   []string{"Link"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("Server spinning on port : %s", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
