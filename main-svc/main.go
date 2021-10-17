package main

import (
	"main-svc/controllers"
	"main-svc/interfaces"
	"main-svc/routes"
	"main-svc/usecases"
	"os"

	middlewareApp "main-svc/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Data struct {
	Message string
}

func main() {
	_ = godotenv.Load()

	configDB := ConfigDB{
		DB_Host:     os.Getenv("DB_HOST"),
		DB_User:     os.Getenv("DB_USER"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Name:     os.Getenv("DB_NAME"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_SSLMode:  os.Getenv("DB_SSLMODE"),
		DB_TimeZone: os.Getenv("DB_TIMEZONE"),
	}

	db := configDB.InitDB()

	// PUT ALL ENV HERE
	SERVICE_PORT := os.Getenv("SERVICE_PORT")

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, Data{
			Message: "pong from main svc",
		})
	})

	// JWT Middleware
	configJWT := middlewareApp.ConfigJWT{
		SecretJWT:       os.Getenv("JWT_SECRET"),
		ExpiresDuration: int64(60 * 60 * 24 * 7),
	}

	// Injector
	userRepo := interfaces.NewUserRepo(db)
	userInteractor := usecases.NewUserInteractor(userRepo)
	userController := controllers.NewUserController(userInteractor)

	quoteRepo := interfaces.NewQuoteRepo(db)
	quoteSaver := interfaces.ProvideQuoteSaver(quoteRepo)
	quoteGetter := interfaces.NewQuoteGetter(quoteSaver)
	quoteInteractor := usecases.NewQuoteInteractor(*userInteractor, quoteRepo, quoteGetter)
	quoteController := controllers.NewQuoteController(quoteInteractor)
	// Controller
	controller := routes.ControllerSetup{
		UserController:  *userController,
		QuoteController: *quoteController,
		JWTMiddleware:   configJWT.Init(),
	}

	controller.InitRoute(e)
	e.Logger.Fatal(e.Start(":" + SERVICE_PORT))
}
