package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	WEATHER_API_TOKEN := os.Getenv("WEATHER_API_TOKEN")
	if WEATHER_API_TOKEN == "" {
		log.Fatal("Missing WEATHER_API_TOKEN")
	}

	r := gin.Default()

	r.GET("/weather", func(c *gin.Context) {
		location := c.Query("location")
		if location == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing location query parameter"})
			return
		}
		url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=yes", WEATHER_API_TOKEN, location)

		response, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Weather API"})
			return
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Weather API"})
		}

		c.Data(response.StatusCode, "application/json", body)
	})

	r.Run(":8080")
}
