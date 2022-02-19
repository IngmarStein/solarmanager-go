package solarmanager

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

// newTestServer returns a *httptest.Server serving mock responses for the SolarManager API.
func newTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/info/gateway/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
{
	"_id": "5c8fb8e7cdcda169da9d5fe3",
	"signal": "connected",
	"name": "1234123412341234",
	"sm_id": "1234123412341234",
	"owner": "5c8fb8fccdcda169da9d5fe4",
	"firmware": "0.20.1",
	"lastErrorDate": "2021-01-20T06:00:01.150Z",
	"mac": "2C:3E:51:C5:A0:CA",
	"ip": "192.168.1.51"
}
`))
	})
	mux.HandleFunc("/v1/info/sensors/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
[
	{
		"_id": "5da07dc5d32a997fd7fb80aa",
		"priority": 7,
		"device_type": "device",
		"signal": "connected",
		"type": "Water Heater",
		"device_group": "myPV AC THOR",
		"ip": "78.45.12.130",
		"tag": null,
		"createdAt": "2019-10-16T11:23:43.229Z",
		"updatedAt": "2020-10-20T11:31:14.842Z"
	},
	{
		"_id": "5da6fdbf6f9aab5013a5cb9f",
		"priority": 5,
		"device_type": "device",
		"signal": "connected",
		"device_group": "KEBA Wallbox P30",
		"type": "Car Charging",
		"createdAt": "2019-10-16T11:23:43.229Z",
		"updatedAt": "2020-10-20T11:31:14.842Z",
		"tag": {
			"_id": "5dfb86da1cdd557b82303f92",
			"name": "Tag_1"
		},
		"ip": "1.5.5.10"
	}
]`))
	})
	return httptest.NewServer(mux)
}

func newTestClient(t *testing.T, svr *httptest.Server) *Client {
	t.Helper()
	baseURL, err := url.Parse(svr.URL)
	if err != nil {
		t.Fatal(err)
	}
	return NewClient(nil, baseURL, "username", "password")
}

func TestGetGatewayInfo(t *testing.T) {
	svr := newTestServer()
	defer svr.Close()
	client := newTestClient(t, svr)
	resp, err := client.GetGatewayInfo("")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Owner != "5c8fb8fccdcda169da9d5fe4" {
		t.Fatalf("unexpected owner, expected 5c8fb8fccdcda169da9d5fe4, but got %s", resp.Owner)
	}
}

func TestGetSensors(t *testing.T) {
	svr := newTestServer()
	defer svr.Close()
	client := newTestClient(t, svr)
	resp, err := client.GetSensors("")
	if err != nil {
		t.Fatal(err)
	}

	if len(resp) != 2 {
		t.Fatalf("unexpected number of sensors, expected 2, but got %d", len(resp))
	}

	if resp[0].DeviceGroup != "myPV AC THOR" {
		t.Fatalf("unexpected device group for sensor 0, expected %q, but got %s", "myPV AC THOR", resp[0].DeviceGroup)
	}
}

func ExampleGetSensors() {
	username := os.Getenv("SOLARMANAGER_USERNAME")
	password := os.Getenv("SOLARMANAGER_PASSWORD")
	smID := os.Getenv("SOLARMANAGER_ID")

	client := NewClient(nil, nil, username, password)
	sensors, err := client.GetSensors(smID)
	if err != nil {
		panic(err)
	}
	fmt.Println(sensors)
}
