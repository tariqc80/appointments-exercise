package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/tariqc80/appointments-exercise/internal/config"
	"github.com/tariqc80/appointments-exercise/internal/data"
	"github.com/tariqc80/appointments-exercise/internal/route"
)

func main() {
	// get Gin engine
	engine := gin.Default()

	// get config values from environment
	config := config.ParseFromEnv()

	// open connection to database using config
	database, err := data.Connect(config)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// get instance of data provider using the database connection
	scheduleData := data.NewSchedule(database)

	// get instance of route handler using data provider
	scheduleRouter := route.NewSchedule(scheduleData)

	// setup the routes using Gin engine
	scheduleRouter.ConfigRoutes(engine)

	// start the api server
	engine.Run("0.0.0.0:8080")
}
