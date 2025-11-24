package sort_test

import (
	"aggregator/models"
	"aggregator/sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSortFlightsByPrice(t *testing.T) {
	
	flights := []models.Flight{
		{Price: 300},
		{Price: 100},
		{Price: 200},
	}

	sort.SortFlights(flights, "price")

	require.Equal(t, 100, flights[0].Price)
	require.Equal(t, 200, flights[1].Price)
	require.Equal(t, 300, flights[2].Price)
}

func TestSortFlightsByDepartureDate(t *testing.T) {
	t1 := time.Date(2024, 1, 5, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 3, 10, 0, 0, 0, time.UTC)
	t3 := time.Date(2024, 1, 7, 10, 0, 0, 0, time.UTC)

	flights := []models.Flight{
		{DepartureTime: t1},
		{DepartureTime: t2},
		{DepartureTime: t3},
	}

	sort.SortFlights(flights, "departure_date")

	require.True(t, flights[0].DepartureTime.Equal(t2))
	require.True(t, flights[1].DepartureTime.Equal(t1))
	require.True(t, flights[2].DepartureTime.Equal(t3))
}

func TestSortFlightsByTravelTime(t *testing.T) {
	f1Dep := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	f1Arr := time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC)

	f2Dep := time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC)
	f2Arr := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)

	flights := []models.Flight{
		{DepartureTime: f1Dep, ArrivalTime: f1Arr},
		{DepartureTime: f2Dep, ArrivalTime: f2Arr},
	}

	sort.SortFlights(flights, "travel_time")

	require.Equal(t, f2Arr, flights[0].ArrivalTime)
	require.Equal(t, f1Arr, flights[1].ArrivalTime)
}
