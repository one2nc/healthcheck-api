package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Service struct {
	name       string
	statusCode int
}

type InputJson struct {
	TargetServices []struct {
		ServiceName string `json:"service_name"`
		Endpoint    string `json:"endpoint"`
	} `json:"target_services"`
}

func gaugeVectorInit() prometheus.GaugeVec {
	metric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "status_code",
			Help: "Status Code returned after hitting the target service",
		},
		[]string{
			"service_name",
		},
	)
	log.Printf("[INFO] initializing & registering a new gauge vector")
	prometheus.MustRegister(metric)
	return *metric
}

func addMetrics(svc Service, metric prometheus.GaugeVec) {
	serviceName := svc.name
	statusCode := svc.statusCode
	log.Printf("[INFO] adding metric to gauge vector")
	metric.With(prometheus.Labels{"service_name": serviceName}).Add(float64(statusCode))
}

func parseJson(filePath string) []InputJson {
	jsonFile, err := os.ReadFile(filePath)
	log.Print("[INFO] reading endpoints from a file")
	if err != nil {
		log.Fatal(err)
	}

	inputJson := []InputJson{}
	err = json.Unmarshal(jsonFile, &inputJson)
	log.Print("[INFO] Unmarshalling json")
	if err != nil {
		log.Fatal(err)
	}
	return inputJson
}

func getStatusCode(serviceName string, endpoint string) Service {
	resp, err := http.Get(endpoint)
	log.Printf("[INFO] making a get request to %v", endpoint)
	if err != nil {
		log.Fatal(err)
	}

	return Service{
		name:       serviceName,
		statusCode: resp.StatusCode,
	}
}

func metricExporter() {
	log.Print("[INFO] getting status code")
	filePath := os.Getenv("INPUT_FILE")
	inputJson := parseJson(filePath)
	metric := gaugeVectorInit()
	for _, v := range inputJson {
		for _, val := range v.TargetServices {
			go addMetrics(getStatusCode(val.ServiceName, val.Endpoint), metric)
		}
	}
}

func main() {
	metricExporter()
	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	log.Print("[INFO] server started at port 8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
