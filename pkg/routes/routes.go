package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MiguelMA3/pintou-um-clima/pkg/data"
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/weather/{woeid}", getWeatherHandler).Methods("GET")
}

func getWeatherHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	woeidStr := vars["woeid"]

	woeid, err := strconv.Atoi(woeidStr)
	if err != nil {
		http.Error(w, "WOEID inválido", http.StatusBadRequest)
		return
	}

	weatherResp, err := data.GetWeather(woeid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherResp)
}
