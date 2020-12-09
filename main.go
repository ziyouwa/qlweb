package main

import (
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {

}

//
func main() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	initRouters(r)

	r.Run()
}
