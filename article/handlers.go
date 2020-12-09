package article

import (
	"net/http"
	"qlweb/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	utils.Render(
		c,
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
		"index.html",
	)
}

func getArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			utils.Render(
				c,
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
				"article.html",
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
