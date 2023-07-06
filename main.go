package main

import (
	"fmt"
	"net/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/kelseyhightower/envconfig"
	"golang_demo/config"
	"golang_demo/models"

)


func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	var config config.Config
	envconfig.Process("APP", &config)

	fmt.Println(config.MySQL)

	dsn := config.MySQL.User + ":@tcp(" + config.MySQL.Host + ":" + config.MySQL.Port + ")/" + config.MySQL.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(err)

	var articles []models.Article
	if err := db.Find(&articles).Error; err != nil {
		panic(err)
	}
	fmt.Println(articles)

	e := echo.New()

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":8080"))
}
