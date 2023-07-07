package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"golang_demo/config"
	"golang_demo/controllers/article"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
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
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	articleController := article.NewArticleController(db)

	e.GET("/", hello)
	e.GET("/articles", articleController.GetAll)

	e.Logger.Fatal(e.Start(":8080"))
}
