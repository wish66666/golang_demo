package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"golang_demo/config"
	"golang_demo/controllers/article"
	"golang_demo/controllers/auth"
	am "golang_demo/middleware/auth"
	ar "golang_demo/repositories/article"

	"github.com/gorilla/sessions"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo-contrib/session"
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
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	articleController := article.NewArticleController(ar.NewArticleRepository(db))
	authConttoller := auth.NewAuthController()

	e.GET("/", hello)
	e.POST("/login", authConttoller.Login)

	userGourp := e.Group("/user", am.AuthMiddleware)
	userGourp.GET("/articles", articleController.GetAll)

	e.Logger.Fatal(e.Start(":8080"))
}
