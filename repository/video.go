package repository

import (
	"github.com/biswajitpain/golang-gin-api/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	FindAll() []entity.Video
	Delete(entity.Video)
	//CloseDB()
}

type database struct {
	connection *gorm.DB
}

// Delete implements VideoRepository.
func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)

}

// CloseDB implements VideoRepository.
// func (db *database) CloseDB() {
// 	err := db.connection.Close()
// 	if err != nil {
// 		panic("Faild to close database")
// 	}
// }

// FindAll implements VideoRepository.
func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}

// Save implements VideoRepository.
func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

// Update implements VideoRepository.
func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}
