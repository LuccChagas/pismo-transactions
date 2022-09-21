package router

import (
	_ "pismo-transactions/docs/app"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// @title Pismo-Transactions Swagger API
// @version 1.0
// @description API Documentation

// @contact.name Luccas Machado
// @contact.url https://www.linkedin.com/in/luccas-machado-das-chagas-ab5897105/
// @contact.email luccaa.chagas23@gmail.com

// @host
// @BasePath /
func Handler() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(CorsConfig))
	e.GET("/health_check/", HealthCheck)
	e.POST("/auth", Auth)

	accounts := e.Group("/accounts")
	accounts.Use(middleware.JWTWithConfig(Config))

	accounts.POST("", CreateAccount)
	accounts.GET("", GetAccountByID)

	transactions := e.Group("/transaction")
	transactions.Use(middleware.JWTWithConfig(Config))

	transactions.POST("", CreateTransaction)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
