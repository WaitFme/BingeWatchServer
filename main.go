package main

import (
	"log"

	"github.com/WaitFme/BingeWatchService/internal/handler"
	"github.com/WaitFme/BingeWatchService/internal/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	if err := storage.InitDB(); err != nil {
		log.Fatal("Database init failed:", err)
	}

	// 设置路由
	service := gin.Default()

	service.GET("/", func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"msg": "server no page"})
	})

	service.GET("/test", handler.Test)

	apiGounp := service.Group("/api")
	{
		apiGounp.POST("/upload", handler.Update)
		apiGounp.GET("/sync", handler.Sync)
	}

	// 启动服务器
	if err := service.Run(":8082"); err != nil {
		log.Fatal("Server failed:", err)
	}

	log.Println("服务器启动，访问 http://localhost:8082")
}
