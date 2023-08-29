package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == ""{
		log.Fatal("NO PORT FOUND")
	}

	fmt.Print("PORT NO: ", port)
}