package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/icco/cron/stats"
	"github.com/icco/unifi/metrics"
	log "github.com/sirupsen/logrus"
	"github.com/unifi-poller/unifi"
)

var (
	host  = flag.String("host", "unifi", "Controller hostname")
	user  = flag.String("user", "", "Controller username")
	pass  = flag.String("pass", "", "Controller password")
	port  = flag.Int("port", 8443, "Controller port")
	token = flag.String("token", "", "Graphql Token")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	c := &unifi.Config{
		User:     *user,
		Pass:     *pass,
		URL:      fmt.Sprintf("https://%s:%d/", *host, *port),
		ErrorLog: log.Printf,
		DebugLog: log.Printf,
	}

	u, err := unifi.NewUnifi(c)
	if err != nil {
		log.Fatal(err)
	}

	v, err := metrics.GetClients(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%f clients found", v)

	sc := &stats.Config{
		Log:          log,
		GraphQLToken: *token,
	}
	if err := sc.UploadStat(ctx, "Network Clients", v); err != nil {
		log.Fatal(err)
	}
}
