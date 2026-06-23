package model

type UserClaims struct {
	Subject           string   `json:"sub"`
	Email             string   `json:"email"`
	PreferredUsername string   `json:"preferred_username"`
	Groups            []string `json:"groups"`
	Expiry            int64    `json:"exp"`
}
