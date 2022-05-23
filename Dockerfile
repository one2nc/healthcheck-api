FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY ./test_endpoints.json ./
COPY *.go ./

RUN go mod download && \
    go build -o ./health-check-api
    
EXPOSE 8090

ENV INPUT_FILE="/app/endpoints.json"

ENTRYPOINT [ "/app/health-check-api" ]
