package stat

import (
	"log"
	"qlweb/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/procfs"
)

func mountInfo() ([]*procfs.MountInfo, error) {
	stat, err := procfs.GetProcMounts(1)
	if err != nil {
		return []*procfs.MountInfo{}, err
	}
	return stat, nil
}

// ShowMountStatus get mounted filesystem infomation
func ShowMountStatus(c *gin.Context) {
	mountinfo, err := mountInfo()
	if err != nil {
		log.Fatal("Get Network information error: ", err.Error())
	}
	utils.Render(
		c,
		gin.H{
			"title":   "System Status",
			"Name":    "System base information",
			"payload": mountinfo,
		},
		"stat.html",
	)
}
