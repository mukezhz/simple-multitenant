package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sachin-gautam/gin-api/database"
	"github.com/sachin-gautam/gin-api/handler"
	"github.com/sachin-gautam/gin-api/middleware"
)

func main() {
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	db := database.NewDatabase()
	handler := handler.NewHandler(db)
	middleware := middleware.NewTenantMiddleware(db)
	router.LoadHTMLFiles("index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	api := router.Group("/", middleware.ExtractTenantIDFromDomain())
	api.GET("/details", handler.GetDetail)
	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
