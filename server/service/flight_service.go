package service

import (
	"aggregator/models"
	"aggregator/repository"
	"aggregator/sort"
)

// FlightService représente le service chargé de construire la réponse contenant les vols
type FlightService struct {}

// NewFlightService retourne une nouvelle instance du service
func NewFlightService() *FlightService {
	return &FlightService{}
}

/*BuildResponse construit la réponse finale en fusionnant les données qui viennent
des deux serveurs (j-server1 et j-server2), puis en triant les vols.*/

func (s *FlightService) BuildResponse(
	flights1DTO []repository.FlightsDTO, // Vols récupérés depuis j-server1
	flights2DTO []repository.FlightsToBookDTO, // Vols récupérés depuis j-server2
	sortBy string, // Critère de tri demandé
) (models.Response, error) {

	var flights []models.Flight

	// Convertit chaque DTO de j-server1 en vols et les ajoute à la liste
	for _, dto := range flights1DTO {
		flights = append(flights, dto.SetFlights()...)
	}
	// Convertit chaque DTO de j-server2 en vols et les ajoute également
	for _, dto := range flights2DTO {
		flights = append(flights, dto.SetFlights()...)
	}
	// Si aucun critère n’est fourni, on trie par prix par défaut
	if sortBy == "" {
		sortBy = "price"
	}


	// Trie la liste complète des vols selon le critère choisi
	sort.SortFlights(flights, sortBy)

	// Construit et retourne la réponse finale
	return models.Response{
		Sort:    sortBy,
		Flights: flights,
	}, nil
}

