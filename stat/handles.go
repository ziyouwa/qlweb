package stat

import (
	"qlweb/utils"

	"github.com/gin-gonic/gin"
)

func showStatus(c *gin.Context) {
	stat, _ := Stat()
	utils.Render(
		c,
		gin.H{
			"title":   "System Status",
			"payload": stat,
		},
		"stat.html",
	)
}
