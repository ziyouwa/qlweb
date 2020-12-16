package stat

import "github.com/gin-gonic/gin"

// InitRouters init routers
func InitRouters(r *gin.Engine) {
	r.GET("/stat", ShowAllStatus)
	statGroup := r.Group("/stat")
	{
		statGroup.GET("/base", ShowBaseInfo)
		statGroup.GET("/cpu", ShowCPUStatus)
		statGroup.GET("/mem", ShowMemStatus)
		statGroup.GET("/loadavg", ShowLoadavgStatus)
		statGroup.GET("/disk", ShowDiskStatus)
		// Get Mounted Filesystem information
		statGroup.GET("/mount", ShowMountStatus)
		// Get Ethernet dev name and IP address.
		statGroup.GET("/arp", ShowArpStatus)
		statGroup.GET("/netdev", ShowNetDevStatus)
		statGroup.GET("/transstat", ShowNetTransStatus)
	}
}
