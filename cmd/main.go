package main

import (
	"github.com/Nikaprilya/test"

	"log"
)

func main() {
	srv := new(test.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
