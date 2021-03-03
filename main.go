package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"try-go/server"
)

func main() {
	var c server.Config

	if err := cleanenv.ReadConfig("conf/conf.yaml", &c); err != nil {
		processError(err)
	}
	log.Printf("Starting server...")
	log.Printf("Reading conf/conf.yaml values: %+v", c)
	s, _ := server.NewServer(&c)

	log.Fatal(s.Run(c.Server.Host, c.Server.Port))
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
