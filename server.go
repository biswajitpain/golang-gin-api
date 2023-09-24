package main

import (
	"io"
	"net/http"
	"os"

	"github.com/biswajitpain/golang-gin-api/controller"
	"github.com/biswajitpain/golang-gin-api/middlewares"
	"github.com/biswajitpain/golang-gin-api/repository"
	"github.com/biswajitpain/golang-gin-api/service"
	"github.com/gin-gonic/gin"
)

var (
	repo            repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(repo)
	loginsvc        service.LoginService       = service.NewLoginService()
	jwtSvc          service.JWTService         = service.NewJWTService()
	videoController controller.VideoController = controller.New(videoService)
	loginctlr       controller.LoginController = controller.NewLoginController(loginsvc, jwtSvc)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()
	// server.Use(
	// 	gin.Recovery(),
	// 	middlewares.Logger(),
	// 	middlewares.BasicAuth(),
	// )

	// Login Endpoint: Authentication + Token Creation

	server.POST("/login", func(ctx *gin.Context) {
		//log.Println(ctx)
		token := loginctlr.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is valid"})
		})
		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Video Updated"})
		})
		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Video Deleted"})
		})
	}
	server.Run(":8000")
}
