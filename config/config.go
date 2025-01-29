package config

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	fio "github.com/indrora/feed-eater/io"
	"github.com/indrora/feed-eater/sources"
	"go.bug.st/serial"
)

type Config struct {
	General GeneralConfig `toml:"general"`
	Sources []Source      `toml:"sources"`
	Output  OutputConfig  `toml:"output"`
}

type GeneralConfig struct {
	Slow       bool `toml:"slow"`
	SpeedLimit int  `toml:"speed_limit"`
	Loop       bool `toml:"loop"`
	Shuffle    bool `toml:"shuffle"`
}

type Source struct {
	Inhibit bool                `toml:"ignore"`
	Type    string              `toml:"type"`
	Name    string              `toml:"name"`
	Options map[string]string   `toml:"options"`
	Impl    *sources.DataSource `toml:"-"`
}

type OutputConfig struct {
	Type      string `toml:"type"` // "serial", "stdio", "fifo"
	Device    string `toml:"device"`
	Mode      string `toml:"mode"`
	FilterTTY bool   `toml:"filter_tty"`
}

func parseSerialMode(mode string) (serial.Mode, error) {
	parts := strings.Fields(mode)
	if len(parts) != 2 {
		return serial.Mode{}, fmt.Errorf("invalid mode format: %s", mode)
	}

	baud, err := strconv.Atoi(parts[0])
	if err != nil {
		return serial.Mode{}, fmt.Errorf("invalid baud rate: %s", parts[0])
	}

	format := parts[1]
	if len(format) != 3 {
		return serial.Mode{}, fmt.Errorf("invalid format: %s", format)
	}

	dataBits := format[0] - '0'
	if dataBits != 7 && dataBits != 8 {
		return serial.Mode{}, fmt.Errorf("invalid data bits: %c", format[0])
	}

	var parity serial.Parity
	switch format[1] {
	case 'N':
		parity = serial.NoParity
	case 'E':
		parity = serial.EvenParity
	case 'O':
		parity = serial.OddParity
	default:
		return serial.Mode{}, fmt.Errorf("invalid parity: %c", format[1])
	}

	var stopBits serial.StopBits
	switch format[2] {
	case '1':
		stopBits = serial.OneStopBit
	case '2':
		stopBits = serial.TwoStopBits
	default:
		return serial.Mode{}, fmt.Errorf("invalid stop bits: %c", format[2])
	}

	return serial.Mode{
		BaudRate: baud,
		DataBits: int(dataBits),
		Parity:   parity,
		StopBits: stopBits,
	}, nil
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}

	if config.Output.Type == "" {
		config.Output.Type = "stdio"
	}

	for i, source := range config.Sources {
		implSource, err := sources.NewSource(source.Type, source.Options)
		if err != nil {
			return nil, fmt.Errorf("Failed to create source: %v", err)
		}
		config.Sources[i].Impl = implSource
	}

	return &config, nil
}

func (c *Config) GetOutput() (io.ReadWriteCloser, error) {
	switch c.Output.Type {
	case "serial":
		mode, err := parseSerialMode(c.Output.Mode)
		if err != nil {
			return nil, fmt.Errorf("failed to parse: %v", err)
		}
		k, err := serial.Open(c.Output.Device, &mode)
		if err != nil {
			return nil, fmt.Errorf("failed to open serial port: %v", err)
		}
		return k, nil

	case "stdio":
		return fio.Glue2(os.Stdin, os.Stdout), nil

	case "fifo":
		return os.OpenFile(c.Output.Device, os.O_RDWR, os.ModeNamedPipe)

	default:
		return nil, fmt.Errorf("unknown output type: %s", c.Output.Type)
	}
}
