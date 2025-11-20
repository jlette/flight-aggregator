package repository

import (
	"aggregator/models"
	"encoding/json"
	"time"
)

type FlightsToBookDTO struct {
	Reference int `json:"reference"`
	Status string `json:"status"`
	Travelers Travelers `json:"travelers"`
    Segments []Segments `json:"segments"`
    Total []Total `json:"total"`
    Id int `json:"id"`
}

type Travelers struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type Segments struct {
	Fligtht []Flight `json:"flight"`
}

type Flight struct {
	Number int `json:"number"`
	From string `json:"segments"`
	To string `json:"to"`
	Depart time.Time `json:"depart"`
	Arrive time.Time `json:"arrive"`
}

type Total struct {
	Amount int `json:"amount"`
	Currency string `json:"currency"`
}

func (flightsToBookDTO FlightsToBookDTO) SetFlights() ([]models.Flight) {

	flight:= models.Flight{
		BookingId:       flightsToBookDTO.Reference,
		Status:          flightsToBookDTO.Status,
		Passenger: models.Passenger{
			
				FirstName: flightsToBookDTO.Travelers.FirstName,
				LastName:  flightsToBookDTO.Travelers.LastName,
				
		},
		FlightNumber:    "",
		DepartureAirport: "",
	}
	return []models.Flight{flight}
}	

func GetFlightsToBook(data []byte) ([]FlightsToBookDTO, error) {
	var flightsToBookDTO []FlightsToBookDTO
	err := json.Unmarshal(data, &flightsToBookDTO);
	if err != nil {
		return nil, err
	}
	return flightsToBookDTO, nil
}