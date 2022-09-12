package routes

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 告诉gin框架静态资源放在哪里的
	r.Static("/static", "static")

	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 增加
		v1Group.POST("/todo", controller.CreatTodo)
		// 查看
		// 查询所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateAllTodo)
		// 删除某个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
