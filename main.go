package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"log"
	"mandart/view"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r.Use(middleware.Logger)
	r.Get("/", view.LoginPage)
	r.Get("/oauth/kakao", view.KakaoLogin)
	err = http.ListenAndServe(":3001", r)
	if err != nil {
		return
	}
}
