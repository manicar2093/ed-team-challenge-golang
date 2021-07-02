package config

import "os"

var (
	FilesPath   = ""
	NomicsKey   = ""
	NominicsAPI = "https://api.nomics.com/v1/currencies/sparkline"
	Port        = ":5000"
)

func InitEnv() {
	FilesPath = os.Getenv("FILES_PATH")
	if FilesPath == "" {
		panic("FILES_PATH is a must")
	}
	NomicsKey = os.Getenv("NOMICS_KEY")
	if NomicsKey == "" {
		panic("NOMICS_KEY is a must")
	}
}
