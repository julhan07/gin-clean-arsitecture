package controller

import (
	"clean-code/constant"
	"clean-code/service"

	router "clean-code/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service service.UserService
	route   router.MappingRouter
}

func NewUserController(service service.UserService, r router.MappingRouter) userController {
	controller := &userController{
		service: service,
		route:   r,
	}

	controller.routing()
	return *controller
}

func (v *userController) routing() {
	v.route.GET(constant.USER, v.GetAll)
}

func (v *userController) GetAll(c *gin.Context) {

	users, err := v.service.GetAll(c)

	if err != nil {
		c.JSON(200, err.Error())
	}

	c.JSON(200, users)
}
