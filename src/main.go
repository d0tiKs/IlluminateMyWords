package main

import (
	configutils "IlluminateMyWords/src/utils/config"
	loggerutils "IlluminateMyWords/src/utils/logger"
	"flag"
)

func initCLI() {
	configFile := flag.String("conf", configutils.DEFAULT_CONFIG_FILE, "The configuration file to load.")
	verbose := flag.Bool("verbose", false, "Enable debug logs.")
	flag.Parse()

	configutils.InitConfig(configFile, verbose)

	loggerutils.LogMessage(loggerutils.LOG_DEBUG, "config file: %s, verbosity: %v", *configFile, *verbose)
}

func main() {
	initCLI()

	err := configutils.LoadConfig()
	if err != nil {
		loggerutils.LogMessage(loggerutils.LOG_ERROR, err.Error())
		return
	}
}
