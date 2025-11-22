package models

type Response struct {
	Sort    string   `json:"sort"`
	Flights []Flight `json:"flights"`
}