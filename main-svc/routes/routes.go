package routes

import (
	"main-svc/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Let the caller define the interface it needs

type ControllerSetup struct {
	UserController  controllers.UserController
	QuoteController controllers.QuoteController
	JWTMiddleware   middleware.JWTConfig
}

func (controller ControllerSetup) InitRoute(e *echo.Echo) {
	users := e.Group("/users")
	users.POST("/", controller.UserController.Register)
	users.GET("/:userId", controller.UserController.GetUser)

	quote := e.Group("/quote")
	quote.GET("/", controller.QuoteController.Quote)
	quote.POST("/:quiteId/save", controller.QuoteController.SaveFavoriteQuote)
	quote.GET("/favorites", controller.QuoteController.ListAllFavoriteQuotes)
}
