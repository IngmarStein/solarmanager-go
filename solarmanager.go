package solarmanager

import (
	"fmt"
	"net/url"
	"time"
)

type GatewayInfo struct {
	Id            string    `json:"_id"`           // db id for gateway
	Signal        string    `json:"signal"`        // gateway signal
	Name          string    `json:"name"`          // gateway name in system
	SmId          string    `json:"sm_id"`         // gateway unique id
	Owner         string    `json:"owner"`         // id of user - owner of gateway
	Firmware      string    `json:"firmware"`      // gateway firmware version
	LastErrorDate time.Time `json:"lastErrorDate"` // date of last error
	Mac           string    `json:"mac"`           // gateway mac address
	Ip            string    `json:"ip"`            // gateway ip
}

type GetGatewayInfoResponse GatewayInfo

type SensorInfo struct {
	Id          string `json:"_id"`
	Priority    int    `json:"priority"`
	DeviceType  string `json:"device_type"`
	Signal      string `json:"signal"`
	Type        string `json:"type"`
	DeviceGroup string `json:"device_group"`
	Ip          string `json:"ip"`
	Tag         struct {
		Id   string `json:"_id"`
		Name string `json:"name"`
	} `json:"tag"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetSensorsResponse []SensorInfo
type GetSensorResponse SensorInfo

type SensorData struct {
	Id                    string `json:"_id"`
	AccumulatedErrorCount int    `json:"accumulatedErrorCount"`
	CurrentPowerInvSm     int    `json:"currentPowerInvSm,omitempty"`
	CurrentEnergy         int    `json:"currentEnergy,omitempty"`
	Errors                []int  `json:"errors"`
	Signal                string `json:"signal"`
	ActiveDevice          int    `json:"activeDevice,omitempty"`
	CurrentPower          int    `json:"currentPower,omitempty"`
	SwitchState           int    `json:"switchState,omitempty"`
	CurrentWaterTemp      int    `json:"currentWaterTemp,omitempty"`
	Status                int    `json:"status,omitempty"`
	SOC                   int    `json:"SOC,omitempty"`
}

type GatewayData struct {
	InterfaceVersion              string       `json:"Interface Version"`
	TimeStamp                     time.Time    `json:"TimeStamp"`
	CurrentBatteryChargeDischarge int          `json:"currentBatteryChargeDischarge"`
	CurrentPowerConsumption       int          `json:"currentPowerConsumption"`
	CurrentPvGeneration           int          `json:"currentPvGeneration"`
	Devices                       []SensorData `json:"devices"`
	Errors                        []int        `json:"errors"`
	Soc                           int          `json:"soc"`
}

type GetSensorDataResponse struct {
	Date time.Time  `json:"date"`
	Data SensorData `json:"data"`
}

type SensorConsumptionStatistics struct {
	SensorId string `json:"sensorId"`
	Period   string `json:"period"`
	Data     []struct {
		CreatedAt   string  `json:"createdAt"`
		Consumption float64 `json:"consumption"`
	} `json:"data"`
	TotalConsumption float64 `json:"totalConsumption"`
}

type GatewayConsumptionStatistics struct {
	GatewayId string `json:"gatewayId"`
	Period    string `json:"period"`
	Data      []struct {
		CreatedAt   string `json:"createdAt"`
		Consumption int    `json:"consumption"`
		Production  int    `json:"production"`
	} `json:"data"`
	TotalConsumption int `json:"totalConsumption"`
}

type GetGatewayPieChartResponse struct {
	LastUpdate  time.Time `json:"lastUpdate"`
	Production  int       `json:"production"`
	Consumption int       `json:"consumption"`
	Battery     struct {
		Capacity           int `json:"capacity"`
		BatteryCharging    int `json:"batteryCharging"`
		BatteryDischarging int `json:"batteryDischarging"`
	} `json:"battery"`
	Arrows []struct {
		Direction string `json:"direction"`
		Value     int    `json:"value"`
	} `json:"arrows"`
}

type ForecastEntry struct {
	Timestamp int64 `json:"timestamp"`
	Expected  int   `json:"expected"`
	Min       int   `json:"min"`
	Max       int   `json:"max"`
}

type GetGatewayForecastResponse []ForecastEntry

type GetLowRateTariffResponse struct {
	MondayFridayFrom string `json:"Monday_Friday_from"`
	MondayFridayTo   string `json:"Monday_Friday_to"`
	SatudayFrom      string `json:"Satuday_from"`
	SatudayTo        string `json:"Satuday_to"`
	SundayFrom       string `json:"Sunday_from"`
	SundayTo         string `json:"Sunday_to"`
}

type HeatPumpOperationState int

const (
	NoInformation HeatPumpOperationState = iota
	Standby
	Heating
	WarmWater
	PartialError
	Failure
	Cooling
	EVU
	Defrosting
)

func (c *Client) GetGatewayInfo(solarManagerID string) (GetGatewayInfoResponse, error) {
	u := fmt.Sprintf("v1/info/gateway/%s", url.PathEscape(solarManagerID))

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return GetGatewayInfoResponse{}, err
	}
	var response GetGatewayInfoResponse
	_, err = c.do(req, &response)
	return response, err
}

func (c *Client) GetSensors(solarManagerID string) (GetSensorsResponse, error) {
	u := fmt.Sprintf("v1/info/sensors/%s", url.PathEscape(solarManagerID))

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return GetSensorsResponse{}, err
	}
	var response GetSensorsResponse
	_, err = c.do(req, &response)
	return response, err
}
