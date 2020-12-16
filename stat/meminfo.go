package stat

import (
	"log"
	"qlweb/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/procfs"
)

func memStat() (*procfs.Meminfo, error) {
	fs, err := procfs.NewFS("/proc")
	if err != nil {
		return nil, err
	}
	stats, err := fs.Meminfo()
	if err != nil {
		return nil, err
	}
	return &stats, nil

}

// ShowMemStatus get memory infomation
func ShowMemStatus(c *gin.Context) {
	meminfo, err := memStat()
	if err != nil {
		log.Fatal("Get memory information Error: ", err.Error())
	}
	utils.Render(
		c,
		gin.H{
			"title":   "System Status",
			"Name":    "System base information",
			"payload": meminfo,
		},
		"stat.html",
	)
}
