package main

import (
	"net/http"
	"time"

	"github.com/aruna-cande/europeanaSearch"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://192.168.31.13:8081"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://192.168.31.13:8080"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/search", handleSearchRequest)

	r.Run()
}

func handleSearchRequest(c *gin.Context) {
	//var chrs search.Service
	client := &http.Client{}
	chrs := search.NewCulturalHeritageRecordService()

	var searchData search.Search

	if err := c.Bind(&searchData); err == nil {
		response := chrs.SearchCulturalHeritageRecords(client, searchData)
		//fmt.Printf("response %s", response)
		c.JSON(200, response)
	}

	if err := c.BindJSON(&searchData); err == nil {
		response := chrs.SearchCulturalHeritageRecords(client, searchData)
		//fmt.Printf(response)
		c.JSON(200, response)
	}
}
