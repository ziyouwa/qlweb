package stat

import (
	"log"
	"qlweb/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/procfs"
)

func statInfo() (*procfs.Stat, error) {
	fs, err := procfs.NewFS("/proc")
	if err != nil {
		return nil, err
	}
	stat, err := fs.Stat()
	if err != nil {
		return nil, err
	}
	return &stat, nil
}

// ShowAllStatus get network infomation
func ShowAllStatus(c *gin.Context) {
	allinfo, err := statInfo()
	if err != nil {
		log.Fatal("Get Network information error: ", err.Error())
	}
	utils.Render(
		c,
		gin.H{
			"title":   "System Status",
			"Name":    "System base information",
			"payload": allinfo,
		},
		"stat.html",
	)
}
