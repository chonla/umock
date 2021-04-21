package models

import "fmt"

type Address struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (a Address) String() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}
