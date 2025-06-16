package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title string
	Code  string
	Price uint
}

type WatchEntity struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	CurrentEpi  int       `gorm:"column:cepi;not null" json:"cepi"`
	AllEpi      int       `gorm:"column:aepi;not null" json:"aepi"`
	CreatedTime int64     `gorm:"column:ctime;not null" json:"ctime"`
	ChangeTime  int64     `gorm:"column:changetime;not null" json:"changetime"`
	IsDelete    bool      `gorm:"column:isdelete;default:false" json:"isdelete"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// 插入内容
	db.Create(&Product{Title: "新款手机", Code: "D42", Price: 1000})
	db.Create(&Product{Title: "新款电脑", Code: "D43", Price: 3500})

	// 读取内容
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// 更新操作：更新单个字段
	db.Model(&product).Update("Price", 2000)

	// 更新操作：更新多个字段
	db.Model(&product).Updates(Product{Price: 2000, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 2000, "Code": "F42"})

	// 删除操作：
	db.Delete(&product, 1)

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
