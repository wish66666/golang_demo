package main

import (
	"fmt"
	"net/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/kelseyhightower/envconfig"
)

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

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	var config Config
	envconfig.Process("APP", &config)

	fmt.Println(config.MySQL)

	dsn := config.MySQL.User + ":@tcp(" + config.MySQL.Host + ":" + config.MySQL.Port + ")/" + config.MySQL.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	e := echo.New()

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":8080"))
}
