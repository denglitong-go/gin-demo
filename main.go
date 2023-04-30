package main

import (
	"gindemo/examples"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
)

var (
	g errgroup.Group
)

func showQuickStart() error {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// listen and serve on 0.0.0.0:8080
	return router.Run(":8080")
}

func main() {
	// Using sync/errgroup + router to listen on multiple ports for multiple routes
	g.Go(showQuickStart)
	g.Go(examples.ShowAsciiJSON)
	g.Go(examples.ShowBindDataRequestWithCustomStruct)
	g.Go(examples.ShowBindHtmlCheckBoxes)
	g.Go(examples.ShowBindQueryStringOrPostData)
	g.Go(examples.ShowBindUri)
	g.Go(examples.ShowBuildSingleBinaryWithAssertTemplate)
	g.Go(examples.ShowControlLogOutputColor)

	if err := g.Wait(); err != nil {
		log.Fatalln(err)
	}
}
