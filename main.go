package main

import (
	"Moments/biz/mq"
	"Moments/pkg/log"
	"Moments/router"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	g    errgroup.Group
	port = ":9999"
)

func main() {
	log.InitLogger(false)
	//mq.InitMQ()
	//mq.InitExpander()
	defer mq.StopMQ()

	server := &http.Server{
		Addr:         port,
		Handler:      router.InitRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(nil, err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		panic(err)
	}
}
