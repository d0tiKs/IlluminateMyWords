package main

import (
	configutils "IlluminateMyWords/src/utils/config"
	loggerutils "IlluminateMyWords/src/utils/logger"
	"flag"
)

func initCLI() {
	configFile := flag.String("conf", configutils.DEFAULT_CONFIG_FILE, "The configuration file to load.")
	verbose := flag.Bool("verbose", false, "Enable debug logs.")

	config := configutils.LoadConfig(configFile, verbose)

	flag.Parse()

	loggerutils.LogMessage(loggerutils.LOG_DEBUG, "config file: %s, verbosity: %v", *config.ConfigFile, *config.Verbose)
}

func main() {
	initCLI()
}
