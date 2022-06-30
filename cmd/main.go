package main

import (
	"github.com/boldfaced7/golang_web_programming/app"
	"log"
)

func main() {
	log.Fatal(app.NewEcho(*app.DefaultConfig()).Start(":8080"))
}
