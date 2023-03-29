package main

import (
	"errors"
	"github.com/ichtrojan/go-todo/routes"
	"github.com/ichtrojan/thoth"
	"log"
	"net/http"
	"os"
)

func main() {
	logger, _ := thoth.Init("log")

	port, exist := os.LookupEnv("PORT")

	if !exist {
		logger.Log(errors.New("PORT not set in .env"))
		log.Fatal("PORT not set in .env")
	}

	err := http.ListenAndServe(":"+port, routes.Init())

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}
}
