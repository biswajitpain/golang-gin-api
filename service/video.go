package service

import (
	"github.com/biswajitpain/golang-gin-api/entity"
	"github.com/biswajitpain/golang-gin-api/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

type videoService struct {
	videoRepoitory repository.VideoRepository
}

// Delete implements VideoService.
func (service *videoService) Delete(video entity.Video) {
	service.videoRepoitory.Delete(video)
}

// Update implements VideoService.
func (service *videoService) Update(video entity.Video) {
	service.videoRepoitory.Update(video)
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepoitory: repo,
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videoRepoitory.Save(video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepoitory.FindAll()
}
