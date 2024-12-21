package main

import (
	"html/template"
	"io"
	"kys/services"
	"log"
	"slices"

	//  "golang.org/x/exp/maps"
	"github.com/labstack/echo/v4"
)

func FirstElement(list []string) string {
  if len(list) > 0 {
    return list[0]
  }
  return ""
}


func main(){
	e := echo.New()
	e.Renderer = newTemplate()

	e.Static("/public", "public")

	e.GET("/", func(c echo.Context) error {
      result, err := db.QueryEntity()
      if err != nil {
        log.Printf("ERROR")
        return c.JSON(200, err)
    }

		return c.Render(200, "index", result)
	})

  e.GET("/data", func(c echo.Context) error {
			//log.Print("Data")
      result, err := db.QueryEntity()
      if err != nil {
        return c.JSON(200, err)
      }
			return c.JSON(200, result)

	})

	e.GET("/entities", func(c echo.Context) error {
    result, err := db.QueryEntity()
      if err != nil {
        return c.JSON(200, "DB down")
    }
			//log.Print("Entities")
			return c.Render(200, "entities", result)
	})


  e.GET("/entities/:entity", func(c echo.Context) error {
		query := c.Param("entity")
    result, err := db.QueryEntity()
      if err != nil {
        return c.JSON(200, "DB down")
    }

		for _, v := range result{
			if slices.Contains(v.Name, query) {
        log.Printf("PROFILE %s", v)
        return c.Render(200, "profile", v)
			}
		}
		return c.Render(200, "profile", nil)
	})
	e.Start(":8080")

}


type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates{
  funcMap := template.FuncMap{
		"firstElement": FirstElement,
	}
	return &Templates{
			templates: template.Must(template.New("").Funcs(funcMap).ParseGlob("*views/*.html")),
	}
}


