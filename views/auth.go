package views

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"mandalart.com/services"
	"mandalart.com/types"
	"mandalart.com/utils/errors"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	err := t.Execute(w, nil)
	errors.Catch(err)
}

func KakaoLogin(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if len(strings.TrimSpace(code)) == 0 {
		render.Status(r, http.StatusBadRequest)	
		render.PlainText(w, r, "")
		return
	}
	authService, err := services.NewAuthService(r.Context())
	if err != nil {
		render.Status(r,http.StatusUnauthorized)
		render.PlainText(w, r, err.Error())
		return
	}
	token, err := authService.HandleSocialLogin(code, types.KAKAO)
	if err != nil {
		render.Status(r,http.StatusUnauthorized)
		render.PlainText(w, r, "")
		return
	}
	
	render.JSON(w,r,token)
}