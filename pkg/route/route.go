package route

import (
	"LotusPart2/pkg/handler"
	"LotusPart2/pkg/infra"
	"LotusPart2/pkg/middleware"
	"LotusPart2/pkg/model"
	"LotusPart2/pkg/repo"
	internalService "LotusPart2/pkg/service"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/service"
)

type Service struct {
	*service.BaseApp
}

func StartNewService() *Service {

	s := &Service{
		BaseApp: service.NewApp("Service user", "v1.0"),
	}

	// Set max memory limit to 8Mib for multipart forms

	dbConn := infra.PostgresConn()
	if err := dbConn.Debug().AutoMigrate(&model.User{}, &model.Auth{}); err != nil {
		panic(err)
	}

	repository := repo.NewPGRepo(dbConn)

	// service
	userService := internalService.NewUserService(repository)

	// handle
	validate := validator.New()
	handlers := handler.NewUserHandler(userService, validate)

	v1Api := s.Router.Group("/api/v1")
	swaggerApi := s.Router.Group("/")

	// swagger
	swaggerApi.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	v1Api.POST("/user/register", ginext.WrapHandler(handlers.Register))
	v1Api.POST("/user/login", ginext.WrapHandler(handlers.Login))
	v1Api.POST("/user/logout", ginext.WrapHandler(handlers.Logout)) // logout => revoke token

	// file
	v1Api.POST("/file/upload", middleware.VerifyToken(), ginext.WrapHandler(handlers.UploadFile))

	// Migrate
	return s
}
