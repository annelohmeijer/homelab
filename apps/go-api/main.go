package main

import (
	"fmt"
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

type Vertex struct {
	X int
	Y int
}

var list = []Vertex{
	{X: 2, Y: 3}, {X: 4, Y: 5},
}

// Store godoc
// @router /store [get]
func getStore(g *gin.Context) {
	g.JSON(http.StatusOK, list)
}

// @Router /store [post]
func addStore(g *gin.Context) {
	var data Vertex

	// here the data is still empty, because we only initalize it
	fmt.Println("Posting data %s", &data)

	if err := g.BindJSON(&data); err != nil {
		// here we obtain the data from the json and return it

		// here the data is still empty, because we only initalize it
		fmt.Println("Posting data %s", &data)

		return
	}

	list = append(list, data)
	g.IndentedJSON(http.StatusCreated, data)
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")

	// healthGroup := v1.Group("")
	//
	{
		v1.GET("/store", getStore)
		v1.POST("/store", addStore)
	}

	r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
