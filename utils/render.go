package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Render render result
func Render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	case "text/html":
		c.HTML(http.StatusOK, templateName, data)
	default:
		c.JSON(http.StatusOK, data["payload"])
	}
}
