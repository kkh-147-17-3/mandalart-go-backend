package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"mandalart.com/repositories"
	"mandalart.com/types"
	"mandalart.com/utils"
)

type UnauthorizedError struct {
	error
}

type AuthService struct {
	queries *repositories.Queries
	ctx    *context.Context
}

func NewAuthService(ctx context.Context) (*AuthService, error){
	conn, ok := ctx.Value("db").(*pgxpool.Pool)
	if !ok {
		return nil, fmt.Errorf("database is not initialized")
	}
	return &AuthService{repositories.New(conn), &ctx}, nil
}

func (s *AuthService) HandleSocialLogin(code string, provider types.SocialProvider) (*utils.Tokens, error) {
	var user repositories.User
	
	accessToken, err := getKakaoToken(code)
	if err != nil {
		return nil, UnauthorizedError{err}
	}

	userInfo, err := getKakaoUserInfo(accessToken)

	if err != nil {
		return nil, err
	}
	userId := strconv.Itoa(int(userInfo["id"].(float64)))

	args := repositories.GetUserBySocialProviderInfoParams{
		SocialID: pgtype.Text{String: userId, Valid: true},
		SocialProvider: pgtype.Text{String: "KAKAO", Valid: true},
	}
	user, err = s.queries.GetUserBySocialProviderInfo(*s.ctx, args)
	
	if errors.Is(err, sql.ErrNoRows) {
		userID, err := s.queries.CreateUser(*s.ctx, repositories.CreateUserParams(args))
		if err != nil {
			return nil, err
		}
		return utils.CreateToken(int(userID))
	}
	return utils.CreateToken(int(user.ID))
}

func getKakaoToken(code string) (string, error) {
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
		return "", err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	accessToken, ok := data["access_token"]
	if !ok || err != nil {
		return "", errors.New("failed to fetch access token")
	}

	return accessToken.(string), nil
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
	json.NewDecoder(resp.Body).Decode(&data)

	return data, nil
}
