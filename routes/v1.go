package routes

import (
	"github.com/gin-gonic/gin"
	v1 "test/controllers/v1"
)

func ReadyDo(c *gin.Engine) {
	//首页
	c.GET("/", v1.Index)
	v1Api := c.Group("/v1") //路由中间件
	{
		//添加
		v1Api.POST("/todo", v1.Add)
		//查看
		v1Api.GET("/todo", v1.Find)
		//删除
		v1Api.DELETE("/todo/:id", v1.Del)
		//修改更新
		v1Api.PUT("/todo/:id", v1.Update)
	}
}
