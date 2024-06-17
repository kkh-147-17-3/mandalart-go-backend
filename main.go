package main

import (
	"context"
	"fmt"
	"log"
	repo "mandalart.com/repositories"
	"mandalart.com/services"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"mandalart.com/utils"
	"mandalart.com/views"
)

func main() {
	fmt.Println("Server starts to run...")
	r := chi.NewRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	utils.InitDatabase(os.Getenv("DATABASE_URL"))

	queries := repo.New(utils.DBPool)
	sheetService := services.NewSheetService(queries)
	cellService := services.NewCellService(queries)
	authService := services.NewAuthService(queries)
	sheetView := views.NewSheetView(sheetService, cellService)
	authView := views.NewAuthView(authService)

	r.Get("/", authView.LoginPage)
	r.Get("/oauth/kakao", authView.KakaoLogin)
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(AuthCtx)
		r.Get("/sheet", sheetView.GetLatestSheetWithMainCells)
		r.Get("/cell/{cellID}/children", sheetView.GetChildrenCells)
		r.Post("/sheet", sheetView.CreateSheet)
	})

	err = http.ListenAndServe(":3001", r)
	if err != nil {
		return
	}

	defer utils.DBPool.Close()
}

func AuthCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			err    error
			userID int
		)
		strs := strings.Split(strings.TrimSpace(r.Header.Get("Authorization")), "Bearer")
		if len(strs) < 2 {
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

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
