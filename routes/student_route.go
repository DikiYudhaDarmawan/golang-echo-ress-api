package routes

import (
	"belajar_rest_api/controller"

	"github.com/labstack/echo/v4"
)

func StudentRoutes(e *echo.Echo){
	e.POST("/student", controller.CreateStudent)
}