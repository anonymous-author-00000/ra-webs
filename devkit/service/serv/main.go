package main

import (
	"github.com/anonymous-author-00000/ra-webs/devkit/service"
	"github.com/anonymous-author-00000/ra-webs/devkit/service/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	g := e.Group("/api")

	l, err := service.Default()
	if err != nil {
		panic(err)
	}

	api.GetApi.Set(g, l)
	api.PostApi.Set(g, l)

	defer l.DB.Close()

	err = e.Start(":8080")

	if err != nil {
		panic(err)
	}
}
