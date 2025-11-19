package repository

import (
	"aggregator/models"
	"time"
)

type FlightsToBook struct {
	Reference int `json: "reference"`
	Status string `json: "status"`
	Travelers []Travelers `json: "travelers"`
    Segments []Segments `json: "segments"`
    Total []Total `json: "total"`
    Id int `json: "id"`
}

type Travelers struct {
	FirstName string `json: "firstName"`
	LastName string `json: "lastName"`
}

type Segments struct {
	Fligtht []Flight `json: "flight"`
}

type Flight struct {
	Number int `json: "number"`
	From string `json: "segments"`
	To string `json: "to"`
	Depart time.Time `json: "depart"`
	Arrive time.Time `json: "arrive"`
}

type Total struct {
	Amount int `json: "amount"`
	Currency string `json: "currency"`
}

func (f FlightsToBook) TransformFlights() ([]models.Flight) {
    
}