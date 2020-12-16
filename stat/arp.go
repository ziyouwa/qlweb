package stat

import (
	"log"
	"qlweb/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/procfs"
)

func arpStat() ([]procfs.ARPEntry, error) {
	fs, _ := procfs.NewFS("/proc")
	stats, err := fs.GatherARPEntries()
	if err != nil {
		return []procfs.ARPEntry{}, err
	}
	return stats, nil

}

// ShowArpStatus get arp infomation
func ShowArpStatus(c *gin.Context) {
	arpinfo, err := arpStat()
	if err != nil {
		log.Fatal("Get CPU info Error: ", err.Error())
	}
	utils.Render(
		c,
		gin.H{
			"title":   "System Status",
			"Name":    "System base information",
			"payload": arpinfo,
		},
		"stat.html",
	)
}
