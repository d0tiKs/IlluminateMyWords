package main

import (
	"IlluminateMyWords/src/config"
	"IlluminateMyWords/src/utils/logger"
	"flag"
)

func initCLI() {
	configFile := flag.String("conf", config.DEFAULT_CONFIG_FILE, "The configuration file to load.")
	verbose := flag.Bool("verbose", false, "Enable debug logs.")
	flag.Parse()

	config.InitConfig(configFile, verbose)

	logger.InitLogger(verbose, logger.STAND_FD)
	logger.LogMessage(logger.LOG_DEBUG, "config file: %s, verbosity: %v", *configFile, *verbose)
}

func main() {
	initCLI()

	err := config.LoadConfig()
	if err != nil {
		logger.LogMessage(logger.LOG_ERROR, err.Error())
		return
	}
}
