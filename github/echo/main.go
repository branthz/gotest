package main

import (
  "net/http"
	"html/template"
	"io"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

// Handler
func hello(c echo.Context) error {
  	return c.String(http.StatusOK, "Hello, World!")
}

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w,name,data)
}

type info struct {
	Name string
	Age int
	Data string
}

func Hello(c echo.Context) error {
	f:=&info{
		"jingxiao",
		13,
		`<a href="https://www.google.com">google</a>`,
	}
    return c.Render(http.StatusOK, "hello", f)
}

func Unescape(a string) template.HTML{
	return template.HTML(a)
}

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  //e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

  fcs:=make(template.FuncMap,0)
  fcs["noescape"]=Unescape
  t := &Template{
    templates: template.Must(template.New("foo").Funcs(fcs).ParseGlob("./public/views/*.html")),
  }
  e.Renderer = t
  e.GET("/hello", Hello)
  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}
