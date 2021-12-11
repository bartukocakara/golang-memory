package config

import (
	"golang-memory/utils"
)

var (
	EXPORT_FILE_PATH = utils.GetEnv("EXPORT_FILE_PATH", "./tmp")
	API_PORT         = utils.GetEnv("API_PORT", "8080")
	RECORD_FREQ      = utils.GetEnv("RECORD_FREQ", "10")
)
