package models

import (
	"fmt"
	"os"
	"strconv"
)

type Address struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (a Address) String() string {
	if a.Port == 0 {
		envPort := os.Getenv("PORT")
		parsedEnvPort, err := strconv.ParseInt(envPort, 10, 64)
		if err != nil {
			parsedEnvPort = int64(80)
		}
		a.Port = int(parsedEnvPort)
	}
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}
