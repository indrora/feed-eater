package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/indrora/feed-eater/sources"
)

type Config struct {
	General GeneralConfig `toml:"general"`
	Sources []Source      `toml:"sources"`
}

type GeneralConfig struct {
	Slow       bool `toml:"slow"`
	SpeedLimit int  `toml:"speed_limit"`
}

type Source struct {
	Type    string              `toml:"type"`
	Name    string              `toml:"name"`
	Options map[string]string   `toml:"options"`
	Impl    *sources.DataSource `toml:"-"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
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
