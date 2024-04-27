package config

import (
	errorFactory "IlluminateMyWords/src/utils/errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	DEFAULT_CONFIG_FILE = ".ilmw.yml"
)

var ColorCode = map[string]string{
	"black":         "30",
	"red":           "31",
	"green":         "32",
	"yellow":        "33",
	"blue":          "34",
	"magenta":       "35",
	"cyan":          "36",
	"light gray":    "37",
	"dark gray":     "90",
	"light red":     "91",
	"light green":   "92",
	"light yellow":  "93",
	"light blue":    "94",
	"light magenta": "95",
	"light cyan":    "96",
	"white":         "97",
}

var Config appConfig

type KeywordsMapping struct {
	Types []struct {
		Name     string   `yaml:"name"`
		Color    string   `yaml:"color"`
		Keywords []string `yaml:"keywords"`
	} `yaml:"types"`
}

type appConfig struct {
	ConfigFile string
	Verbose    *bool
	Mapping    KeywordsMapping
}

func GetConfig() appConfig {
	return Config
}

func InitConfig(conf *string, verbose *bool) (appConfig, error) {
	Config = appConfig{
		ConfigFile: "",
		Verbose:    verbose,
	}
	filePath, err := filepath.Abs(*conf)
	if err != nil {
		return Config, errorFactory.BuildError(err, "error while at path '%s'", *conf)
	}
	Config.ConfigFile = filePath
	return Config, nil
}

func LoadConfig() error {
	file, err := os.ReadFile(Config.ConfigFile)
	if err != nil {
		return errorFactory.BuildError(err, "error while opening file '%s'", Config.ConfigFile)
	}
	err = yaml.Unmarshal(file, &Config.Mapping)
	if err != nil {
		return errorFactory.BuildError(err, "error while parsing the file '%s'", Config.ConfigFile)
	}

	return nil
}
