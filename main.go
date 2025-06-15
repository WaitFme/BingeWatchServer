package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()

	ginServer.LoadHTMLGlob("web/templates/*")

	ginServer.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "server msg"})
	})

	ginServer.GET("/test", test)

	apiGounp := ginServer.Group("/api")
	{
		apiGounp.POST("/update", update)
		apiGounp.GET("/sync", sync)
	}

	log.Println("服务器启动，访问 http://localhost:8082")

	ginServer.Run(":8082")
}

func test(ctx *gin.Context) {
	num := ctx.Query("num")
	ctx.JSON(http.StatusOK, gin.H{
		"num": num,
	})
}

func update(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var m map[string]interface{}
	_ = json.Unmarshal(data, &m)
	ctx.JSON(http.StatusOK, m)
}

func sync(ctx *gin.Context) {

}
