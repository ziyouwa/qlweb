package stat

import (
	"log"
	"qlweb/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/procfs"
)

// type CPUInfo procfs.CPUInfo

func cpuStat() ([]procfs.CPUInfo, error) {
	fs, err := procfs.NewFS("/proc")
	if err != nil {
		return []procfs.CPUInfo{}, err
	}
	stats, err := fs.CPUInfo()
	if err != nil {
		return []procfs.CPUInfo{}, err
	}
	return stats, nil

}

// ShowCPUStatus get cpu infomation
func ShowCPUStatus(c *gin.Context) {
	cpuinfo, err := cpuStat()
	if err != nil {
		log.Fatal("Get CPU info Error: ", err.Error())
	}
	utils.Render(
		c,
		gin.H{
			"title":   "System Status",
			"Name":    "System base information",
			"payload": cpuinfo[0],
		},
		"stat.html",
	)
}
