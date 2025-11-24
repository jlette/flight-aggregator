package models

// Structure de la réponse retournée par notre API.
type Response struct {
	Sort    string   `json:"sort"`
	Flights []Flight `json:"flights"`
}