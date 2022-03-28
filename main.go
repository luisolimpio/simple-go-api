package main

import (
	"github.com/joho/godotenv"
	"github.com/luisolimpio/api-go-gin/database"
	"github.com/luisolimpio/api-go-gin/routes"
)

func init() {
	godotenv.Load()
}

func main() {
	database.Connection()

	routes.Router()
}
