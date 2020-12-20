package main

import (
	"Moments/pkg/log"
	"Moments/router"
)

func init() {
	log.InitLogger(false)
}

func startExpander() {

}

func main() {
	r := router.InitRouter()
	err := r.Run(":6666")
	if err != nil {
		log.Fatal(err)
	}
}
