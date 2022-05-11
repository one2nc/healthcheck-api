package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func getStatusCode() string {
	instahmsEndpoint := os.Getenv("INSTAHMS_ENDPOINT")
	resp, err := http.Get(instahmsEndpoint)
	log.Printf("[INFO] making a get request to %v", instahmsEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	return "status_code " + strconv.Itoa(resp.StatusCode) + "\n"
}

func metricExporter(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/metrics" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Print("[INFO] getting status code")
	resp := getStatusCode()
	fmt.Fprintf(w, resp)
}

func main() {

	fmt.Println("Starting health-check server at port 8090")
	http.HandleFunc("/metrics", metricExporter)

	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
