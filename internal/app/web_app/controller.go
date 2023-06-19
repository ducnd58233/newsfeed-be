package web_app

import (
	"fmt"

	service "github.com/ducnd58233/newsfeed-be/internal/app/web_app/service"
	v1 "github.com/ducnd58233/newsfeed-be/internal/app/web_app/v1"
	"github.com/gin-gonic/gin"
)

type WebController struct {
	WebService service.WebService
	Port       int
}

func (c WebController) Run() {
	r := gin.Default()

	v1Router := r.Group("v1")
	v1.AddNewsfeedRouter(v1Router, &c.WebService)
	v1.AddUserRouter(v1Router, &c.WebService)

	r.Run(fmt.Sprintf(":%d", c.Port))
}
