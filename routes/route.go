package routes

import (
	user "user/controller"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uc user.UserController) {
	e.GET("/users", uc.GetAllUser)
	e.GET("/users/:id", uc.GetUserbyid)
	e.POST("/users", uc.InsertNewUser)
	e.PUT("/users/:id", uc.UpdateUser)
	e.DELETE("/users/:id", uc.DeleteUser)
}
