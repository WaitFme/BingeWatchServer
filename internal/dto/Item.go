package dto

type Data struct {
	Title          string `json:"title"`
	Remarks        string `json:"remarks"`
	State          int    `json:"state"`
	CurrentEpisode int    `json:"currentEpisode"`
	TotalEpisode   int    `json:"totalEpisode"`
	CreateTime     int64  `json:"createTime"`
	ChangeTime     int64  `json:"changeTime"`
	IsDelete       bool   `json:"isDelete"`
}
