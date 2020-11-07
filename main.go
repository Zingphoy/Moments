package main

import (
	"Moments/models"
	"Moments/pkg/log"
	"Moments/routers"
)

func init() {
	log.InitLogger(false)
	models.SetUp()
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
