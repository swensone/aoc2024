package config

import (
	"log"
	"os"

	"github.com/spf13/pflag"
)

// Config is the configuration for the application
type Config struct {
	Input string
	Debug bool
}

// Parse parses config flags and returns the configuration
func Parse() *Config {
	cfg := &Config{}

	fs := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	fs.StringVarP(&cfg.Input, "input", "i", "input.txt", "input file")
	fs.BoolVarP(&cfg.Debug, "debug", "d", false, "debug mode")

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}
