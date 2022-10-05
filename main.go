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
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	fmt.Println(os.Getenv("MONGO_ETC"))

	port := os.Getenv("PORT")

	config := config.GetConfig()
	dbCon := utils.NewConnectionDatabase(config)

	defer dbCon.CloseConnection()

	controllers := modules.RegistrationModules(dbCon, config)
	// gin.SetMode(gin.ReleaseMode)
	e := gin.Default()
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))
	fmt.Println("cors enabled")

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
