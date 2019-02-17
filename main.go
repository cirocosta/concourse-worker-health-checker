package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/cirocosta/concourse-worker-health-checker/healthcheck"
)

var (
	gardenUrl       = flag.String("garden-url", "http://127.0.0.1:7777", "url of the garden server")
	baggageclaimUrl = flag.String("baggageclaim-url", "http://127.0.0.1:7788", "url of the baggageclaim server")
	timeout         = flag.Duration("timeout", 5*time.Second, "maximum amount of time to wait for checkers to run")
)

func main() {
	aggregate := &healthcheck.Aggregate{
		Checkers: []healthcheck.Checker{
			&healthcheck.Baggageclaim{Url: *baggageclaimUrl},
			&healthcheck.Garden{Url: *gardenUrl},
		},
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(*timeout))
	defer cancel()

	err := aggregate.Check(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("success!")
}
