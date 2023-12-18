package solarmanager

import (
	"fmt"
	"net/url"
	"time"
)

type GatewayInfo struct {
	Id                      string    `json:"_id"`    // db id for gateway
	Signal                  string    `json:"signal"` // gateway signal
	Name                    string    `json:"name"`   // gateway name in system
	SmId                    string    `json:"sm_id"`  // gateway unique id
	Owner                   string    `json:"owner"`  // id of user - owner of gateway
	IsInstallationCompleted bool      `json:"isInstallationCompleted"`
	Firmware                string    `json:"firmware"`      // gateway firmware version
	LastErrorDate           time.Time `json:"lastErrorDate"` // date of last error
	Mac                     string    `json:"mac"`           // gateway mac address
	Ip                      string    `json:"ip"`            // gateway ip
}

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

type GetGatewayInfoResponse struct {
	Gateway GatewayInfo `json:"gateway"`

	Settings struct {
		OffsetWatt     int    `json:"offset_watt"`
		LowMFFrom      string `json:"low_m_f_from"`
		LowMFTo        string `json:"low_m_f_to"`
		LowSatFrom     string `json:"low_sat_from"`
		LowSatTo       string `json:"low_sat_to"`
		LowSunFrom     string `json:"low_sun_from"`
		LowSunTo       string `json:"low_sun_to"`
		KWp            int    `json:"kWp"`
		HouseFuse      int    `json:"houseFuse"`
		LoadManagement bool   `json:"loadManagement"`
		CommonSeasons  struct {
			MondayFriday []struct {
				From   string `json:"from"`
				Tariff string `json:"tariff"`
			} `json:"mondayFriday"`
			Saturday []struct {
				From   string `json:"from"`
				Tariff string `json:"tariff"`
			} `json:"saturday"`
			Sunday []struct {
				From   string `json:"from"`
				Tariff string `json:"tariff"`
			} `json:"sunday"`
		} `json:"commonSeasons"`
		HighTariff          float64 `json:"highTariff"`
		IsWinterTimeEnabled bool    `json:"isWinterTimeEnabled"`
		LowTariff           float64 `json:"lowTariff"`
		Provider            string  `json:"provider"`
		TariffType          string  `json:"tariffType"`
		WinterSeason        struct {
			MondayFriday []struct {
				From   string `json:"from"`
				Tariff string `json:"tariff"`
			} `json:"mondayFriday"`
			Saturday []struct {
				From   string `json:"from"`
				Tariff string `json:"tariff"`
			} `json:"saturday"`
			Sunday []struct {
				From   string `json:"from"`
				Tariff string `json:"tariff"`
			} `json:"sunday"`
		} `json:"winterSeason"`
	} `json:"settings"`
	User struct {
		FirstName    string `json:"first_name"`
		UserId       string `json:"user_id"`
		LastName     string `json:"last_name"`
		Email        string `json:"email"`
		Status       string `json:"status"`
		Country      string `json:"country"`
		City         string `json:"city"`
		Zip          string `json:"zip"`
		Plant        string `json:"plant"`
		CompanyName  string `json:"company_name"`
		Company      string `json:"company"`
		ConnectedOem string `json:"connectedOem"`
	} `json:"user"`
	Versions struct {
		SupportContract bool `json:"supportContract"`
	} `json:"versions"`
}

type GetSensorsResponse []SensorInfo

type GetSensorResponse SensorInfo

type GetGatewayDataResponse GatewayData

type GetSensorDataResponse struct {
	Date time.Time  `json:"date"`
	Data SensorData `json:"data"`
}

type GetSensorConsumptionStatisticsResponse struct {
	SensorId string `json:"sensorId"`
	Period   string `json:"period"`
	Data     []struct {
		CreatedAt   string  `json:"createdAt"`
		Consumption float64 `json:"consumption"`
	} `json:"data"`
	TotalConsumption float64 `json:"totalConsumption"`
}

type GetGatewayConsumptionStatisticsResponse struct {
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

type StatisticPeriod string

const (
	Day   = "day"
	Month = "month"
	Year  = "year"
)

func (c *Client) GetGatewayInfo(solarManagerID string) (GetGatewayInfoResponse, error) {
	u := fmt.Sprintf("v1/info/gateway/%s", url.PathEscape(solarManagerID))

	var response GetGatewayInfoResponse
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return response, err
	}
	_, err = c.do(req, &response)
	return response, err
}

func (c *Client) GetSensors(solarManagerID string) (GetSensorsResponse, error) {
	u := fmt.Sprintf("v1/info/sensors/%s", url.PathEscape(solarManagerID))

	var response GetSensorsResponse
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return response, err
	}
	_, err = c.do(req, &response)
	return response, err
}

func (c *Client) GetSensor(sensorID string) (GetSensorResponse, error) {
	u := fmt.Sprintf("v1/info/sensor/%s", url.PathEscape(sensorID))

	var response GetSensorResponse
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return response, err
	}
	_, err = c.do(req, &response)
	return response, err
}

func (c *Client) GetGatewayData(solarManagerID string) (GetGatewayDataResponse, error) {
	u := fmt.Sprintf("v1/info/stream/gateway/%s", url.PathEscape(solarManagerID))

	var response GetGatewayDataResponse
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return response, err
	}
	_, err = c.do(req, &response)
	return response, err
}

func (c *Client) GetSensorConsumptionStatistics(sensorID string, period StatisticPeriod) (GetSensorConsumptionStatisticsResponse, error) {
	u := fmt.Sprintf("v1/consumption/sensor/%s?period=%s",
		url.PathEscape(sensorID),
		url.QueryEscape(string(period)))

	var response GetSensorConsumptionStatisticsResponse
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return response, err
	}
	_, err = c.do(req, &response)
	return response, err
}

func (c *Client) GetGatewayConsumptionStatistics(solarManagerID string, period StatisticPeriod) (GetGatewayConsumptionStatisticsResponse, error) {
	u := fmt.Sprintf("v1/consumption/gateway/%s?period=%s",
		url.PathEscape(solarManagerID),
		url.QueryEscape(string(period)))

	var response GetGatewayConsumptionStatisticsResponse
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return response, err
	}
	_, err = c.do(req, &response)
	return response, err
}

func (c *Client) GetSensorData(solarManagerID string, sensorID string) (GetSensorDataResponse, error) {
	u := fmt.Sprintf("v1/stream/sensor/%s/%s", url.PathEscape(solarManagerID), url.PathEscape(sensorID))

	var response GetSensorDataResponse
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return response, err
	}
	_, err = c.do(req, &response)
	return response, err
}

func (c *Client) GetGatewayPieChart(solarManagerID string) (GetGatewayPieChartResponse, error) {
	u := fmt.Sprintf("v1/chart/gateway/%s", url.PathEscape(solarManagerID))

	var response GetGatewayPieChartResponse
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return response, err
	}
	_, err = c.do(req, &response)
	return response, err
}

func (c *Client) GetGatewayForecast(solarManagerID string) (GetGatewayForecastResponse, error) {
	u := fmt.Sprintf("v1/forecast/gateways/%s", url.PathEscape(solarManagerID))

	var response GetGatewayForecastResponse
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return response, err
	}
	_, err = c.do(req, &response)
	return response, err
}

func (c *Client) GetLowRateTariff(solarManagerID string) (GetLowRateTariffResponse, error) {
	u := fmt.Sprintf("v1/low-rate-tariff/gateways/%s", url.PathEscape(solarManagerID))

	var response GetLowRateTariffResponse
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return response, err
	}
	_, err = c.do(req, &response)
	return response, err
}
