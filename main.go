package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"mandalart.com/views"
	"mandalart.com/views/test"
)

func main() {
	fmt.Println("Server starts to run...")
	r := chi.NewRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//utils.InitDatabase()

	r.Use(middleware.Logger)
	r.Get("/", views.LoginPage)
	r.Get("/oauth/kakao", views.KakaoLogin)
	r.Get("/test", views.GetSheetMainCells)
	r.Get("/sheet", test.CreateSheet)

	err = http.ListenAndServe(":3001", r)
	if err != nil {
		return
	}
}
