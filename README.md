# Golang SolarManager API Client

Go client for the [SolarManager API](https://external-web.solar-manager.ch)

## Usage

```go
package main

import (
    "fmt"
    "os"

    "github.com/ingmarstein/solarmanager-go/solarmanager"
)

func main() {
    username := os.Getenv("SOLARMANAGER_USERNAME")
    password := os.Getenv("SOLARMANAGER_PASSWORD")
    smID := os.Getenv("SOLARMANAGER_ID")

    client := solarmanager.NewClient(nil, nil, username, password)
    sensors, err := client.GetSensors(smID)
    if err != nil {
    	panic(err)
    }
    fmt.Println(sensors)
}
```