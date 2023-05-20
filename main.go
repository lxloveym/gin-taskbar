package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test/models"
	"test/routes"
)

func main() {
	r := gin.Default()
	//公共部分
	models.WebFile(r)
	//路由抽离
	routes.ReadyDo(r)
	//访问端口设置
	r.Run(":80")
	fmt.Println("liangxin")
}
