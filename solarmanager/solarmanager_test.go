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
	mux.HandleFunc("/v1/info/sensor/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
{
	"_id": "5da07dc5d32a997fd7fb80aa",
	"priority": 7,
	"device_type": "device",
	"signal": "connected",
	"type": "Water Heater",
	"device_group": "myPV AC THOR",
	"ip": "78.45.12.130",
	"createdAt": "2019-10-11T13:04:05.719Z",
	"updatedAt": "2020-02-20T11:46:12.358Z",
	"tag": null
}`))
	})
	mux.HandleFunc("/v1/stream/gateway/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
{
	"Interface Version": "1.0",
	"TimeStamp": "2021-02-01T16:09:39.744Z",
	"currentBatteryChargeDischarge": 0,
	"currentPowerConsumption": 494,
	"currentPvGeneration": 0,
	"devices": [
		{
			"_id": "5e07c29ecb03704972e486cc",
			"accumulatedErrorCount": 0,
			"currentPowerInvSm": 0,
			“currentEnergy”: 0
			"errors": [],
			"signal": "connected"
		},
		{
			"_id": "5e08ab9ccb03704972029a2c",
			"accumulatedErrorCount": 6,
			"currentPowerInvSm": 0,
			"errors": [],
			"signal": "connected"
		},
		{
			"_id": "5e0cd6f6dde8943e7179ebda",
			"accumulatedErrorCount": 0,
			"activeDevice": 0,
			"currentPower": 0,
			"errors": [],
			"signal": "connected",
			"switchState": 1
		},
		{
			"_id": "5f7d950deb88166c81c56f7a",
			"accumulatedErrorCount": 34,
			"activeDevice": 0,
			"currentPower": 0,
			"currentWaterTemp": 44,
			"errors": [],
			"signal": "connected",
			"status": 0
		},
		{
			"SOC": 0,
			"_id": "5d604d02b364481c2e0c72b5",
			"accumulatedErrorCount": 0,
			"activeDevice": 0,
			"currentPower": 0,
			"errors": [],
			"signal": "connected"
		},
		{
			"_id": "5ef0fe7cb9c6c4306c885133",
			"accumulatedErrorCount": 2067,
			"activeDevice": 0,
			"currentPower": 0,
			"errors": [
				1
			],
			"signal": "not connected",
			"switchState": 0
		},
		{
			"_id": "5d875f92f41d1c0df7b2ca7f",
			"accumulatedErrorCount": 2067,
			"activeDevice": 0,
			"currentPower": 0,
			"currentWaterTemp": 0,
			"errors": [
				5
			],
			"signal": "not connected"
		}
	],
	"errors": [],
	"soc": 0
}`))
	})
	mux.HandleFunc("/v1/consumption/sensor/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
{
	"sensorId": "5da6fdbf6f9aab5013a5cb9f",
	"period": "day",
	"data": [
		{
			"createdAt": "2021-01-01",
			"consumption": 120.83333429999996
		}
	],
	"totalConsumption": 120.83333429999996
}`))
	})
	mux.HandleFunc("/v1/consumption/gateway/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
{
	"gatewayId": "5c8fb8e7cdcda169da9d5fe3",
	"period": "day",
	"data": [
		{
			"createdAt": "2021-04-13",
			"consumption": 0,
			"production": 0
		}
	],
	"totalConsumption": 0
}`))
	})
	mux.HandleFunc("/v1/stream/sensor/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
{
	"date": "2021-03-25T09:26:37.037Z",
	"data": {
		"_id": "5db9b0bc2b591c467a947428",
		"signal": "connected",
		"currentPower": 1,
		"currentWaterTemp": 26,
		"errors": [
			1,
			15,
			10
		]
	}
}`))
	})
	mux.HandleFunc("/v1/chart/gateway/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
{
	"lastUpdate": "2021-04-07T15:24:55.608Z",
	"production": 20000,
	"consumption": 5000,
	"battery": {
		"capacity": 43,
		"batteryCharging": 0,
		"batteryDischarging": 0
	},
	"arrows": [
		{
			"direction": "fromPVToGrid",
			"value": 15000 // {number} - power in W
		},
		{
			"direction": "fromGridToConsumer",
			"value": 0 // {number} - power in W
		},
		{
			"direction": "fromPVToConsumer",
			"value": 5000 // {number} - power in W
		}
	]
}`))
	})
	mux.HandleFunc("/v1/forecast/gateways/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
[
	{
		"timestamp": 1641808800000,
		"expected": 1726,
		"min": 1183,
		"max": 2269
	},
	{
		"timestamp": 1641809700000,
		"expected": 1638,
		"min": 1120,
		"max": 2155
	},
]`))
	})
	mux.HandleFunc("/v1/low-rate-tariff/gateways/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
{
	"Monday_Friday_from": "20:00",
	"Monday_Friday_to": "07:00",
	"Satuday_from": "13:00",
	"Satuday_to": "07:00",
	"Sunday_from": "00:00",
	"Sunday_to": "07:00"
}`))
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

func ExampleClient_GetSensors() {
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
