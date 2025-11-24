package sort

import (
	"aggregator/models"
	"sort"
)

// SortFlights trie la liste des vols selon le choix du filtre.
// Le tri se fait directement sur le slice (les slices sont modifiés en place).
func SortFlights(flights []models.Flight, choice string) {
	switch choice {

		// Tri par date de départ : du plus tôt au plus tard
	case "departure_date":
		sort.Slice(flights, func(i, j int) bool {
			return flights[i].DepartureTime.Before(flights[j].DepartureTime)
		})
		// Tri par durée de voyage : du plus court au plus long
	case "travel_time":
		sort.Slice(flights, func(i, j int) bool {
			di := flights[i].ArrivalTime.Sub(flights[i].DepartureTime)
			dj := flights[j].ArrivalTime.Sub(flights[j].DepartureTime)
			return di < dj
		})
		// Tri par prix : du moins cher au plus cher
		// Par défaut, si le critère est inconnu ou non spécifié
	case "price":
		fallthrough
	default:
		sort.Slice(flights, func(i, j int) bool {
			return flights[i].Price < flights[j].Price
		})
	}
}
