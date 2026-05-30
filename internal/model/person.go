package model

type Person struct {
	ID        string  `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	FirstName string  `json:"firstName" example:"Ture"`
	LastName  string  `json:"lastName" example:"Teknolog"`
	ImageURL  *string `json:"imageUrl"`
}
