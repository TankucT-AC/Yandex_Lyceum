package main

import (
	"calculator/server"
)

func main() {
	app := server.New()
	app.RunServer()
}
