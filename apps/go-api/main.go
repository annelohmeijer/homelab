package main

import (
	"net/http"
	"time"

	docs "api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v2

// Health godoc
// @Summary Health check
// @Description Check if the API is running
// @Tags health
// @Produce json
// @Success 200 {string} string "OK"
// @Router /health [get]
func Health(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"status":    "OK",
		"timestamp": time.Now(),
	})
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")

	// healthGroup := v1.Group("")

	{
		v1.GET("/health", Health)
	}

	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}
	r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
