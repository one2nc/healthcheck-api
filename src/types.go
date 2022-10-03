package healthcheck

type Service struct {
	Name       string
	StatusCode int
}

type InputJson struct {
	TargetServices []struct {
		ServiceName string `json:"service_name"`
		Endpoint    string `json:"endpoint"`
	} `json:"target_services"`
}
