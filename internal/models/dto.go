package models

type Input struct {
	ID []int `json:"id" binding:"required"`
}
type Person struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Born       string `json:"born"`
}
