package config

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Search Search `toml:"search"`
}

type Search struct {
	Paths         []string `toml:"search_paths"`
	IgnorePattern string   `toml:"ignore_pattern"`
	IgnoreHidden  bool     `toml:"ignore_hidden"`
}

func Exists() bool {
	configDir, _ := os.UserConfigDir()
	configPath := filepath.Join(configDir, "tsesh", "config.toml")

	_, err := os.Stat(configPath)
	if err == nil {
		return true
	}

	return false
}

func CreateDefault() error {
	cfg := Config{
		Search: Search{
			Paths: []string{
				"~",
				"~/projects",
			},
			IgnorePattern: "![^()$]",
			IgnoreHidden:  true,
		},
	}

	buf := bytes.NewBuffer([]byte{})
	encoder := toml.NewEncoder(buf)
	if err := encoder.Encode(cfg); err != nil {
		return err
	}

	configDir, _ := os.UserConfigDir()
	configDirPath := filepath.Join(configDir, "tsesh")

	os.MkdirAll(configDirPath, os.ModePerm)
	return os.WriteFile(
		filepath.Join(configDirPath, "config.toml"),
		buf.Bytes(),
		os.ModePerm,
	)
}

func LoadInto(cfg *Config) error {
	configDir, _ := os.UserConfigDir()
	configPath := filepath.Join(configDir, "tsesh", "config.toml")

	f, err := os.Open(configPath)
	if err != nil {
		return err
	}

	decoder := toml.NewDecoder(f)

	return decoder.Decode(cfg)
}
