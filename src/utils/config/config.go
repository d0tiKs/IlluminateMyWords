package configutils

const (
	DEFAULT_CONFIG_FILE = ".ilmw.yml"
)

var Config processConfig

type processConfig struct {
	ConfigFile *string
	Verbose    *bool
}

func GetConfig() processConfig {
	return Config
}

func LoadConfig(conf *string, verbose *bool) processConfig {
	Config = processConfig{
		ConfigFile: conf,
		Verbose:    verbose,
	}

	return Config
}
