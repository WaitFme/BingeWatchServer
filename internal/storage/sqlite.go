package storage

import (
	"errors"

	"github.com/WaitFme/BingeWatchService/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化 SQLite 数据库
func InitDB() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// 自动迁移模型
	if err := db.AutoMigrate(&model.WatchEntity{}); err != nil {
		return err
	}

	DB = db
	return nil
}

/*func temp() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.WatchEntity{})

	// 插入内容
	db.Create(&model.WatchEntity{Title: "新款手机"})
	db.Create(&model.WatchEntity{Title: "新款电脑"})

	// 读取内容
	var product model.WatchEntity
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// 更新操作：更新单个字段
	db.Model(&product).Update("Price", 2000)

	// 更新操作：更新多个字段
	db.Model(&product).Updates(model.WatchEntity{Title: "2000"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 2000, "Code": "F42"})

	// 删除操作：
	db.Delete(&product, 1)
}*/

// CreateWatch 创建新的观看记录
func CreateWatch(watch *model.WatchEntity) error {
	return DB.Create(watch).Error
}

// GetWatchByTitle 根据标题获取观看记录
func GetWatchByTitle(title string) (*model.WatchEntity, error) {
	var item model.WatchEntity
	err := DB.Where("title = ?", title).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 记录不存在不算错误
		}
		return nil, err
	}
	return &item, nil
}

// UpdateOrCreateWatch 更新或创建观看记录
// 如果title已存在，则按照changeTime最近的数据保存，但createTime保存两个之间最久远那条数据
func UpdateOrCreateWatch(watch *model.WatchEntity) error {
	// 开启事务
	return DB.Transaction(func(tx *gorm.DB) error {
		// 查找是否已存在相同标题的记录
		existing, err := GetWatchByTitle(watch.Title)
		if err != nil {
			return err
		}

		if existing != nil {
			// 比较changeTime
			if watch.ChangeTime > existing.ChangeTime {
				// 保留新数据内容但使用最早的createTime
				watch.CreateTime = min(watch.CreateTime, existing.CreateTime)

				// 删除旧记录
				if err := tx.Delete(existing).Error; err != nil {
					return err
				}
			} else {
				// 已有数据的changeTime更新，不需要修改
				return nil
			}
		}

		// 创建新记录
		return tx.Create(watch).Error
	})
}

// DeleteWatch 删除观看记录
func DeleteWatch(id uint) error {
	return DB.Delete(&model.WatchEntity{}, id).Error
}

// GetAllWatches 获取所有观看记录
func GetAllWatches() ([]model.WatchEntity, error) {
	var items []model.WatchEntity
	err := DB.Find(&items).Error
	return items, err
}

// GetWatchByID 根据ID获取观看记录
func GetWatchByID(id uint) (*model.WatchEntity, error) {
	var item model.WatchEntity
	err := DB.First(&item, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

// min 辅助函数，获取最小值
func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
