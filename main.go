package main

import (
	"log"
	"service-user-management/src/config"
	"service-user-management/src/entity"
	"service-user-management/src/wire"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panicln(err)
	}
	db := config.InitDB()
	db.AutoMigrate(&entity.User{})
	e := echo.New()

	uw := wire.UserWire(db)

	// api/user group route
	user := e.Group("/api/user")
	user.GET("", uw.SearchUser)
	user.POST("", uw.AddUser)
	user.POST("/login", uw.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
