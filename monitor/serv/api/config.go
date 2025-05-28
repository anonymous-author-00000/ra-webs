package api

import (
	"net/http"

	goutils "github.com/anonymous-author-00000/go-utils"
	"github.com/anonymous-author-00000/ra-webs/monitor/serv"
	"github.com/labstack/echo/v4"
)

var GetConfigApi = goutils.EchoRoute[serv.MonitorServer]{
	Method: goutils.GET,
	Path:   "/config",
	F: func(server *serv.MonitorServer) goutils.EchoRouteFunc {
		return func(c echo.Context) error {

			return c.JSON(http.StatusOK, map[string]string{
				"taDomain": server.Monitor.TADomain,
			})
		}
	},
}
