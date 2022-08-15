package core

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
	Version     string `env:"VERSION" envDefault:"1.0.0"`
	Port        string `env:"PORT" envDefault:"3000"`
}

var Configuration = Config{}

func init() {
	fmt.Println("Loading config...")
	if err := env.Parse(&Configuration); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(Configuration)
	fmt.Println("config loaded")
}
