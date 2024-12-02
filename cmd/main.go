package main

import (
	"html/template"
	"io"
	"kys/services"
	"log"
	"reflect"
//  "golang.org/x/exp/maps"
	"github.com/labstack/echo/v4"
)

func StructToMap(data interface{}) map[string]interface{}{
  result := make(map[string]interface{})
  val := reflect.ValueOf(data)
  typ := reflect.TypeOf(data)
  for i:= 0; i < val.NumField(); i++ {
    field := typ.Field(i)
    tag := field.Tag.Get("human")
    if tag == ""{
      tag = field.Tag.Get("json")
    }
    if tag == "" {
      tag = field.Name
    }
    result[tag] = val.Field(i).Interface()
    log.Print(result[tag])
  }
  return result
}

func main(){
	e := echo.New()
	e.Renderer = newTemplate()

	e.Static("/public", "public")

	e.GET("/", func(c echo.Context) error {
      result, err := db.QueryEntity()
      if err != nil {
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
    result, err := db.QueryEntity()
      if err != nil {
        return c.JSON(200, "DB down")
    }
		query := c.Param("entity")
		for _, v := range result{
			if v.Name == query {
        entityMap := StructToMap(v)
        //keys := maps.Keys(entity)
        keys := []string{"Date of Birth", "Date of Death", "Gender", "Nationality", "Citizenship", "Birthplace", "Birth Country", "Passport Number", "Social Security Number", "Phone", "Source Name", "Source"}
        data := map[string]interface{}{
          "Map": entityMap,
          "Keys": keys,
        }

			 return c.Render(200, "profile", data)
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
	return &Templates{
			templates: template.Must(template.ParseGlob("*views/*.html")),
	}
}


