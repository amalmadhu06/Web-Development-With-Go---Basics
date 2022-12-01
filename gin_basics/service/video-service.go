package service

import (
	"golang-gin/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	findAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) findAll() []entity.Video {
	return service.videos
}
