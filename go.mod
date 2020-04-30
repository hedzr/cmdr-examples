module github.com/hedzr/cmdr-examples

go 1.14

// replace github.com/hedzr/logex => ../logex

replace github.com/hedzr/cmdr => ../cmdr

replace github.com/hedzr/cmdr-addons => ../cmdr-addons

replace github.com/kardianos/service => ../../kardianos/service

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/gin-gonic/gin v1.6.2
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gorilla/mux v1.7.4
	github.com/hedzr/cmdr v1.6.35
	github.com/hedzr/logex v1.1.8
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kardianos/service v1.0.0
	github.com/kataras/iris/v12 v12.1.8
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/valyala/fasthttp v1.12.0 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180219094614-024efa345fa9 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073
	gopkg.in/hedzr/errors.v2 v2.0.11
)
