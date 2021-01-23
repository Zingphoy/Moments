package main

import (
	"Moments/pkg/log"
	"Moments/router"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func init() {
	log.InitLogger(false)
	//model.InitDatabase()()
}

var (
	g    errgroup.Group
	port = ":6666"
)

func startExpander() {

}

func main2() {
	server := &http.Server{
		Addr:         port,
		Handler:      router.InitRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
