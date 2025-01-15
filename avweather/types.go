package avweather

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Set of surface observations (METARs)
type Observations []Observation
type Observation struct {
	// Altim corresponds to the JSON schema field "altim".
	Altim float64 `json:"altim"`

	// Clouds corresponds to the JSON schema field "clouds".
	Clouds []ObservationsElemCloudsElem `json:"clouds"`

	// Dewp corresponds to the JSON schema field "dewp".
	Dewp *float64 `json:"dewp"`

	// Elev corresponds to the JSON schema field "elev".
	Elev int `json:"elev"`

	// IcaoId corresponds to the JSON schema field "icaoId".
	IcaoId string `json:"icaoId"`

	// Lat corresponds to the JSON schema field "lat".
	Lat float64 `json:"lat"`

	// Lon corresponds to the JSON schema field "lon".
	Lon float64 `json:"lon"`

	// MaxT corresponds to the JSON schema field "maxT".
	MaxT *float64 `json:"maxT"`

	// MaxT24 corresponds to the JSON schema field "maxT24".
	MaxT24 *float64 `json:"maxT24"`

	// MetarType corresponds to the JSON schema field "metarType".
	MetarType ObservationsElemMetarType `json:"metarType"`

	// MetarId corresponds to the JSON schema field "metar_id".
	MetarId int `json:"metar_id"`

	// MinT corresponds to the JSON schema field "minT".
	MinT *float64 `json:"minT"`

	// MinT24 corresponds to the JSON schema field "minT24".
	MinT24 *float64 `json:"minT24"`

	// MostRecent corresponds to the JSON schema field "mostRecent".
	MostRecent int `json:"mostRecent"`

	// Name corresponds to the JSON schema field "name".
	Name string `json:"name"`

	// ObsTime corresponds to the JSON schema field "obsTime".
	ObsTime int `json:"obsTime"`

	// Pcp24Hr corresponds to the JSON schema field "pcp24hr".
	Pcp24Hr *float64 `json:"pcp24hr"`

	// Pcp3Hr corresponds to the JSON schema field "pcp3hr".
	Pcp3Hr *float64 `json:"pcp3hr"`

	// Pcp6Hr corresponds to the JSON schema field "pcp6hr".
	Pcp6Hr *float64 `json:"pcp6hr"`

	// Precip corresponds to the JSON schema field "precip".
	Precip *float64 `json:"precip"`

	// PresTend corresponds to the JSON schema field "presTend".
	PresTend *float64 `json:"presTend"`

	// Prior corresponds to the JSON schema field "prior".
	Prior int `json:"prior"`

	// QcField corresponds to the JSON schema field "qcField".
	QcField int `json:"qcField"`

	// RawOb corresponds to the JSON schema field "rawOb".
	RawOb string `json:"rawOb"`

	// ReceiptTime corresponds to the JSON schema field "receiptTime".
	ReceiptTime string `json:"receiptTime"`

	// ReportTime corresponds to the JSON schema field "reportTime".
	ReportTime string `json:"reportTime"`

	// Slp corresponds to the JSON schema field "slp".
	Slp *float64 `json:"slp"`

	// Snow corresponds to the JSON schema field "snow".
	Snow *float64 `json:"snow"`

	// Temp corresponds to the JSON schema field "temp".
	Temp float64 `json:"temp"`

	// VertVis corresponds to the JSON schema field "vertVis".
	VertVis *int `json:"vertVis"`

	// Visib corresponds to the JSON schema field "visib".
	Visib interface{} `json:"visib"`

	// Wdir corresponds to the JSON schema field "wdir".
	Wdir interface{} `json:"wdir"`

	// Wgst corresponds to the JSON schema field "wgst".
	Wgst *int `json:"wgst"`

	// Wspd corresponds to the JSON schema field "wspd".
	Wspd *int `json:"wspd"`

	// WxString corresponds to the JSON schema field "wxString".
	WxString *string `json:"wxString"`
}

type ObservationsElemCloudsElem struct {
	// Base corresponds to the JSON schema field "base".
	Base int `json:"base"`

	// Cover corresponds to the JSON schema field "cover".
	Cover ObservationsElemCloudsElemCover `json:"cover"`
}

type ObservationsElemCloudsElemCover string

const ObservationsElemCloudsElemCoverBKN ObservationsElemCloudsElemCover = "BKN"
const ObservationsElemCloudsElemCoverCAVOK ObservationsElemCloudsElemCover = "CAVOK"
const ObservationsElemCloudsElemCoverCLR ObservationsElemCloudsElemCover = "CLR"
const ObservationsElemCloudsElemCoverFEW ObservationsElemCloudsElemCover = "FEW"
const ObservationsElemCloudsElemCoverOVC ObservationsElemCloudsElemCover = "OVC"
const ObservationsElemCloudsElemCoverOVX ObservationsElemCloudsElemCover = "OVX"
const ObservationsElemCloudsElemCoverSCT ObservationsElemCloudsElemCover = "SCT"

var enumValues_ObservationsElemCloudsElemCover = []interface{}{
	"CLR",
	"CAVOK",
	"FEW",
	"SCT",
	"BKN",
	"OVC",
	"OVX",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObservationsElemCloudsElemCover) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ObservationsElemCloudsElemCover {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ObservationsElemCloudsElemCover, v)
	}
	*j = ObservationsElemCloudsElemCover(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObservationsElemCloudsElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["base"]; raw != nil && !ok {
		return fmt.Errorf("field base in ObservationsElemCloudsElem: required")
	}
	if _, ok := raw["cover"]; raw != nil && !ok {
		return fmt.Errorf("field cover in ObservationsElemCloudsElem: required")
	}
	type Plain ObservationsElemCloudsElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ObservationsElemCloudsElem(plain)
	return nil
}

type ObservationsElemMetarType string

const ObservationsElemMetarTypeBUOY ObservationsElemMetarType = "BUOY"
const ObservationsElemMetarTypeCMAN ObservationsElemMetarType = "CMAN"
const ObservationsElemMetarTypeMETAR ObservationsElemMetarType = "METAR"
const ObservationsElemMetarTypeSPECI ObservationsElemMetarType = "SPECI"
const ObservationsElemMetarTypeSYNOP ObservationsElemMetarType = "SYNOP"

var enumValues_ObservationsElemMetarType = []interface{}{
	"METAR",
	"SPECI",
	"SYNOP",
	"BUOY",
	"CMAN",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObservationsElemMetarType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ObservationsElemMetarType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ObservationsElemMetarType, v)
	}
	*j = ObservationsElemMetarType(v)
	return nil
}
