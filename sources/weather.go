package sources

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/indrora/feed-eater/avweather"
	"github.com/paulmach/orb/geojson"
)

// WeatherReport implements the Source interface for WeatherReport weather data
type WeatherReport struct {
	station string
	hours   int
}

func (m *WeatherReport) Print(writer io.Writer) {

	weather, err := avweather.GetObservations(m.station, m.hours)

	if err != nil {
		fmt.Fprintf(writer, "Failed to retrieve weather: %v", err)
		return
	}

	if len(weather) < 1 {
		fmt.Fprintf(writer, "No weather data available")
		return
	}

	fmt.Fprintln(writer, "Weather report:")

	for _, observation := range weather {

		fmt.Fprintf(writer, "%s: %s ", observation.Name, observation.ReportTime)

		fmt.Fprintf(writer, "%f *C ", observation.Temp)
		if observation.WxString != nil {
			fmt.Fprintf(writer, " Observations: %s ", avweather.DecodeWeatherCodes(*observation.WxString))
		}

		// Wind
		w := observation.GetWind()
		fmt.Fprintf(writer, "Wind  %s at %d kts ", w.Direction, w.Speed)
		fmt.Fprintf(writer, "Visibility %v mi ", observation.Visib)

		fmt.Fprintf(writer, "Clouds:")
		for _, cloud := range observation.Clouds {
			fmt.Fprintf(writer, " %dft: %s ", cloud.Base, cloud.Cover)
		}
		fmt.Fprintf(writer, "\n")

	}

}

func (m *WeatherReport) Configure(config map[string]string) error {
	station, ok := config["station"]
	if !ok {
		return fmt.Errorf("station is required")
	}
	m.station = station
	hours, ok := config["hours"]
	if !ok {
		m.hours = 4
	} else {
		h, err := strconv.Atoi(hours)
		if err != nil {
			return fmt.Errorf("invalid hours value: %v", err)
		}
		m.hours = h
	}
	return nil
}

type WeatherAlert struct {
	// specify one or the other,
	State string
	Zone  string
}

func (m *WeatherAlert) Print(writer io.Writer) {
	fmt.Fprintf(writer, "Weather Alert: %s %s", m.State, m.Zone)

	// For a given state, you use https://api.weather.gov/alerts/active?area=WA
	// for a given zone, you use https://api.weather.gov/alerts/active?zone=AKZ001

	// construct the appropriate URL
	var url string
	if m.State != "" {
		url = fmt.Sprintf("https://api.weather.gov/alerts/active?area=%s", m.State)
	} else {
		url = fmt.Sprintf("https://api.weather.gov/alerts/active?zone=%s", m.Zone)
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(writer, "Failed to retrieve weather: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(writer, "Failed to retrieve weather: %v", err)
		return
	}

	geodata, err := geojson.UnmarshalFeatureCollection(body)
	if err != nil {
		fmt.Fprintf(writer, "Failed to retrieve weather: %v", err)
		return
	}
	for _, feature := range geodata.Features {
		fmt.Fprintf(writer, "Alert: %s\n", feature.Properties["headline"])
		fmt.Fprintf(writer, "Event: %s\n", feature.Properties["event"])
		fmt.Fprintf(writer, "Severity: %s\n", feature.Properties["severity"])
		fmt.Fprintf(writer, "Description: %s\n", feature.Properties["description"])
		fmt.Fprintf(writer, "Instruction: %s\n", feature.Properties["instruction"])
	}

}
func (m *WeatherAlert) Configure(config map[string]string) error {
	State, ok := config["state"]
	if !ok {

		// try zone
		zone, ok := config["zone"]
		if !ok {
			return fmt.Errorf("state or zone is required")
		} else {
			m.Zone = zone
		}
		return nil
	} else {
		m.State = State
		return nil
	}
}
