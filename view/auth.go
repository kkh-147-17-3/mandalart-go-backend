package view

import (
	"encoding/json"
	"html/template"
	"mandart/utils"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type KAuthTokenReqBody struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUri  string `json:"redirect_uri"`
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	//requestUrl := fmt.Sprintf("")
	//req, err := http.NewRequest(http.MethodPost, requestUrl, nil)

	t, _ := template.ParseFiles("templates/index.html")
	err := t.Execute(w, nil)
	utils.Catch(err)
}

func KakaoLogin(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if len(strings.TrimSpace(code)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	var data map[string]interface{}

	reqBody := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {os.Getenv("KAKAO_CLIENT_ID")},
		"code":          {code},
		"redirect_uri":  {"http://localhost:3001/oauth/kakao"},
		"client_secret": {os.Getenv("KAKAO_CLIENT_SECRET")},
	}

	requestUrl := "https://kauth.kakao.com/oauth/token"
	resp, err := http.Post(requestUrl, "application/x-www-form-urlencoded", strings.NewReader(reqBody.Encode()))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	userInfo, err := getKakaoUserInfo(data["access_token"].(string))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	result, err := json.Marshal(userInfo)
	w.Write(result)

}

func getKakaoUserInfo(token string) (map[string]interface{}, error) {
	baseUrl := "https://kapi.kakao.com/v2/user/me"
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)

	return data, nil
}
