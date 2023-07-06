package config

type Config struct {
	MySQL MySQL
}
type MySQL struct {
	Host string
	Port string
	User string
	Password string
	Database string
}
