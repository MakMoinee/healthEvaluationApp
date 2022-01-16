package main

import (
	"healtEvaluationApp/cmd/webapp/config"
	"healtEvaluationApp/cmd/webapp/routes"
	"healtEvaluationApp/internal/pkg/service"
	"healtEvaluationApp/internal/utility"
	"log"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config.Set()
	port := os.Getenv("HTTP_PORT")
	if len(port) == 0 {
		port = config.Registry.GetString("SERVER_PORT")
	}

	isProfiling := config.Registry.GetBool("SERVER_ENABLE_PROFILING")
	service := service.NewService(port)
	service.EnableProfiling(isProfiling)
	service.SetGlobalMetrics(promhttp.Handler())
	routes.Set(service)
	prometheus.MustRegister(utility.AppFailures)
	prometheus.MustRegister(utility.AppMetric)
	log.Printf("Starting the health evaluation backend in Port %s", port)
	err := service.Start()
	if err != nil {
		panic(err)
	}
}
