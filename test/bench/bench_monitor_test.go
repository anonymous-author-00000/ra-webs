package benchtest

import (
	"os"
	"testing"
	"time"

	"github.com/anonymous-author-00000/ctstream/core"
	"github.com/anonymous-author-00000/ctstream/monitor/sslmate"
	"github.com/anonymous-author-00000/ra-webs/monitor"
	browsernotifier "github.com/anonymous-author-00000/ra-webs/monitor/notifier/browser"
)

func BenchmarkMonitorNext(b *testing.B) {
	core.DefaultEpochSleep = 0
	core.DefaultPullingSleep = 0

	domain := os.Getenv("RA_WEBS_TA_DOMAIN")
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

	b.Run("monitor", func(b *testing.B) {
		for b.Loop() {
			m.CT.Next(m.Monitor)

			m, err = monitor.Default(ct, notifier)
			if err != nil {
				panic(err)
			}

			b.StopTimer()
			time.Sleep(time.Second * 5)
			b.StartTimer()
		}
	})
}
