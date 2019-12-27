package main

import (
	"log"
	"time"

	"github.com/gotoolkit/config"
)

func main() {
	c, err := config.New(config.WithFile(".env"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(c.GetString("test"))
	time.Sleep(5 * time.Second)
	log.Println(c.GetString("test"))
}
