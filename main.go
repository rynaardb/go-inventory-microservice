package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"rynaardb.com/inventory-service/common"
	"rynaardb.com/inventory-service/controllers"
	"rynaardb.com/inventory-service/databases"
	"rynaardb.com/inventory-service/docs"
)

// Main manages the main golang application
type Main struct {
	router *gin.Engine
}

// Initialized the server
func (m *Main) initServer() error {
	var err error
	// Load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize postgres database
	err = databases.Database.Init()
	if err != nil {
		return err
	}

	// Setting Gin Logger
	if common.Config.EnableGinFileLog {
		f, _ := os.Create("logs/gin.log")
		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	} else {
		if !common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	m.router = gin.Default()

	return nil
}

// @title Inventory Service API Document
// @version 1.0
// @description List APIs of Store Service
// @termsOfService http://swagger.io/terms/

// @host localhost:8805
// @BasePath /api/v1
func main() {
	m := Main{}

	// Swagger Docs
	docs.SwaggerInfo.Title = "Inventory Service API Document"
	docs.SwaggerInfo.Description = "Documenation for Inventory Service."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8801"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Initialize server
	if m.initServer() != nil {
		return
	}

	// Close the connection once no longer needed
	db, err := databases.Database.DBConn.DB()
	if err != nil {
		panic("Failed to get database connection")
	}
	defer db.Close()

	itemController := controllers.Item{}

	// Simple group: v1
	v1 := m.router.Group("/api/v1")
	{
		items := v1.Group("/items")

		items.POST("", itemController.AddItem)
		items.GET("", itemController.ListItems)
		items.GET(":id", itemController.GetItemByID)
		items.DELETE("", itemController.DeleteItem)
		items.PUT("", itemController.UpdateItem)
	}

	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	m.router.Run(common.Config.Port)
}
