package main

import (
	"go-redis-sample/pkg/timeline"
	"io"
	"net/http"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func formatDate(unix int64) string {
	t := time.Unix(unix, 0)
	return t.Format("2006/04/01 15:04")
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	renderer := &TemplateRenderer{
		templates: template.Must(template.New("").Funcs(template.FuncMap{
			"formatDate": formatDate,
		}).ParseGlob("./templates/*.html")),
	}
	e.Renderer = renderer
	ctl := timeline.NewTimelineController()

	e.GET("/timeline", ctl.GetTimeline)
	e.GET("/", func(c echo.Context) error { return c.JSON(http.StatusOK, "Hello World!!") })

	e.Logger.Fatal(e.Start(":8080"))
}
