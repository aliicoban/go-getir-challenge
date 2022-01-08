package main

import (
	config "github.com/alicobanserver/config"
	LOG "github.com/alicobanserver/log"
	routes "github.com/alicobanserver/routes"
)

func main() {
	config.Connect()
	LOG.InitLog()

	routes.Routes()
}
