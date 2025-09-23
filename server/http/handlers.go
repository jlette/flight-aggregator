package handlers

import (
	"fmt"
	"net/http"
)

func Healt(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w,"Serveur is good !")
}