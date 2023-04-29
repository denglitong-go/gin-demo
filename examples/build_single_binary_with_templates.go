package examples

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

// Prepare Packages
// 	go get -u github.com/jessevdk/go-assets-builder
// 	go install github.com/jessevdk/go-assets-builder
// Generate asserts.go:
// 	cd examples && go-assets-builder -p examples ./html -o ./asserts.go
// Build the single file server
// 	go build -o assets-in-binary
// Run the single file server
// 	./assets-in-binary

func ShowBuildSingleBinaryWithAssertTemplate() error {
	r := gin.New()
	t, err := loadTemplates()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.GET("/", serveIndexPage)
	r.GET("/bar", serveBarPage)

	return r.Run(":8086")
}

func serveIndexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "/html/index.html", gin.H{
		"Foo": "World",
	})
}

func serveBarPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "/html/bar.html", gin.H{
		"Bar": "World",
	})
}

func loadTemplates() (*template.Template, error) {
	t := template.New("default")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
