package config

type APIConfig struct {
	Host string `yaml:"address" env-default:"localhost"`
	Port string `yaml:"port" env-default:"8080"`
}
