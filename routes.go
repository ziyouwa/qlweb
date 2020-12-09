package main

import (
	"qlweb/article"
	"qlweb/stat"

	"github.com/gin-gonic/gin"
)

func initRouters(r *gin.Engine) {
	article.InitRouters(r)
	stat.InitRouters(r)
}
