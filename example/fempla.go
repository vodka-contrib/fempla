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
	v.SetRenderer(fempla.Renderor())
	v.Static("/static", "./static")
	v.Get("/", func(self vodka.Context) error {
		data := make(map[string]interface{})
		data["oh"] = "no"
		data["name"] = "Insion Ng"
		self.Set("title", "你好，世界")
		self.SetStore(data)
		self.Set("oh", "yes")
		return self.Render(200, "index.html")
	})

	v.Run(fasthttp.New(":9000"))

}
