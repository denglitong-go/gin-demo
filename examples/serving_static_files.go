package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowServingStaticFiles() error {
	r := gin.New()
	// can reflect the newest files under the dir
	r.StaticFS("/", http.Dir("./"))

	return r.Run(":8100")
}
