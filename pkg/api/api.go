package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MiguelMA3/pintou-um-clima/pkg/utils"
)

// WeatherAPIHandler é o manipulador HTTP para a API de previsão do tempo.
func WeatherAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Definir o cabeçalho Content-Type para JSON
	w.Header().Set("Content-Type", "application/json")

	// Obter o parâmetro WOEID da query string da solicitação HTTP
	woeidStr := r.URL.Query().Get("woeid")
	woeid, err := strconv.Atoi(woeidStr)
	if err != nil {
		http.Error(w, "Parâmetro WOEID inválido", http.StatusBadRequest)
		return
	}

	// Obter os dados da previsão do tempo para o WOEID fornecido
	weatherResponse, err := utils.ReadWeather(woeid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter dados da previsão do tempo para WOEID %d", woeid), http.StatusInternalServerError)
		return
	}

	// Codificar os dados da previsão do tempo como JSON e enviá-los como resposta
	if err := json.NewEncoder(w).Encode(weatherResponse); err != nil {
		http.Error(w, "Erro ao codificar dados da previsão do tempo como JSON", http.StatusInternalServerError)
		return
	}
}
