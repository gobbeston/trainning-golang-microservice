package main

import (
	//"github.com/gobbeston/boredom"
	//"github.com/gobbeston/ginney"
	//"github.com/gobbeston/trainning-golang-microservice/app/constants"
	"github.com/gobbeston/trainning-golang-microservice/app/database"
	"github.com/gobbeston/trainning-golang-microservice/app/env"

	_ "github.com/gobbeston/trainning-golang-microservice/app/docs"
	_healthcheck "github.com/gobbeston/trainning-golang-microservice/app/layers/deliveries/http/health_check"

	customerHandler "github.com/gobbeston/trainning-golang-microservice/app/layers/deliveries/http/customer"
	customerRepo "github.com/gobbeston/trainning-golang-microservice/app/layers/repositories/customer"
	customerUseCase "github.com/gobbeston/trainning-golang-microservice/app/layers/usecases/customer"

	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pseidemann/finish"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Trainning Golang MicroService
// @version 0.1.0
// @description Trainning Golang MicroService.
// @schemes http https
// @BasePath /

func main() {
	env.Init()
	//boredom.InitL(env.LogLevel)
	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())

	_healthcheck.NewEndpointHTTPHandler(ginEngine)
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//ginEngine.Use(ginney.LogWithCorrelationIdMiddleware(gin.DefaultWriter, constants.NoLoggedRoutes))
	//ginEngine.Use(ginney.MicroServiceCorrelationIdMiddleware())
	//ginEngine.Use(ginney.FromGinContextToContextMiddleware())

	dbConn := database.ConnectDB()
	defer func() {
		_ = dbConn.Close()
	}()
	database.DBMigration()

	customerRepo := customerRepo.InitRepo(dbConn)
	customerUseCase := customerUseCase.InitUseCase(customerRepo)
	customerHandler.NewEndpointHttpHandler(ginEngine, customerUseCase)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: ginEngine,
	}

	timeOut, err := strconv.Atoi(os.Getenv("GRACEFUL_TIMEOUT"))
	if err != nil {
		timeOut = 30 // second
	}

	graceful := &finish.Finisher{Timeout: time.Duration(timeOut) * time.Second}
	graceful.Add(srv)

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	graceful.Wait()
}
