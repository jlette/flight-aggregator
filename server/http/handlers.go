package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Healt(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Serveur is good !\n")
}

func Flight(w http.ResponseWriter, req *http.Request) {

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

    if res.StatusCode > 299 {
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
