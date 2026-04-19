package model

type ApiKey struct {
	ID     string `json:"id"`
	Value  string `json:"value"`
	Active bool   `json:"active"`
}
