package main

import (
	"GolangLivePT01/golang_web_programming/config"
	"GolangLivePT01/golang_web_programming/routes"
)

func main() {
	cfg := config.GetInstance()
	e := routes.InitializeRoutes(cfg)
	e.Logger.Fatal(e.Start(":" + cfg.GetServicePort()))
}
