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

	nomicsService := services.NewNomicsServiceWClient()
	chartService := services.ChartService{}

	controller := controllers.NewChartController(&chartService, nomicsService)

	server.HandleFunc("/generate_chart", controller.CreateChartHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(config.Port, server))

}
