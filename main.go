package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GameChangerCorp/cari-kkn-be/api"
	"github.com/GameChangerCorp/cari-kkn-be/app/modules"
	"github.com/GameChangerCorp/cari-kkn-be/config"
	"github.com/GameChangerCorp/cari-kkn-be/utils"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	config := config.GetConfig()
	dbCon := utils.NewConnectionDatabase(config)

	defer dbCon.CloseConnection()

	controllers := modules.RegistrationModules(dbCon, config)
	e := gin.Default()

	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})
	api.RegistrationPath(e, controllers)
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	go func() {
		if port == "" {
			port = "8080"
		}
		address := fmt.Sprintf(":%s", port)

		if err := e.Run(address); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	<-quit
}
