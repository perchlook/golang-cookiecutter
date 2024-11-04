package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/{{ cookiecutter.org_name }}/{{ cookiecutter.project_slug }}/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := app.InitApp(os.Getenv("DATABASE_URL"))

	serverHost, exists := os.LookupEnv("SERVER_HOST")
	if !exists {
		serverHost = "127.0.0.1:3000"
	}

	e.Logger.Fatal(e.Start(serverHost))
}
