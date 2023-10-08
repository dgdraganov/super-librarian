package model

type AppConfig struct {
	AppEnv   string
	Host     string
	Domain   string
	HttpPort string
	Secure   string
}

type MySqlConfig struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
}
