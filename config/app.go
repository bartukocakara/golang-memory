package config

import (
	"golang-memory/utils"
)

var (
	EXPORT_TO = utils.GetEnv("EXPORT_TO", "./data")
	API_PORT = utils.GetEnv("API_PORT", "8090")
	RECORD_CYCLE = utils.GetEnv("RECORD_CYCLE", "1")
)