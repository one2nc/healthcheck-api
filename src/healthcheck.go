package healthcheck

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/projectdiscovery/gologger"
	"github.com/prometheus/client_golang/prometheus"
)

func parseJson(filePath string) []InputJson {
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		gologger.Fatal().Msg(err.Error())
	}

	inputJson := []InputJson{}
	err = json.Unmarshal(jsonFile, &inputJson)
	if err != nil {
		gologger.Fatal().Msg(err.Error())
	}
	return inputJson
}

func initializeGaugeVector() prometheus.GaugeVec {
	metric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "status_code",
			Help: "Status Code returned after hitting the target service",
		},
		[]string{
			"service_name",
		},
	)
	prometheus.MustRegister(metric)
	return *metric
}

func getStatusCode(serviceName string, endpoint string) (Service, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		gologger.Info().Msg(err.Error())
		return Service{}, err
	}

	return Service{
		Name:       serviceName,
		StatusCode: resp.StatusCode,
	}, nil
}

func addMetrics(svc Service, metric prometheus.GaugeVec) {
	serviceName := svc.Name
	statusCode := svc.StatusCode
	metric.With(prometheus.Labels{"service_name": serviceName}).Add(float64(statusCode))
}

func InitializeMetricExporter() {
	filePath := os.Getenv("INPUT_FILE")
	inputJson := parseJson(filePath)
	metric := initializeGaugeVector()

	for _, v := range inputJson {
		for _, val := range v.TargetServices {
			svc, err := getStatusCode(val.ServiceName, val.Endpoint)
			if err == nil {
				addMetrics(svc, metric)
			}
		}
	}
}
