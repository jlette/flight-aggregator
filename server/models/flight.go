package models

import "time"

//Modèle final des vols apres traitement des données brutes des deux serveurs
type Flight struct {

	BookingId string `json:"bookingId"`
	Status string `json:"status"`
    Passenger Passenger `json:"passengerName"`
    FlightNumber string `json:"flightNumber"`
    DepartureAirport string `json:"departureAirport"`
    ArrivalAirport string `json:"arrivalAirport"`
    DepartureTime time.Time `json:"departureTime"`
    ArrivalTime time.Time `json:"arrivalTime"`
    Price int `json:"price"`
    Currency string `json:"currency"`
    Id string `json:"id"`

}

type Passenger struct {
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
}
