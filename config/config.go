package config

type Config struct {
	GRPCPort int
	GRPCHost string
	PDConfig []string
}

func GetConfig() *Config {
	return &Config{
		GRPCPort: 10002,
		GRPCHost: "localhost",
		PDConfig: []string{"pd0:2379","pd1:2379","pd2:2379"},
	}
}
