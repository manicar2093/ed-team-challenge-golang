package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manicar2093/ed-team-go/config"
	"github.com/manicar2093/ed-team-go/controllers"
	"github.com/manicar2093/ed-team-go/services"
)

func main() {

	config.InitEnv()

	server := mux.NewRouter()

	nomicsService := services.NewNomicsServiceWClient(config.NominicsAPI)
	chartService := services.ChartServiceImpl{}

	controller := controllers.NewChartController(&chartService, nomicsService)

	server.HandleFunc("/generate_chart", controller.CreateChartHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(config.Port, server))

}
