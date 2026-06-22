package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"golang.org/x/oauth2"
)

type AuthService struct {
	config oauth2.Config
}

type OAuthState struct {
	Nonce   string `json:"nonce"`
	NextURL string `json:"nextUrl"`
}

func generateStateOauthCookie(nextURL string) (string, error) {
	b := make([]byte, 16)
	_, _ = rand.Read(b)

	s := OAuthState{
		Nonce:   base64.URLEncoding.EncodeToString(b),
		NextURL: nextURL,
	}

	j, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return base64.URLEncoding.EncodeToString(j), nil
}

func NewAuthService(config oauth2.Config) AuthService {
	return AuthService{
		config,
	}
}

func (s *AuthService) DecodeState(state string) (OAuthState, error) {
	decoded, err := base64.URLEncoding.DecodeString(state)
	if err != nil {
		fmt.Println(err)
		return OAuthState{}, err
	}

	var st OAuthState
	if err := json.Unmarshal(decoded, &st); err != nil {
		fmt.Println(err)
		return OAuthState{}, err
	}

	return st, nil
}

func (s *AuthService) GenerateURL(nextURL string) (string, string, error) {
	state, err := generateStateOauthCookie(nextURL)
	if err != nil {
		return "", "", err
	}

	u := s.config.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return u, state, nil
}

func (s *AuthService) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := s.config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}
