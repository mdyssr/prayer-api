package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mdyssr/prayer-api/handlers"
	"os"
)

func main() {
	app := echo.New()

	app.Use(middleware.CORS())
	app.Use(middleware.Logger())

	app.GET("/", handlers.Home)
	app.GET("/api/times", handlers.GetPrayerTimes)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Logger.Fatal(app.Start(":" + port))
}
