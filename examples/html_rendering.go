package examples

import (
	"gindemo/templates"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowHtmlRendering() error {
	router := gin.New()
	t, err := templates.LoadTemplates()
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(t)

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/index.html", gin.H{
			"Foo": "Main website",
		})
	})
	router.GET("/bar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/bar.html", gin.H{
			"Foo": "Main website",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/users/index.html", gin.H{
			"Foo": "Users Home",
		})
	})

	return router.Run(":8097")
}
