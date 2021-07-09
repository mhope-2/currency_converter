package handler

import (

	"github.com/gin-gonic/gin"
	"github.com/mhope-2/currency_converter/repository"
	"gorm.io/gorm"

	)

type Handler struct {
	DB                   *gorm.DB
	Repo                 *repository.Repository
}


func New(DB *gorm.DB) *Handler {
	repo := repository.New(DB)

	return &Handler{
		DB:                   DB,
		Repo:                 repo,
	}
}

func (h *Handler) Register(v1 *gin.RouterGroup){

	users := v1.Group("/users")
	users.GET("/", h.ListUsers)
	users.GET("/:id", h.ViewUser)
	users.POST("/signup", h.SignupUser)
	users.POST("/login", h.LoginUser)
}





