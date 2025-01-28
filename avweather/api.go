package avweather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetObservations(where string, hours int) (Observations, error) {

	callurl := fmt.Sprintf("https://aviationweather.gov/api/data/metar?ids=%s&hours=%d&format=json", where, hours)

	uu, err := url.Parse(callurl)

	if err != nil {
		return nil, err
	}
	resp, err := http.Get(uu.String())
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, err
	}

	bodybytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	var obs Observations
	err = json.Unmarshal(bodybytes, &obs)
	return obs, err
}

type Wind struct {
	Direction string
	Speed     int
}

func degreesToCardinal(degrees float64) string {
	directions := []string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"}
	index := int((degrees + 11.25) / 22.5)
	return directions[index%16]
}

func (o Observation) GetWind() Wind {

	if o.Wspd == nil {
		return Wind{
			Direction: "Calm",
			Speed:     0,
		}
	}

	spd := *(o.Wspd)
	if spd == 0 {
		return Wind{
			Direction: "Calm",
			Speed:     0,
		}
	}

	switch o.Wdir.(type) {
	case int:
		return Wind{
			Direction: degreesToCardinal(float64(o.Wdir.(int))),
			Speed:     spd,
		}
	case float64:
		return Wind{
			Direction: degreesToCardinal(o.Wdir.(float64)),
			Speed:     spd,
		}
	case string:
		return Wind{
			Direction: o.Wdir.(string),
			Speed:     spd,
		}
	default:
		return Wind{
			Direction: "Unknown",
			Speed:     spd,
		}
	}

}
