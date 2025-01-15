package sources

import (
	"fmt"
	"io"
	"strconv"

	"github.com/indrora/feed-eater/avweather"
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

		fmt.Fprintf(writer, "%f *C", observation.Temp)
		if observation.WxString != nil {
			fmt.Fprintf(writer, " Observations: %s ", avweather.DecodeWeatherCodes(*observation.WxString))
		}

		// Wind
		w := observation.GetWind()
		fmt.Fprintf(writer, "Wind  %s at %f ", w.Direction, w.Speed)
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
