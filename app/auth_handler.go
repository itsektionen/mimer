package app

import (
	"log"
	"net/http"
	"time"

	"github.com/itsektionen/mimer/app/ctxs"
	"github.com/itsektionen/mimer/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return AuthHandler{
		authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	nextURL := r.URL.Query().Get("next")
	if nextURL == "" {
		nextURL = "/admin"
	}

	url, state, err := h.authService.GenerateURL(nextURL)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   300,
	})

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *AuthHandler) HandleAuthentikCallback(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	state := r.FormValue("state")

	s, err := h.authService.DecodeState(state)
	if err != nil {
		// Do nothing?
	}

	cookie, err := r.Cookie("oauth_state")
	if err != nil || cookie.Value != state {
		log.Println(err)
		http.Error(w, "state mismatch / csrf alert", http.StatusBadRequest)
		return
	}

	token, err := h.authService.ExchangeCode(r.Context(), code)
	if err != nil {
		log.Println(err)
		http.Error(w, "exchange failed", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "mimer_token",
		Value:    token.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(time.Until(token.Expiry).Seconds()),
	})

	http.Redirect(w, r, s.NextURL, http.StatusTemporaryRedirect)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	_, ok := ctxs.UserFromContext(r.Context())

	if !ok {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "mimer_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
