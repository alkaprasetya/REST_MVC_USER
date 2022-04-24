package routes

import (
	user "user/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc user.UserController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/users", uc.GetAllUser)
	e.GET("/users/:id", uc.GetUserbyid)
	e.POST("/users", uc.InsertNewUser)
	e.PUT("/users/:id", uc.UpdateUser)
	e.DELETE("/users/:id", uc.DeleteUser)

	kelompokGET := e.Group("/users")
	kelompokGET.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time:${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))

}
