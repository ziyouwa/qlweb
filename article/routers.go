package article

import (
	"github.com/gin-gonic/gin"
)

// InitRouters init routers
func InitRouters(r *gin.Engine) {

	r.GET("/article", showIndexPage)

	r.GET("/article/view/:id", getArticle)

}
