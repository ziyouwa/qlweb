package stat

import (
	"log"
	"qlweb/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/procfs"
)

func transInfo() (*procfs.NetDev, error) {
	fs, err := procfs.NewFS("/proc")
	if err != nil {
		return nil, err
	}
	stat, err := fs.NetDev()
	if err != nil {
		return nil, err
	}
	return &stat, nil
}

// ShowNetTransStatus get network infomation
func ShowNetTransStatus(c *gin.Context) {
	transinfo, err := transInfo()
	if err != nil {
		log.Fatal("Get Network information error: ", err.Error())
	}
	utils.Render(
		c,
		gin.H{
			"title":   "System Status",
			"Name":    "System base information",
			"payload": transinfo,
		},
		"stat.html",
	)
}
