package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//go mod init x-clone/backend

func main() {
	err := godotenv.Load()
	apiEnv := os.Getenv("ENV")
	if err != nil && apiEnv == "" {
		log.Println("fail to load env", err)
	}

}
