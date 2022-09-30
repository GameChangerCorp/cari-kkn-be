package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GameChangerCorp/cari-kkn-be/config"
	"github.com/GameChangerCorp/cari-kkn-be/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	config := config.GetConfig()
	dbCon := utils.NewConnectionDatabase(config)

	defer dbCon.CloseConnection()

	// controllers := modules.RegistrationModules(dbCon, config)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API Is Active")
	})
	// api.RegistrationPath(e, controllers)
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	go func() {
		if port == "" {
			port = "8080"
		}
		address := fmt.Sprintf(":%s", port)

		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	<-quit
}
