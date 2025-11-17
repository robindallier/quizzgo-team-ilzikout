package main

import (
	"log"

	"github.com/robindallier/quizzgo-team-ilzikout/web"
)

func main() {
	if err := web.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
