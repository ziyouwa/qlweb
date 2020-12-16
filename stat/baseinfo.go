package stat

import (
	"io/ioutil"
	"log"
	"os"
	"qlweb/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// Node /proc/uptime
//  第一列:系统启动到现在的时间（以秒为单位），这里简记为num1
//  第二列:系统空闲的时间（以秒为单位），这里简记为num2
//  系统的空闲率(%) = num2/(num1*N) 其中N是SMP系统中的CPU个数
type Node struct {
	Name   string
	Boot   string
	Uptime string
	Time   string
}

func baseInfo() (*Node, error) {
	b, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			b = b[:i]
			break
		}
	}
	t := string(b) + "s"
	d, err := time.ParseDuration(t)
	if err != nil {
		return nil, err
	}
	boot := time.Now().Add(-d).Local().Format("2006-01-02 15:01:01 Mon")
	// boot := time.Now().Add(-d).Format("2006-01-02 15:01:01 Mon")
	tt := time.Now().Local().Format("2006-01-02 15:01:01 Mon")

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	return &Node{hostname, boot, d.String(), tt}, nil
}

//ShowBaseInfo get node system info
func ShowBaseInfo(c *gin.Context) {
	sysinfo, err := baseInfo()
	if err != nil {
		log.Fatal("Get node system info error: ", err)
	}
	utils.Render(
		c,
		gin.H{
			"title":   "System Status",
			"name":    "System Base Infomation",
			"payload": sysinfo,
		},
		"stat.html",
	)
}
