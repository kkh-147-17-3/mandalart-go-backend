package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"mandalart.com/utils"
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
	utils.InitDatabase(os.Getenv("DATABASE_URL"))

	r.Get("/", views.LoginPage)
	r.With(DbCtx).Get("/oauth/kakao", views.KakaoLogin)
	r.Route("/", func(r chi.Router){
		r.Use(middleware.Logger)
		r.Use(DbCtx)
		r.Use(AuthCtx)
		r.Get("/sheet", test.CreateSheet)
	})

	err = http.ListenAndServe(":3001", r)
	if err != nil {
		return
	}

	defer utils.DBPool.Close()
}

type ContextKey string

func DbCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn := utils.DBPool
		ctx := context.WithValue(r.Context(), "db", conn)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var (
			err error
			userID string
		)
		strs := strings.Split(strings.TrimSpace(r.Header.Get("Authorization")), "Bearer")
		if(len(strs) < 2){
			err = fmt.Errorf("bearer token is required")
		} else {
			tokenStr := strings.TrimSpace(strs[1])
			userID, err = utils.GetUserIdFromToken(tokenStr)
		}
	
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, err.Error())
			return 
		}

		ctx := context.WithValue(r.Context(), ContextKey("userID"), userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
