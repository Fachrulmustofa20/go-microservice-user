package handler

import (
	"github.com/Fachrulmustofa20/go-microservice-user/service"
	"github.com/Fachrulmustofa20/go-microservice-user/service/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUsecase service.UsersUsecase
}

func NewUserHandler(r *gin.Engine,
	userUsecase service.UsersUsecase,
) {
	handler := &Handler{
		userUsecase: userUsecase,
	}
	// test
	r.GET("/api/users/test", handler.Welcome)

	// users
	r.POST("/api/users/register", handler.Register)
	r.POST("/api/users/login", handler.Login)
	r.PUT("/api/users/profile", middleware.Authentication(), handler.UpdateProfile)
	r.GET("/api/users/profile", middleware.Authentication(), handler.GetProfile)
}
