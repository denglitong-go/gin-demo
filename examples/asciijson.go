package examples

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ShowAsciiJSON() {
	r := gin.Default()

	r.GET("/someJSON", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "<br/>",
		}
		context.AsciiJSON(http.StatusOK, data)
	})

	log.Fatalln(r.Run())
}
