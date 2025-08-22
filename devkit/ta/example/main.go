package main

import (
	"fmt"
	"net/http"

	"github.com/anonymous-author-00000/ra-webs/devkit/ta"
	// "github.com/anonymous-author-00000/ra-webs/devkit/core/attest/debug"
)

func main() {
	// debug.EnableDebug()

	config, err := ta.DefaultConfig()
	if err != nil {
		panic(err)
	}

	ta, err := ta.Init(config)
	if err != nil {
		panic(err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		fmt.Fprintln(w, "Hello from TA running on TEE :)<br/>")

		for _, m := range config.MonitorBases {
			fmt.Fprintf(w, `<button onclick="window.open('%s');">Monitor Page (%s)</button><br/>`, m, m)
		}
	}

	tlsConfig, err := ta.TLSConfig()
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr:      ":443",
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/", handler)

	err = server.ListenAndServeTLS("", "")
	panic(err)
}
