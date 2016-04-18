package frontend

import (
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"path/filepath"
)

func init() {
	router := gin.New()
	router.HTMLRender = loadTemplates()

	router.GET("/", indexHandler)

	http.Handle("/", router)
}

func indexHandler(c *gin.Context) {
	res := gin.H{"Default": 10}
	c.HTML(http.StatusOK, "index.tmpl", res)
}

func loadTemplates() multitemplate.Render {
	dir := "templates/"
	r := multitemplate.New()

	layouts, err := filepath.Glob(dir + "layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(dir + "includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	for _, layout := range layouts {
		files := append(includes, layout)
		r.Add(filepath.Base(layout), template.Must(template.ParseFiles(files...)))
	}
	return r
}
