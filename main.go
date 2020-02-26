package main

import (
	"github.com/Planck1858/yesya-coin/router"
	"log"
)

func main() {
	r := router.New()

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
