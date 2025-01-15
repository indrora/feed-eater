package sources

import (
	"errors"
	"io"
)

var (
	// ErrInvalidSourceType is returned when an invalid source type is provided to NewSource.
	ErrInvalidSourceType = errors.New("invalid source type")
)

// DataSource represents a data source that can be configured and print output
type DataSource interface {
	Print(writer io.Writer)
	Configure(config map[string]string) error
}

// NewSource creates a Source from a type string and config
func NewSource(sourceType string, config map[string]string) (*DataSource, error) {
	var source DataSource
	switch sourceType {
	case "rssfeed":
		source = &RSSFeed{}
	case "weather":
		source = &WeatherReport{}
	case "textfile":
		source = &TextFile{}
	case "command":
		source = &CommandSource{}
	default:
		return nil, ErrInvalidSourceType
	}
	err := source.Configure(config)
	if err != nil {
		return nil, err
	}
	return &source, nil
}
