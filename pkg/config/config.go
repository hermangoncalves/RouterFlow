package config

type Server struct {
	Name     string `mapstructure:"name" yaml:"name"`
	Port     int    `mapstructure:"port" yaml:"port"`
	Host     string `mapstructure:"host" yaml:"host"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
}

type Config struct {
	Servers []Server `mapstructure:"servers" yaml:"servers"`
}
