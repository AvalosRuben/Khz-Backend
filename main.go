package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	API_URL := os.Getenv("API_URL")
	API_KEY := os.Getenv("API_KEY")

	client, err := supabase.NewClient(API_URL, API_KEY, &supabase.ClientOptions{})
	if (err != nil){
		fmt.Println("Failed to initialize the db client:" , err)
	} else{
		// I'm using client here rn ONLY to avoid the error of not using it elsewhere.
		//I'll change it later
		fmt.Println("Connected to the db successfully: ", client)
	}

	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"Welcome to the Khz api!",
		})
	})

	

	r.Run()
}
