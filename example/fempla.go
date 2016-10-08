package main

import (
	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/fasthttp"
	"github.com/insionng/vodka/middleware"
	"github.com/vodka-contrib/fempla"
)

func main() {

	v := vodka.New()
	v.Use(middleware.Logger())
	v.Use(middleware.Recover())
	r := fempla.Renderor()
	v.SetRenderer(r)
	v.Static("/static", "./static")
	v.Get("/", func(ctx vodka.Context) error {
		ctx.Render(200, "index.html", map[string]interface{}{
			"title": "你好，世界",
		})
		return nil
	})

	v.Run(fasthttp.New(":9000"))

}
