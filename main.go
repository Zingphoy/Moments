package main

import (
	"Moments/pkg/log"
	"Moments/routers"
)

func init() {
	log.InitLogger(false)
}

func startExpander() {

}

func main() {
	r := routers.InitRouter()
	err := r.Run(":6666")
	if err != nil {
		log.Fatal(err)
	}
}
