package examples

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ShowUploadFile() error {
	r := gin.New()

	// Set a lower memory limit for multipart forms (default is 32MB)
	r.MaxMultipartMemory = 8 << 20 // 8MB

	// curl -X POST http://localhost:8101/upload \
	//  -F "upload[]=@/Users/litong.deng/Downloads/aws-lambda-functions.png" \
	//  -F "upload[]=@/Users/litong.deng/Downloads/aws-textract-project.png"
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./upload/%v", file.Filename)
			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			}
		}
	})

	return r.Run(":8101")
}
