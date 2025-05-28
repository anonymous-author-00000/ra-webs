package main

import (
	"os"

	"github.com/anonymous-author-00000/ctstream/monitor/sslmate"
	"github.com/anonymous-author-00000/ra-webs/monitor"
	browsernotifier "github.com/anonymous-author-00000/ra-webs/monitor/notifier/browser"
)

func main() {
	domain := os.Getenv("RA_WEBS_TA_DOMAIN")
	if domain == "" {
		panic("RA_WEBS_TA_DOMAIN is empty")
	}

	ct, err := sslmate.DefaultCTClient(domain)

	if err != nil {
		panic(err)
	}

	notifier, err := browsernotifier.Default()
	if err != nil {
		panic(err)
	}

	m, err := monitor.Default(ct, notifier)
	if err != nil {
		panic(err)
	}

	defer m.Close()

	m.Run()
}
