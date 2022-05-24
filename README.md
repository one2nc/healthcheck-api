# Health-Check API

Hits a HTTP request to an endpoint and exporters the status code in prometheus metrics format.

## How to run?
- `export INPUT_FILE=./endpoints.json`
- `go run main.go`

### How to run the docker container?
- `docker run -p 8091:8090 -v <path to endpoints.json>:/app/endpoints.json ghcr.io/one2nc/healthcheck-api:<image_tag>`

### Build for mac
- `GOOS=darwin GOARCH=amd64 go build -o ./health-check-api-mac`

## TODO
- [x] Handle signals like ctrl+c
- [x] Use gorilla mux instead of net/http package
- [x] Emit metric using prometheus go library ex: `status_code{service_name=<>} 200`
- [x] Read json file containing list of endpoints to be tested
- [x] Loop and make concurrent calls to test this endpoint. Wait for the results and send the results
