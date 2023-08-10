package interfaces

import "github.com/gin-gonic/gin"

func RegisterHTTPServer(us *UserUseCase) *gin.Engine {
	r := gin.Default()

	r.Any("/v1/user", us.Register)

	return r
}
