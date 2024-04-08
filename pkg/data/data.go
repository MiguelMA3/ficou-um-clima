package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/MiguelMA3/pintou-um-clima/pkg/types"
)

const apiKey = "821d62f1"
const apiURL = "https://api.hgbrasil.com/weather?key=" + apiKey

// Where On Earth IDentifier / City Code
func GetWeather(woeid int) (types.WeatherResponse, error) {
	url := fmt.Sprintf("%s&woeid=%d", apiURL, woeid)
	response, err := http.Get(url)
	if err != nil {
		return types.WeatherResponse{}, fmt.Errorf("erro ao fazer a requisição: %v", err)
	}
	defer response.Body.Close()

	var weatherResp types.WeatherResponse
	if err := json.NewDecoder(response.Body).Decode(&weatherResp); err != nil {
		return types.WeatherResponse{}, fmt.Errorf("erro ao decodificar a resposta JSON: %v", err)
	}

	if err := writeToFile(weatherResp, woeid); err != nil {
		fmt.Printf("Erro ao gravar os dados da previsão do tempo em arquivo: %v\n", err)
	}

	return weatherResp, nil
}

func writeToFile(resp types.WeatherResponse, woeid int) error {
	fileName := fmt.Sprintf("%d.json", woeid)

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(resp); err != nil {
		return fmt.Errorf("erro ao codificar os dados: %v", err)
	}

	return nil
}
