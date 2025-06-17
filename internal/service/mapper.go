package service

import (
	"github.com/WaitFme/BingeWatchService/internal/dto"
	"github.com/WaitFme/BingeWatchService/internal/model"
)

func Data2Model(data dto.Data) *model.WatchEntity {
	return &model.WatchEntity {
		Title:          data.Title,
		Remarks:        data.Remarks,
		CurrentEpisode: data.CurrentEpisode,
		TotalEpisode:   data.TotalEpisode,
		CreateTime:     data.CreateTime,
		ChangeTime:     data.ChangeTime,
		IsDelete:       data.IsDelete,
		State:          data.State,
	}
}

func Model2Data(data model.WatchEntity) dto.Data {
	return dto.Data {
		Title:          data.Title,
		Remarks:        data.Remarks,
		CurrentEpisode: data.CurrentEpisode,
		TotalEpisode:   data.TotalEpisode,
		CreateTime:     data.CreateTime,
		ChangeTime:     data.ChangeTime,
		IsDelete:       data.IsDelete,
		State:          data.State,
	}
}
