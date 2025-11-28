package handlers

import (
	"aggregator/repository"
	"aggregator/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "Home Page\n")
}
func Healt(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Serveur is good !\n")
}
// fonction obselete je laisse ici car ça peux service pour plus tard 
func Flight(w http.ResponseWriter, req *http.Request) {

    //J-Server 1
    res, err := http.Get("http://j-server1:4001/flights")
    if err != nil {
        http.Error(w, "Echec de récuparation de L'API", http.StatusInternalServerError)
        return
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        http.Error(w, "Echec de lecteur de la réponse de l'API", http.StatusInternalServerError)
        return
    }
	
    if res.StatusCode > 299 {
        http.Error(w, fmt.Sprintf("API returned status %d: %s", res.StatusCode, body), res.StatusCode)
        return
    }

    //J-Server 2
    res2, err := http.Get("http://j-server2:4002/flight_to_book")
    if err != nil {
        http.Error(w, "Echec de récuparation de L'API", http.StatusInternalServerError)
        return
    }
    defer res2.Body.Close()
	body2, err := io.ReadAll(res2.Body)

    if err != nil {
        http.Error(w, "Echec de lecteur de la réponse de l'API", http.StatusInternalServerError)
        return
    }

    if res2.StatusCode > 299 {
        http.Error(w, fmt.Sprintf("API returned status %d: %s", res.StatusCode, body), res.StatusCode)
        return
    }

	result := map[string]json.RawMessage{
        "flights": body,
        "flight_to_book":   body2,
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
   	json.NewEncoder(w).Encode(result)
}

func FlightSorted(w http.ResponseWriter, req *http.Request) {
	//Appel à j-server1
	resp1, err := http.Get("http://j-server1:4001/flights")
	if err != nil {
		http.Error(w, "error calling j-server1: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp1.Body.Close()

	body1, err := io.ReadAll(resp1.Body)
	if err != nil {
		http.Error(w, "error reading j-server1 response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//Appel à j-server2
	resp2, err := http.Get("http://j-server2:4002/flight_to_book")
	if err != nil {
		http.Error(w, "error calling j-server2: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp2.Body.Close()

	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		http.Error(w, "error reading j-server2 response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//Transformation des données brutes en DTOs
	flightsDTOs, err := repository.GetFlights(body1)
	if err != nil {
		http.Error(w, "error decoding j-server1 flights: "+err.Error(), http.StatusInternalServerError)
		return
	}

	flightsToBookDTOs, err := repository.GetFlightsToBook(body2)
	if err != nil {
		http.Error(w, "error decoding j-server2 flights: "+err.Error(), http.StatusInternalServerError)
		return
	}

	/* Récupère le critère de tri dans l’URL (price, departure_date, travel_time)
    Le service met "price" par défaut si vide*/
	sortBy := req.URL.Query().Get("sort") 

	//Construit la réponse finale (fusion + tri)
	response, err := service.BuildResponse(flightsDTOs, flightsToBookDTOs, sortBy)
	if err != nil {
		http.Error(w, "error building response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 6. Retourne la réponse final en JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
