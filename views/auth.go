package views

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"mandalart.com/services"
	"mandalart.com/utils/errors"
)

type AuthView struct {
	authService *services.AuthService
}

func NewAuthView(authService *services.AuthService) *AuthView {
	return &AuthView{authService}
}

func (v *AuthView) LoginPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	err := t.Execute(w, nil)
	errors.Catch(err)
}

func (v *AuthView) KakaoLogin(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if len(strings.TrimSpace(code)) == 0 {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, "")
		return
	}

	token, err := v.authService.HandleSocialLogin(r.Context(), code)
	if err != nil {
		render.Status(r, http.StatusUnauthorized)
		render.PlainText(w, r, "")
		return
	}

	render.JSON(w, r, token)
}
