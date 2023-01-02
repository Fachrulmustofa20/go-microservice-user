package handler

import (
	"net/http"

	"github.com/Fachrulmustofa20/go-microservice-user/models"
	"github.com/Fachrulmustofa20/go-microservice-user/service/utils"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func (handler Handler) Register(ctx *gin.Context) {
	var users models.Users
	var profile models.Profile

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	valid, err := govalidator.ValidateStruct(users)
	if err != nil || !valid {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "A validation error occurred. Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	userId, err := handler.userUsecase.Register(users)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Status Bad Request",
			"error":   err.Error(),
		})
		return
	}

	profile.UserId = userId
	err = handler.userUsecase.CreateProfile(profile)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Status Bad Request",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success Created Account!",
	})
}

func (handler Handler) Login(ctx *gin.Context) {
	var users models.Users
	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	if users.Email == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please input your email!",
		})
		return
	}
	if users.Password == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please input your password!",
		})
		return
	}

	password := users.Password
	token, err := handler.userUsecase.Login(users.Email, password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}

func (handler Handler) UpdateProfile(ctx *gin.Context) {
	var profile models.Profile
	userId := utils.GetUserIdJWT(ctx)
	profile.UserId = userId

	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "please check the form and try again!",
			"error":   err.Error(),
		})
		return
	}

	if err := handler.userUsecase.UpdateProfileByUserId(profile); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Update Profile",
		"data": map[string]interface{}{
			"age":         profile.Age,
			"hoby":        profile.Hoby,
			"description": profile.Description,
			"photo":       profile.Photo,
		},
	})

}

func (handler Handler) GetProfile(ctx *gin.Context) {
	userId := utils.GetUserIdJWT(ctx)

	user, err := handler.userUsecase.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	profile, err := handler.userUsecase.GetProfileByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"email":      user.Email,
		"fullname":   user.FullName,
		"age":        profile.Age,
		"hoby":       profile.Hoby,
		"photo":      profile.Photo,
		"created_at": user.CreatedAt,
	})
}
