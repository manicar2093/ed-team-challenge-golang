package config

import "os"

var (
	NomicsKey = ""
	NomicsAPI = "https://api.nomics.com/v1/currencies/sparkline"
	Port      = ":5000"
)

func InitEnv() {
	NomicsKey = os.Getenv("NOMICS_KEY")
	if NomicsKey == "" {
		panic("NOMICS_KEY is a must")
	}
}
