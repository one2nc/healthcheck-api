package main

import (
	"testing"
)

type Svc struct {
	ServiceName string
	Endpoint    string
}

func TestParseJson(t *testing.T) {
	got := parseJson("./test_endpoints.json")
	want := Svc{
		ServiceName: "google",
		Endpoint:    "https://www.google.com/",
	}

	if got[0].TargetServices[0].ServiceName != want.ServiceName {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if got[0].TargetServices[0].Endpoint != want.Endpoint {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
