package model

type Person struct {
	ID        string  `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	ImageURL  *string `json:"imageUrl"`
}
