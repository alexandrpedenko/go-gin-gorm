package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"

	database "gin-gorm-crud/core/database"
	middlewares "gin-gorm-crud/core/middlewares"
	utils "gin-gorm-crud/core/utils"
	posts "gin-gorm-crud/posts"
)

func initMiddlewares(app *gin.Engine) {
	middlewares.GlobalErrorHandler(app)
}

func main() {
	app := gin.Default()
	validate := validator.New()

	// NOTE: Database connection
	db := database.DatabaseConnection()

	// NOTE: Middlewares
	initMiddlewares(app)

	// NOTE: Modules
	apiRouteGroup := app.Group("/api")
	posts.PostsModule(posts.PostsModuleArgs{
		Db:          db,
		Validate:    validate,
		RouterGroup: apiRouteGroup,
	})

	// NOTE: Run server
	log.Info().Msg("Starting server...")
	err := app.Run(":9001")
	utils.ErrorLog("Server running", err)
}
