package controller

import (
	"strconv"

	"github.com/biswajitpain/golang-gin-api/entity"
	"github.com/biswajitpain/golang-gin-api/service"
	"github.com/biswajitpain/golang-gin-api/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type VideoController interface {
	FindAll() []entity.Video
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

// Delete implements VideoController.
func (c *controller) Delete(ctx *gin.Context) error {
	var video entity.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	c.service.Delete(video)
	return nil
}

// Update implements VideoController.
func (c *controller) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Update(video)
	return nil
}

// FindAll implements VideoController.
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// Save implements VideoController.
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return err
}

func New(svc service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: svc,
	}
}
