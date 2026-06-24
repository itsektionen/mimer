package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/itsektionen/mimer/model"
	"golang.org/x/oauth2"
)

type AuthService struct {
	config      oauth2.Config
	userinfoURL string
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

func NewAuthService(config oauth2.Config, userinfoURL string) AuthService {
	return AuthService{
		config:      config,
		userinfoURL: userinfoURL,
	}
}

func (s *AuthService) ValidateToken(ctx context.Context, accessToken string) (*model.UserClaims, error) {
	if accessToken == "" {
		return nil, fmt.Errorf("empty access token")
	}

	req, err := http.NewRequestWithContext(ctx, "GET", s.userinfoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create userinfo request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send userinfo request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("userinfo request failed with status %d", resp.StatusCode)
	}

	var claims model.UserClaims
	if err := json.NewDecoder(resp.Body).Decode(&claims); err != nil {
		return nil, fmt.Errorf("failed to decode userinfo response: %w", err)
	}

	return &claims, nil
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

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error) {
	if refreshToken == "" {
		return nil, fmt.Errorf("empty refresh token")
	}

	t := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	ts := s.config.TokenSource(ctx, t)
	newToken, err := ts.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	return newToken, nil
}
