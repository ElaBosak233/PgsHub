package router

import (
	"github.com/elabosak233/pgshub/controller"
	"github.com/gin-gonic/gin"
)

func NewGroupRouter(groupRouter *gin.RouterGroup, groupController *controller.GroupController) {
	groupRouter.GET("/", groupController.FindAll)
	groupRouter.GET("/id/:id", groupController.FindById)
	groupRouter.POST("/", groupController.Create)
	groupRouter.PATCH("/:id", groupController.Update)
	groupRouter.DELETE("/:id", groupController.Delete)
	groupRouter.POST("/join", groupController.AddUserToGroup)
}
