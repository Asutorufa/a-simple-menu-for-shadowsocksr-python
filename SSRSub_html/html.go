package main

import (
	"os"

	getdelay "../net"
	ssr_process "../process"
	"../subscription"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var configPath = os.Getenv("HOME") + "/.config/SSRSub"
var sqlPath = configPath + "/SSR_config.db"

func main() {
	router := gin.Default()

	// file, _ := exec.LookPath(os.Args[0])
	// path2, _ := filepath.Abs(file)
	// rst := filepath.Dir(path2)

	router.LoadHTMLGlob("./**/**")
	router.GET("/", func(c *gin.Context) {
		_, ssrStatus := ssr_process.Get(configPath)
		c.HTML(200, "sidebar.html", gin.H{
			"now_node":       subscription.GetNowNode(sqlPath),
			"ssr_status":     ssrStatus,
			"title":          "SSRSub",
			"server_remarks": subscription.GetAllNodeRemarksAndID(sqlPath),
			"home":           true,
		})
	})

	router.GET("/test", func(c *gin.Context) {
		c.String(200, "tests")
	})

	router.POST("/submit", func(c *gin.Context) {
		id, _ := c.GetPostForm("server")
		subscription.SsrSQLChangeNode(id, sqlPath)
		_, ssrStatus := ssr_process.Get(configPath)
		if ssrStatus == true {
			ssr_process.Stop(configPath)
			ssr_process.Start(configPath, sqlPath)
		}
		node := subscription.GetOneNodeAll(id, sqlPath)
		delay, _ := getdelay.Tcp_delay(node["server"], node["server_port"])
		c.HTML(200, "sidebar.html", gin.H{
			"id":          node["id"],
			"remarks":     node["remarks"],
			"server":      node["server"],
			"server_port": node["server_port"],
			"protocol":    node["protocol"],
			"method":      node["method"],
			"obfs":        node["obfs"],
			"password":    node["password"],
			"obfsparam":   node["obfsparam"],
			"protoparam":  node["protoparam"],
			"delay":       delay,
			"server_bool": true,
		})
		// c.String(200, getOneNodeAll(id)["remarks"])
	})

	router.GET("/link", func(c *gin.Context) {
		link := subscription.Get_subscription_link(sqlPath)
		c.HTML(200, "sidebar.html", gin.H{
			"subscription": true,
			"link":         link,
		})
	})

	router.Run(":8081")
}
