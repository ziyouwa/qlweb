package stat

import (
	"log"
	"qlweb/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/procfs"
)

func loadavgStat() (*procfs.LoadAvg, error) {
	fs, err := procfs.NewFS("/proc")
	if err != nil {
		return nil, err
	}
	stats, err := fs.LoadAvg()
	if err != nil {
		return nil, err
	}
	return stats, nil

}

// ShowLoadavgStatus get memory infomation
func ShowLoadavgStatus(c *gin.Context) {
	loadavginfo, err := loadavgStat()
	if err != nil {
		log.Fatal("Get Loadavg information error: ", err.Error())
	}
	utils.Render(
		c,
		gin.H{
			"title":   "System Status",
			"Name":    "System base information",
			"payload": loadavginfo,
		},
		"stat.html",
	)
}
