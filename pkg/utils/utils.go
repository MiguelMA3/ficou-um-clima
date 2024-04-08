package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MiguelMA3/pintou-um-clima/pkg/types"
)

func ReadWeather(woeid int) ([]types.WeatherResponse, error) {
	// Formatar o nome do arquivo com o WOEID
	fileName := fmt.Sprintf("%d.json", woeid)
	var weatherResponses []types.WeatherResponse

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&weatherResponses); err != nil {
		return nil, fmt.Errorf("erro ao decodificar os dados do arquivo: %v", err)
	}

	return weatherResponses, nil
}
