package routes

import (
	"ujk-golang/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	
	e.GET("/comics", controllers.GetComicsController)
	e.POST("/comics", controllers.AddComicController)
	e.GET("/comics/:id", controllers.GetComicDetailController)
	e.PUT("/comics/:id", controllers.UpdateComicController)
	e.DELETE("/comics/:id", controllers.DeleteComicController)
}