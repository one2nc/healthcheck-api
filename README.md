# Health-Check API

Hits a HTTP request to an endpoint and exporters the status code in prometheus metrics format.

## TODO
1. Instead of reading the target endpoint from env variable, add the functionality to read from json to support list of endpoints.
2. Need to make async calls for the list of target endpoints.


## How to run?
- `export INSTAHMS_ENDPOINT="https://instahms1203.instahmsdev.com/instahms/loginForm.do"`
- `go run main.go`
