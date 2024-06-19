package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	repo "mandalart.com/repositories"
	"mandalart.com/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	todoService := services.NewTodoService(queries, cellService)
	sheetView := views.NewSheetView(sheetService, cellService)
	authView := views.NewAuthView(authService)
	cellView := views.NewCellView(cellService, todoService)

	r.Get("/", authView.LoginPage)
	r.Get("/oauth/kakao", authView.KakaoLogin)
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(AuthCtx)
		r.Get("/sheet", sheetView.GetLatestSheetWithMainCells)
		r.Route("/cell/{cellID}", func(r chi.Router) {
			r.Use(CellCtx)
			r.Patch("/", cellView.UpdateCell)
			r.Get("/children", sheetView.GetChildrenCells)
		})
		r.Post("/sheet", sheetView.CreateSheet)
	})

	err = http.ListenAndServe(":3001", r)
	if err != nil {
		return
	}

	defer utils.DBPool.Close()
}

func CellCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			cellID, err := strconv.Atoi(chi.URLParam(r, "cellID"))
			if err != nil {
				views.Respond(w,r,http.StatusBadRequest, err)
			}

			ctx := context.WithValue(r.Context(), "cellID", cellID)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
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
			views.Respond(w,r,http.StatusUnauthorized, err)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}