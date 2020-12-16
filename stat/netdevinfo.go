package stat

import (
	"log"
	"qlweb/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/procfs/sysfs"
)

func netDevInfo() (sysfs.NetClass, error) {
	fs, err := sysfs.NewFS("/sys")
	if err != nil {
		return nil, err
	}
	stat, err := fs.NetClass()
	if err != nil {
		return nil, err
	}
	return stat, nil
}

// ShowNetDevStatus get network infomation
func ShowNetDevStatus(c *gin.Context) {
	netdevinfo, err := netDevInfo()
	if err != nil {
		log.Fatal("Get Network information error: ", err.Error())
	}
	utils.Render(
		c,
		gin.H{
			"title":   "System Status",
			"Name":    "System base information",
			"payload": netdevinfo,
		},
		"stat.html",
	)
}
