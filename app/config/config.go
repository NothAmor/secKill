package config

type Config struct {
	Sys   Sys   `yaml:"sys"`
	Log   Log   `yaml:"log"`
	DB    DB    `yaml:"db"`
	Redis Redis `yaml:"redis"`
}
