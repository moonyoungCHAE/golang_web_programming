package main

import "golang_web_programming/internal"

func main() {
	server := internal.NewDefaultServer()
	server.Run()
}
