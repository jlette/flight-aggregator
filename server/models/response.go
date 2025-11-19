package models

type Response struct {
	Sort string `json: "message"`
	Flights []Flight `json: "flights"`
}