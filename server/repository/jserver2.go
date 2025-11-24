package repository

import (
	"aggregator/models"
	"encoding/json"
	"time"
)

/* DTO utilisé par j-server2 pour représenter une réservation de vols.
On convertis ça ensuite en models.Flight grâce à la fonction SetFlights().*/
type FlightsToBookDTO struct {
	Reference string `json:"reference"`
	Status string `json:"status"`
	Traveler Traveler `json:"traveler"`
    Segments []Segments `json:"segments"`
    Total Total `json:"total"`
    Id string `json:"id"`
}

type Traveler struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type Segments struct {
	Flight Flight `json:"flight"`
}

type Flight struct {
	Number string `json:"number"`
	From string `json:"from"`
	To string `json:"to"`
	Depart time.Time `json:"depart"`
	Arrive time.Time `json:"arrive"`
}

type Total struct {
	Amount int `json:"amount"`
	Currency string `json:"currency"`
}

/* adapte les données de j-server2 pour les rendre compatibles
avec notre modèle Flight commun au deux serveurs. */
func (f FlightsToBookDTO) SetFlights() []models.Flight {

    //Comme y'a qu'un seul segment par réservation dans j-server2, on le prend directement
	seg := f.Segments[0]
        flight := models.Flight{
            BookingId: f.Reference,
            Status:    f.Status,
            Passenger: models.Passenger{
                FirstName: f.Traveler.FirstName,
                LastName:  f.Traveler.LastName,
            },
            FlightNumber:     seg.Flight.Number,
            DepartureAirport: seg.Flight.From,
            ArrivalAirport:   seg.Flight.To,
            DepartureTime:    seg.Flight.Depart,
            ArrivalTime:      seg.Flight.Arrive,
            Price:            f.Total.Amount,
            Currency:         f.Total.Currency,
            Id:               f.Id,
        }

    return []models.Flight{flight}
}

//Convertit les données JSON venant de j-server2 en tableau de FlightsToBookDTO pour réutiliser ensuite dans la fonction SetFlights()
func GetFlightsToBook(data []byte) ([]FlightsToBookDTO, error) {
	var flightsToBookDTOs []FlightsToBookDTO
	err := json.Unmarshal(data, &flightsToBookDTOs)
	if err != nil {
		return nil, err
	}
	return flightsToBookDTOs, nil
}