package model

type ServiceConfig struct {
	AppEnv   string
	Host     string
	HttpPort string
}

type MySqlConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
}
