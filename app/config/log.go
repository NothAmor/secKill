package config

type Log struct {
	Level     string `yaml:"level"`
	Caller    bool   `yaml:"caller"`
	Timestamp bool   `yaml:"timestamp"`
	Prefix    string `yaml:"prefix"`
}
