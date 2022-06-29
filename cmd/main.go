package main

import (
	"GolangLivePT01/golang_web_programming/app/config"
	"GolangLivePT01/golang_web_programming/app/routes"
)

func main() {
	cfg := config.GetInstance()
	routes.InitializeRoutes(cfg.GetGroup())
	cfg.GetEcho().Logger.Fatal(cfg.GetEcho().Start(":" + cfg.GetServicePort()))
}
