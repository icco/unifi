package main

import (
	"context"
	"flag"
	"log"

	"github.com/dim13/unifi"
	"github.com/icco/unifi/metrics"
)

var (
	host = flag.String("host", "unifi", "Controller hostname")
	user = flag.String("user", "", "Controller username")
	pass = flag.String("pass", "", "Controller password")
	port = flag.String("port", "8443", "Controller port")
)

func main() {
	ctx := context.Background()

	u, err := unifi.Login(*user, *pass, *host, *port, "", 6)
	if err != nil {
		log.Fatal(err)
	}

	v, err := metrics.GetClients(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%d", v)
}