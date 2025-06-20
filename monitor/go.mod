module github.com/anonymous-author-00000/ra-webs/monitor

go 1.23.6

replace github.com/anonymous-author-00000/crtsh => ../thirdparty/crtsh

replace github.com/anonymous-author-00000/extract-embed => ../thirdparty/extract-embed

replace github.com/anonymous-author-00000/go-utils => ../thirdparty/go-utils

replace github.com/anonymous-author-00000/golang-utils => ../thirdparty/golang-utils

replace github.com/anonymous-author-00000/ra-webs/devkit/service => ../devkit/service

replace github.com/anonymous-author-00000/ra-webs/monitor/serviceclient => ./serviceclient

replace github.com/anonymous-author-00000/ctstream => ../thirdparty/ctstream

replace github.com/anonymous-author-00000/sslmate-cert-search-api => ../thirdparty/sslmate-cert-search-api/

require (
	entgo.io/ent v0.14.4
	github.com/SherClockHolmes/webpush-go v1.4.0
	github.com/anonymous-author-00000/ctstream v0.0.39
	github.com/anonymous-author-00000/extract-embed v0.0.1
	github.com/anonymous-author-00000/go-utils v0.0.7
	github.com/anonymous-author-00000/golang-utils v0.0.1
	github.com/anonymous-author-00000/ra-webs/devkit/core v0.0.0-00010101000000-000000000000
	github.com/anonymous-author-00000/ra-webs/devkit/service v0.0.0-00010101000000-000000000000
	github.com/anonymous-author-00000/ra-webs/monitor/serviceclient v0.0.0-00010101000000-000000000000
	github.com/cockroachdb/errors v1.11.3
	github.com/edgelesssys/ego v1.5.0
	github.com/google/certificate-transparency-go v1.2.1
	github.com/labstack/echo/v4 v4.13.3
	github.com/mattn/go-sqlite3 v1.14.22
)

require (
	ariga.io/atlas v0.31.1-0.20250212144724-069be8033e83 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/anonymous-author-00000/sslmate-cert-search-api v0.0.15 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/getsentry/sentry-go v0.27.0 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/zclconf/go-cty v1.14.4 // indirect
	github.com/zclconf/go-cty-yaml v1.1.0 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)

replace github.com/anonymous-author-00000/ra-webs/devkit/core => ../devkit/core
