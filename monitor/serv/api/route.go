package api

import (
	"github.com/anonymous-author-00000/ra-webs/monitor/serv"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Group, monitor *serv.MonitorServer) {
	GetLogs.Set(e, monitor)
	PostNotifierApi.Set(e, monitor)
	GetConfigApi.Set(e, monitor)
}
