package stat

import "github.com/gin-gonic/gin"

// InitRouters init routers
func InitRouters(r *gin.Engine) {
	r.GET("/stat", showStatus)
}
