package models

type Config struct {
	Server Address `yaml:"server"`
	Routes []Route `yaml:"routes"`
}
