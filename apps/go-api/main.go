package main

import (
	"api/internal/container"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	docs "api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Project Management API
// @version 1.0
// @description A simple CRUD API for managing projects
// @host localhost:8080
// @BasePath /api/v1

// healthCheck godoc
// @Summary Health check
// @Description Check if the API is running
// @Tags health
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":    "OK",
		"timestamp": time.Now(),
	})
}

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/homelab?sslmode=disable"
	}

	ctx := context.Background()
	c, err := container.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}
	defer c.Close()

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/health", healthCheck)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("/api/v1")
	{
		projects := v1.Group("/projects")
		{
			projects.POST("", c.ProjectHandler.CreateProject)
			projects.GET("", c.ProjectHandler.GetAllProjects)
			projects.GET("/:id", c.ProjectHandler.GetProject)
			projects.PUT("/:id", c.ProjectHandler.UpdateProject)
			projects.DELETE("/:id", c.ProjectHandler.DeleteProject)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
