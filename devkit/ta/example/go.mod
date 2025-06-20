module github.com/anonymous-author-00000/ra-webs/monitor/example

go 1.21.4

replace github.com/anonymous-author-00000/ra-webs/ta => ../

replace github.com/anonymous-author-00000/ra-webs/devkit/core => ../../core

require (
	github.com/anonymous-author-00000/ra-webs/devkit/core v0.0.0-00010101000000-000000000000
	github.com/anonymous-author-00000/ra-webs/ta v0.0.0-00010101000000-000000000000
)

require (
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/edgelesssys/ego v1.5.0 // indirect
	github.com/go-acme/lego/v4 v4.16.1 // indirect
	github.com/go-jose/go-jose/v4 v4.0.2 // indirect
	github.com/miekg/dns v1.1.58 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)
