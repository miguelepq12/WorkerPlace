package main

import (
	"WorkerPlace/app/application"
	"WorkerPlace/app/infrastructure/kafka/access/consumer"
	"WorkerPlace/app/infrastructure/mongodb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.HideBanner = true

	db,_ := mongodb.NewMongoDbConnection()
	defer db.Disconnect()
	repository := mongodb.NewAccessMongoDbRepository(db)
	usecase :=  application.NewAccessRecordUseCase(repository)
	workerAccess := consumer.NewAccessRecordKafkaConsumer(usecase)
	workerAccess.Run()

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}

