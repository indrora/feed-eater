package avweather

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	intensityMap = map[string]string{
		"-": "Light ",
		"+": "Heavy ",
	}

	descriptorMap = map[string]string{
		"MI": "Shallow",
		"PR": "Partial",
		"BC": "Patches",
		"DR": "Low Drifting",
		"BL": "Blowing",
		"SH": "Showers",
		"TS": "Thunderstorms",
		"FZ": "Freezing",
		"RE": "(Recent)",
		"VC": "in vicinity",
	}

	precipMap = map[string]string{
		"DZ": "Drizzle",
		"RA": "Rain",
		"SN": "Snow",
		"SG": "Grainy Snow",
		"IC": "Ice Crystals",
		"PL": "Ice Pellets",
		"GR": "Hail",
		"GS": "Graupel",
		"UP": "Unknown Precipitation",
	}

	obscurationMap = map[string]string{
		"BR": "Mist",
		"FG": "Fog",
		"FU": "Smoke",
		"VA": "Volcanic Ash",
		"DU": "Dust",
		"SA": "Sand",
		"HZ": "Haze",
		"PY": "Spray",
	}

	otherMap = map[string]string{
		"PO": "Dust Whirls",
		"SQ": "Squall",
		"FC": "Funnel Cloud",
		"SS": "Sandstorm",
		"DS": "Duststorm",
	}
)

func DecodeWeatherCodes(input string) string {
	if input == "" {
		return ""
	}

	groups := strings.Split(input, " ")
	var results []string

	for _, group := range groups {
		timeRegex := regexp.MustCompile(`[BE](\d{2})`)
		times := timeRegex.FindAllStringSubmatch(group, -1)

		weatherPart := timeRegex.ReplaceAllString(group, "")

		var intensity, descriptor, phenomena []string

		// Check for intensity
		for prefix, desc := range intensityMap {
			if strings.HasPrefix(weatherPart, prefix) {
				intensity = append(intensity, desc)
				weatherPart = strings.TrimPrefix(weatherPart, prefix)
			}
		}

		// Process remaining weather codes in 2-character chunks
		for len(weatherPart) > 0 {
			if len(weatherPart) < 2 {
				phenomena = append(phenomena, "(Unknown) "+weatherPart)
				break
			}

			code := weatherPart[:2]
			weatherPart = weatherPart[2:]

			if desc, ok := descriptorMap[code]; ok {
				descriptor = append(descriptor, desc)
			} else if precip, ok := precipMap[code]; ok {
				phenomena = append(phenomena, precip)
			} else if obsc, ok := obscurationMap[code]; ok {
				phenomena = append(phenomena, obsc)
			} else if other, ok := otherMap[code]; ok {
				phenomena = append(phenomena, other)
			} else {
				phenomena = append(phenomena, "(Unknown) "+code)
			}
		}

		var result string
		if len(intensity) > 0 {
			result += strings.Join(intensity, " ")
		}
		if len(descriptor) > 0 {
			result += strings.Join(descriptor, " ") + " "
		}
		if len(phenomena) > 0 {
			result += strings.Join(phenomena, ", ")
		}

		if len(times) > 0 {
			timeDesc := ""
			for _, t := range times {
				if strings.HasPrefix(group, "B") {
					timeDesc = fmt.Sprintf("from :%s", t[1])
				} else if strings.HasPrefix(group, "E") {
					timeDesc = fmt.Sprintf("until :%s", t[1])
				}
			}
			if timeDesc != "" {
				result += " " + timeDesc
			}
		}

		results = append(results, strings.TrimSpace(result))
	}

	return strings.Join(results, ", ")
}
