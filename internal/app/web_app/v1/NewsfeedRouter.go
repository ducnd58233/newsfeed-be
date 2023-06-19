package v1

import (
	"github.com/ducnd58233/newsfeed-be/internal/app/web_app/service"
	"github.com/gin-gonic/gin"
)

func AddNewsfeedRouter(r *gin.RouterGroup, svc *service.WebService) {
	newsfeedRouter := r.Group("newsfeeds")
	newsfeedRouter.GET("", func(context *gin.Context) {
		
	})
}