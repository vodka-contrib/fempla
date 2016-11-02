package fempla

import (
	"net/http"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/test"
	. "github.com/smartystreets/goconvey/convey"
)

func request(method, path string, e *vodka.Vodka) (int, string) {
	req := test.NewRequest(method, path, nil)
	rec := test.NewResponseRecorder()
	e.ServeHTTP(req, rec)
	return rec.Status(), rec.Body.String()
}

func TestRenderHtml(t *testing.T) {
	Convey("Render HTML", t, func() {
		e := vodka.New()
		r := Renderor(FemplaOption{
			Directory: "test",
		})
		e.SetRenderer(r)
		e.Get("/vodka", func() vodka.HandlerFunc {
			return func(self vodka.Context) error {
				return self.Render(http.StatusOK, "vodka.html")
			}
		}())
		status, body := request("GET", "/vodka", e)
		So(status, ShouldEqual, http.StatusOK)
		So(body, ShouldEqual, "<h1>Hello vodka</h1>")
	})

	Convey("Render HTML with Reload", t, func() {
		e := vodka.New()
		r := Renderor(FemplaOption{
			Directory: "test",
			Reload:    true,
		})
		e.SetRenderer(r)
		e.Get("/vodka", func() vodka.HandlerFunc {
			return func(self vodka.Context) error {
				return self.Render(http.StatusOK, "vodka.html")
			}
		}())
		status, body := request("GET", "/vodka", e)
		So(status, ShouldEqual, http.StatusOK)
		So(body, ShouldEqual, "<h1>Hello vodka</h1>")
	})

	Convey("Render HTML with Context", t, func() {
		e := vodka.New()
		r := Renderor(FemplaOption{
			Directory: "test",
		})
		e.SetRenderer(r)
		e.Get("/vodka", func() vodka.HandlerFunc {
			return func(self vodka.Context) error {
				self.Set("name", "vodka")
				return self.Render(http.StatusOK, "vodka_markup.html")
			}
		}())
		status, body := request("GET", "/vodka", e)
		So(status, ShouldEqual, http.StatusOK)
		So(body, ShouldEqual, "<h1>Hello, vodka</h1>")
	})

	Convey("Render HTML with Context and Reload", t, func() {
		e := vodka.New()
		r := Renderor(FemplaOption{
			Directory:  "test",
			Reload:     true,
			LeftDelim:  "{{",
			RightDelim: "}}",
		})
		e.SetRenderer(r)
		e.Get("/vodka", func() vodka.HandlerFunc {
			return func(self vodka.Context) error {
				self.Set("name", "vodka")
				return self.Render(http.StatusOK, "vodka_markup.html")
			}
		}())
		status, body := request("GET", "/vodka", e)
		So(status, ShouldEqual, http.StatusOK)
		So(body, ShouldEqual, "<h1>Hello, vodka</h1>")
	})
}

func ExampleRender() {
	e := vodka.New()
	r := Renderor()
	e.SetRenderer(r)
	e.Get("/", func() vodka.HandlerFunc {
		return func(self vodka.Context) error {
			self.Set("title", "你好，世界")

			// render ./templates/index.html file.
			return self.Render(200, "index.html")
		}
	}())
}
