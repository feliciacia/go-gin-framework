package service

import (
	"github.com/feliciacia/go-gin-framework/go-gin-framework/entity"
	"github.com/feliciacia/go-gin-framework/go-gin-framework/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) Update(video entity.Video) {
	service.videoRepository.Update(video)
}

func (service *videoService) Delete(video entity.Video) {
	service.videoRepository.Delete(video)
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}
