package config

import (
	errorFactory "IlluminateMyWords/src/utils/errors"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

const (
	DEFAULT_CONFIG_FILE = ".ilmw.yml"
	DEFAULT_COLOR       = "white"
)

var ColorCode = map[string]color.Attribute{
	"black":         color.FgBlack,
	"red":           color.FgRed,
	"green":         color.FgGreen,
	"yellow":        color.FgHiYellow,
	"blue":          color.FgBlue,
	"magenta":       color.FgMagenta,
	"cyan":          color.FgCyan,
	"orange":        color.FgYellow,
	"light gray":    color.FgHiWhite,
	"dark gray":     color.FgHiBlack,
	"light red":     color.FgHiRed,
	"light green":   color.FgHiGreen,
	"light yellow":  color.FgHiYellow,
	"light blue":    color.FgHiBlue,
	"light magenta": color.FgHiMagenta,
	"light cyan":    color.FgHiCyan,
	"white":         color.FgWhite,
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
	Verbose    *bool
	ConfigFile string
	Rules      map[string]MatchingRule
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

	Config.Rules = make(map[string]MatchingRule)
	for _, t := range Config.Mapping.Types {
		mr, err := CreateRule(&t.Keywords, t.Color)
		if err != nil {
			errorFactory.BuildAndLogError(err, "erorr while building rule for '%s'", t.Name)
		}
		Config.Rules[t.Name] = mr
	}

	return nil
}
