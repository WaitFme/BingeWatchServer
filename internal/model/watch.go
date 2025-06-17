package model

type WatchEntity struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title         string    `gorm:"size:255;not null" json:"title"`
	CurrentEpisode int      `gorm:"not null" json:"currentEpisode"`
	TotalEpisode  int       `gorm:"not null" json:"totalEpisode"`
	CreateTime    int64     `gorm:"not null" json:"createTime"`
	ChangeTime    int64     `gorm:"not null" json:"changeTime"`
	IsDelete      bool      `gorm:"default:false" json:"isDelete"`
	Remarks       string    `gorm:"size:500" json:"remarks"`
	State         int       `gorm:"default:0" json:"state"`
}

func (WatchEntity) TableName() string {
	return "watch_entities" // 自定义表名
}
