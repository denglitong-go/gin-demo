package main

import (
	"gindemo/examples"
	"github.com/gin-gonic/gin"
	"log"
)

func showQuickStart() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// listen and serve on 0.0.0.0:8080
	log.Fatalln(router.Run())
}

func main() {
	// showQuickStart()
	// examples.ShowAsciiJSON()
	// examples.ShowBindDataRequestWithCustomStruct()
	// examples.ShowBindHtmlCheckBoxes()
	examples.ShowBindQueryStringOrPostData()
}
