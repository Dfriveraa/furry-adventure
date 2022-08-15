package models

import "github.com/Sht97/furry-adventure/core"

type Root struct {
	Status  string `json:"status"`
	Environ string `json:"environment"`
	Version string `json:"version"`
}

var HealthCheck = Root{
	Status:  "Ok",
	Environ: core.Configuration.Environment,
	Version: core.Configuration.Version,
}
