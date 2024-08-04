package config

type DB struct {
	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	Name              string `yaml:"name"`
	MaxOpenConns      int    `yaml:"max_open_conns"`
	MaxIdleConns      int    `yaml:"max_idle_conns"`
	ConnMaxLifetime   int    `yaml:"conn_max_lifetime"`
	MaxIdleTime       int    `yaml:"max_idle_time"`
	MaxIdleTimeExceed int    `yaml:"max_idle_time_exceed"`
}
