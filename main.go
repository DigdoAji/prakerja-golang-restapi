package main

import (
	"os"
	"ujk-golang/configs"
	"ujk-golang/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)



func main(){
	loadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":3030"
	}
	return ":" + port
}

func loadEnv(){
	err := godotenv.Load()
  	if err != nil {
    	panic("Error loading .env file")
  	}
}

