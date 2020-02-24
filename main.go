package main

import (
	"github.com/Planck1858/yesya-coin/router"
)

func main() {
	r := router.New()

	err := r.Run()
	if err != nil {
		panic(err)
	}

	//bc := models.NewBlockchain()
	//spew.Dump(bc)
	//bc.NewBlock("my new second block")
	//bc.NewBlock("my new third block")
	//spew.Dump(bc)
}
