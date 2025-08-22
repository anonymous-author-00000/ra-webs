package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"

	"github.com/anonymous-author-00000/ra-webs/devkit/core"
	"github.com/anonymous-author-00000/ra-webs/monitor/serv"
	"github.com/anonymous-author-00000/ra-webs/monitor/serv/api"

	browsernotify "github.com/anonymous-author-00000/ra-webs/monitor/notifier/browser"
	"github.com/labstack/echo/v4"
)

//go:embed views/*.html
var viewEmbedFiles embed.FS

//go:embed static/js/*.js static/js/*/*.js static/*.json
var staticEmbedFiles embed.FS

const TMP_FOLDER_NAME = "views"

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func InjectSWHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Service-Worker-Allowed", "/")
		return next(c)
	}
}

func main() {
	// debug.EnableDebug()

	e := echo.New()
	apiGroup := e.Group(core.API_ROOT)

	s, err := serv.Default()
	if err != nil {
		panic(err)
	}

	defer s.Close()

	api.Route(apiGroup, s)

	viewSubFS := echo.MustSubFS(viewEmbedFiles, "views")
	e.StaticFS("/", viewSubFS)

	staticSubFS := echo.MustSubFS(staticEmbedFiles, "static")
	e.StaticFS("/static", staticSubFS)

	// e.Use(middleware.Logger())
	e.Use(InjectSWHeader)

	s.Monitor.Notifier.(*browsernotify.BrowserNotifier).Setup(s.Monitor, apiGroup)
	e.Debug = true
	fmt.Printf("public: %v\nprivate: %v",
		s.Monitor.Notifier.(*browsernotify.BrowserNotifier).VapidPublicKey,
		s.Monitor.Notifier.(*browsernotify.BrowserNotifier).VapidPrivateKey)

	err = s.Run(":8000", e)
	e.Logger.Fatal(err)
}
