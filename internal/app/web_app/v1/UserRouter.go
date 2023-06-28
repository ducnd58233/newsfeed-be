package v1

import (
	"github.com/ducnd58233/newsfeed-be/internal/app/web_app/service"
	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup, svc *service.WebService) {
	userRouter := r.Group("users")
	userRouter.GET("", svc.CheckUserNameAndPassword)
	userRouter.POST("register", svc.CreateUser)
}