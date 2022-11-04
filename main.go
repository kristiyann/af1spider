package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logic := NewLogicImpl()

	nikeParams := GetProductFromNikeParams{
		Url:  os.Getenv("PRODUCT_LINK"),
		Size: os.Getenv("SIZE_US"),
	}

	err = logic.GetProductFromNike(nikeParams)
	if err != nil {
		log.Fatalf("error: " + err.Error())
	}
}
