package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowServingStaticFiles() error {
	r := gin.New()
	r.StaticFS("/", http.Dir("./"))

	return r.Run(":8100")
}
