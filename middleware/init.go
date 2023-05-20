package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(c *gin.Context) {
	//打印当前时间
	fmt.Println(time.Now())
	//获取URL
	fmt.Println(c.Request.URL)
	//
	c.Set("name", "梁鑫")

	//定义一个goroutine统计日志
	cCp := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Done! in path" + cCp.Request.URL.Path)
	}()
}
