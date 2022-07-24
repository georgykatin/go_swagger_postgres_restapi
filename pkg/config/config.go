package config

type Config struct {
	ConnectionString string `json:"connection_string"`
	Port             int    `json:"port"`
}
