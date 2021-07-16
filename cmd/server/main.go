package main

import (
	"flag"
	"os"

	"github.com/dbut2/shortener/internal/server"
	"github.com/dbut2/shortener/pkg/config"
)

var (
	configPath = flag.String("config", "config.yml", "")
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	c := config.Load(*configPath)
	err := server.Run(c)
	if err != nil {
		panic(err.Error())
	}
}
