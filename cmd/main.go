package main

import (
	"test/database"
	"test/internal/handler"
	"test/internal/middleware"
	"test/internal/repositories"
	"test/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	db := database.StartPostgres()
	repo := repositories.NewTestRepo(db)
	service := services.NewTestService(repo)

	middleware := middleware.NewTweetMiddleware(service)
	handler.NewTestHandler(r,middleware, service)
	r.Run(":2400")
}
