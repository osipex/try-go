package main

import (
	"log"
	"try-go/server"
)

func main() {
	c, err := readConfiguration()
	if err != nil {
		log.Fatal(err)
	}
	s := server.NewServer(c)
	log.Printf("Starting a fresh server...")
	log.Fatal(s.Run())
}

func readConfiguration() (*server.Config, error) {
	c := &server.Config{}

	err := readYaml(c)
	if err != nil {
		return nil, err
	}

	readEnv(c)

	return c, nil
}

func readYaml(c *server.Config) error {
	// TODO: IMPLEMENT PART C
	return nil
}

func readEnv(c *server.Config) {
	// TODO: IMPLEMENT PART C
}
