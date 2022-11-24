package main

import (
	"flag"
	"log"
	"os"

	"github.com/akmalcm/team-soccer-go/configs"
	"github.com/akmalcm/team-soccer-go/models"

	"github.com/akmalcm/team-soccer-go/controllers"
	"github.com/akmalcm/team-soccer-go/seeds"

	"github.com/gin-gonic/gin"
)

func main() {
	handleArgs()
}

func handleArgs() {
	_, err := configs.DBConn()

	if err != nil {
		log.Fatal(err)
		return
	}

	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			seeds.Execute(configs.DB, args[1:]...)
			os.Exit(0)
		}
	} else {
		router := gin.Default()

		err = configs.DB.AutoMigrate(&models.Player{})

		if err != nil {
			log.Fatal(err)
			return
		}

		router.GET("/players", controllers.FindPlayers)
		router.GET("/players/:id", controllers.FindPlayer)
		router.POST("/players", controllers.CreatePlayer)
		router.PATCH("/players/:id", controllers.UpdatePlayer)
		router.DELETE("/players/:id", controllers.DeletePlayer)

		router.Run("localhost:8080")
	}
}
