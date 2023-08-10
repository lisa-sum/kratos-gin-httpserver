package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"tiktok/internal/service"
)

type UserUseCase struct {
	userService *service.UserService
	log         *log.Helper
}

func NewUserUseCase(us *service.UserService, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		userService: us,
		log:         log.NewHelper(logger),
	}
}

func (us *UserUseCase) Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
