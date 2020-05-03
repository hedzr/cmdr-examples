package svr

import (
	"crypto/tls"
	"github.com/hedzr/cmdr"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net"
	"net/http"
)

// https://echo.labstack.com/
// https://github.com/labstack/echo

func newEcho() *echoImpl {
	d := &echoImpl{}
	d.init()
	return d
}

type echoImpl struct {
	e *echo.Echo
}

func (d *echoImpl) Handler() http.Handler {
	// panic("implement me")
	return d.e
}

func (d *echoImpl) Serve(srv *http.Server, listener net.Listener, certFile, keyFile string) (err error) {
	// panic("implement me")
	// d.e.Logger.Fatal(d.e.Start(":1323"))

	if listener != nil {
		h2listener = tls.NewListener(listener, srv.TLSConfig)
		// d.e.Listener = h2listener
		d.e.TLSListener = h2listener
	}
	
	err = d.e.StartServer(srv)
	return
}

func (d *echoImpl) BuildRoutes() {
	// panic("implement me")

	d.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, echo World!")
	})
}

func (d *echoImpl) init() {
	d.e = echo.New()

	// https://echo.labstack.com/middleware/logger
	l := cmdr.GetLoggerLevel()
	n := log.DEBUG
	switch l {
	case cmdr.OffLevel:
		n = log.OFF
	case cmdr.FatalLevel, cmdr.PanicLevel:
		n = log.ERROR
	case cmdr.ErrorLevel:
		n = log.ERROR
	case cmdr.WarnLevel:
		n = log.WARN
	case cmdr.InfoLevel:
		n = log.INFO
	default:
		n = log.DEBUG
	}
	d.e.Logger.SetLevel(n)

	// d.e.Logger.Fatal(d.e.Start(":1323"))
}
