package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": msg,
		"data":    data,
	})

	//c.Abort()

}

func HTMLResponse(c *gin.Context, htmlFile string, data interface{}) {
	c.HTML(200, htmlFile, data)

}

func RedirectResponse(c *gin.Context, location string) {
	c.Redirect(http.StatusFound, location)
}
