package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routes"
)

func main() {
	// 创建数据库
	// 连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}

	defer dao.Close()

	// 绑定模型
	dao.DB.AutoMigrate(&models.Todo{})

	// 注册路由
	r := routes.SetupRouter()

	r.Run(":8000")
}
