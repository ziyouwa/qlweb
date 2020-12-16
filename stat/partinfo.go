package stat

import (
	"log"
	"qlweb/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/procfs/blockdevice"
)

func diskInfo() ([]blockdevice.Diskstats, error) {
	fs, err := blockdevice.NewFS("/proc", "/sys")
	if err != nil {
		return []blockdevice.Diskstats{}, err
	}
	stat, err := fs.ProcDiskstats()
	if err != nil {
		return []blockdevice.Diskstats{}, err
	}
	return stat, nil
}

// ShowDiskStatus get partitions infomation
func ShowDiskStatus(c *gin.Context) {
	diskinfo, err := diskInfo()
	if err != nil {
		log.Fatal("Get Disk information error: ", err.Error())
	}
	utils.Render(
		c,
		gin.H{
			"title":   "System Status",
			"Name":    "System base information",
			"payload": diskinfo,
		},
		"stat.html",
	)
}
