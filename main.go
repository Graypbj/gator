package main

import (
	"fmt"
	"log"

	"github.com/Graypbj/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("grayson")
	if err != nil {
		log.Fatal(err)
	}

	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", *updatedCfg)
}
