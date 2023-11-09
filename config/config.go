package config

type Config struct {
	Port    string
	DBURL   string
	MongoURL string
}

var cfg = &Config{
	Port:    ":8080",
	DBURL:   "postgres://",
	MongoURL: "mongodburl",
}

func GetConfig() *Config {
	return cfg
}