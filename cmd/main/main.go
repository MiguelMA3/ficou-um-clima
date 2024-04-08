// main.go

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MiguelMA3/pintou-um-clima/pkg/data"
	"github.com/MiguelMA3/pintou-um-clima/pkg/routes"
	"github.com/gorilla/mux"
)

const updateInterval = 3 * time.Minute // Atualizar a cada 3 minutos

func main() {
	// Configurar as rotas da aplicação
	r := mux.NewRouter()
	routes.SetupRoutes(r)

	http.Handle("/", r)

	go func() {
		// Iniciar um loop para atualizar os dados da previsão do tempo
		for {
			// Paranaguá - Antonina - Pontal do Paraná
			woeids := []int{455895, 456682, 26795245}
			for _, woeid := range woeids {
				_, err := data.GetWeather(woeid)
				if err != nil {
					log.Printf("Erro ao obter a previsão do tempo para WOEID %d: %v\n", woeid, err)
					continue
				}

				log.Printf("Previsão do tempo para WOEID %d atualizada.\n", woeid)
			}

			// Aguardar o intervalo de atualização antes de realizar a próxima atualização
			time.Sleep(updateInterval)
		}
	}()

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
