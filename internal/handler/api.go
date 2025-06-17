package handler

import (
	"net/http"

	"github.com/WaitFme/BingeWatchService/internal/dto"
	"github.com/WaitFme/BingeWatchService/internal/model"
	"github.com/WaitFme/BingeWatchService/internal/service"
	"github.com/WaitFme/BingeWatchService/internal/storage"

	"github.com/gin-gonic/gin"
)

// 创建 Watch 记录
func CreateWatch(c *gin.Context) {
	var watch model.WatchEntity
	if err := c.ShouldBindJSON(&watch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := storage.CreateWatch(&watch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, watch)
}

func Test(ctx *gin.Context) {
	num := ctx.Query("num")
	ctx.JSON(http.StatusOK, gin.H{
		"num": num,
	})
}

func Update(ctx *gin.Context) {
	var req dto.RequestData

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, item := range req.Data {
		storage.UpdateOrCreateWatch(service.Data2Model(item))
	}

	// 返回成功响应
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data received and processed",
		"count":   len(req.Data),
	})
}

func Sync(ctx *gin.Context) {
	// 1. 获取所有未删除的记录
	items, err := storage.GetAllWatches()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch data",
			"details": err.Error(),
		})
		return
	}

	rdata := make([]dto.Data, len(items))

	for i, item := range items {
		rdata[i] = service.Model2Data(item)
	}

	response := dto.ResponseData {
		Data: rdata,
	}

	ctx.JSON(http.StatusOK, response)
}
