package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/unpoller/unifi"
	"go.icco.me/gutil/logging"
	"go.icco.me/unifi/metrics"
)

var (
	host = flag.String("host", "unifi", "Controller hostname")
	user = flag.String("user", "", "Controller username")
	pass = flag.String("pass", "", "Controller password")
	port = flag.Int("port", 8443, "Controller port")
)

func main() {
	flag.Parse()
	log := logging.Must(logging.NewLogger("unifi"))

	ctx := context.Background()
	c := &unifi.Config{
		User:     *user,
		Pass:     *pass,
		URL:      fmt.Sprintf("https://%s:%d/", *host, *port),
		ErrorLog: log.Errorf,
		DebugLog: log.Debugf,
	}

	u, err := unifi.NewUnifi(c)
	if err != nil {
		log.Fatal(err)
	}

	v, err := metrics.GetClients(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("%f clients found", v)

	n, err := metrics.GetBytesPerSecond(ctx, u)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("WAN: %+v", n)

	bytesPerMb := 125000.0
	log.Infof("WAN TX: %.2f mbps", n.Upload/bytesPerMb)
	log.Infof("WAN RX: %.2f mbps", n.Download/bytesPerMb)
}
