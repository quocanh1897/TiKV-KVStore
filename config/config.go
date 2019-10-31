package config

type Config struct {
	PDConfig []string
}

func GetConfig() *Config {
	return &Config{
		PDConfig: []string{"pd0:2379","pd1:2379","pd2:2379"},
	}
}
