package main

import (
	"blog-api/config"
	"blog-api/core/handlers"
	"blog-api/core/repositories"
	"blog-api/core/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func init() {
	config.NewAppInitEnvironment()
}

func main() {

	db := config.InitDatabase()

	articleRepository := repositories.NewArticleRepositoryDB(db)
	articleService := services.NewArticleService(articleRepository)
	articleHandler := handlers.NewArticleHandler(articleService)

	app := fiber.New()

	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
	}))
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	apiV1 := app.Group("/api/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	apiV1.Get("/articles", articleHandler.GetArticles)
	apiV1.Get("/article/:uid", articleHandler.GetArticleByUID)
	apiV1.Delete("/article/:uid", articleHandler.DeleteArticleByUID)
	apiV1.Post("/article", articleHandler.CreateArticle)

	app.Listen(fmt.Sprintf(":%v", config.Env.Port))
}
