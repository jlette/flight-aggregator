package service_test

import (
	"aggregator/repository"
	"aggregator/service"

	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBuildResponse(t *testing.T) {
	

	// Jserver1
	dto1 := repository.FlightsDTO{
		BookingId:       "B1",
		Status:          "CONFIRMED",
		PassengerName:   "John Doe",
		FlightNumber:    "S1",
		DepartureAirport: "CDG",
		ArrivalAirport:   "LHR",
		DepartureTime:    time.Date(2024, 1, 2, 10, 0, 0, 0, time.UTC),
		ArrivalTime:      time.Date(2024, 1, 2, 12, 0, 0, 0, time.UTC),
		Price:            300,
		Currency:         "EUR",
		Id:               "1",
	}

	// Jserver2
	dto2 := repository.FlightsToBookDTO{
		Reference: "B2",
		Status:    "CONFIRMED",
		Traveler: repository.Traveler{
			FirstName: "Jean",
			LastName:  "Neymar",
		},
		Segments: []repository.Segments{
			{
				Flight: repository.Flight{
					Number: "S2",
					From:   "ORY",
					To:     "MAD",
					Depart: time.Date(2024, 1, 3, 9, 0, 0, 0, time.UTC),
					Arrive: time.Date(2024, 1, 3, 11, 0, 0, 0, time.UTC),
				},
			},
		},
		Total: repository.Total{
			Amount:   100,
			Currency: "EUR",
		},
		Id: "2",
	}

	resp, err := service.BuildResponse(
		[]repository.FlightsDTO{dto1},
		[]repository.FlightsToBookDTO{dto2},
		"price",
	)

	require.NoError(t, err)
	require.Len(t, resp.Flights, 2)

	// Tri par price
	require.Equal(t, 100, resp.Flights[0].Price)
	require.Equal(t, 300, resp.Flights[1].Price)

	// Vérification mapping
	require.Equal(t, "Jean", resp.Flights[0].Passenger.FirstName)
	require.Equal(t, "Neymar", resp.Flights[1].Passenger.FirstName)
}

func TestBuildResponse_Default(t *testing.T) {

	dto1 := repository.FlightsDTO{
		BookingId:       "B1",
		Status:          "CONFIRMED",
		PassengerName:   "Jean Neymar",
		FlightNumber:    "S1",
		DepartureAirport: "CDG",
		ArrivalAirport:   "LHR",
		DepartureTime:    time.Now(),
		ArrivalTime:      time.Now(),
		Price:            200,
		Currency:         "EUR",
		Id:               "1",
	}

	resp, err := service.BuildResponse(
		[]repository.FlightsDTO{dto1},
		nil,
		"",
	)

	require.NoError(t, err)
	require.Equal(t, "price", resp.Sort)
	require.Len(t, resp.Flights, 1)
}
