package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/Maksatus123/go-final-project/internal/config"
	"github.com/Maksatus123/go-final-project/internal/controller"
	"github.com/Maksatus123/go-final-project/internal/middleware"
	"github.com/Maksatus123/go-final-project/internal/repository"
	"github.com/Maksatus123/go-final-project/internal/service"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg := config.LoadConfig()

	// Database connection
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	// Repositories
	bookRepo := repository.NewBookRepository(db)
	exchangeRequestRepo := repository.NewExchangeRequestRepository(db)

	// Services
	bookSvc := service.NewBookService(bookRepo)
	exchangeRequestSvc := service.NewExchangeRequestService(bookRepo, exchangeRequestRepo)

	// Controllers
	bookCtrl := controller.NewBookController(bookSvc)
	exchangeRequestCtrl := controller.NewExchangeRequestController(exchangeRequestSvc)

	// Gin setup
	r := gin.Default()
	r.Use(middleware.AuthMiddleware(cfg.JWTSecret))

	// Book endpoints
	r.POST("/books", bookCtrl.CreateBook)
	r.GET("/books/:id", bookCtrl.GetBook)
	r.GET("/books", bookCtrl.GetAllBooks)
	r.PUT("/books/:id", bookCtrl.UpdateBook)
	r.DELETE("/books/:id", bookCtrl.DeleteBook)
	r.GET("/my-books", bookCtrl.GetBooksByOwner) // Filter by owner_id query param

	// Exchange request endpoints
	r.POST("/exchange-requests", exchangeRequestCtrl.CreateExchangeRequest)
	r.GET("/exchange-requests/:id", exchangeRequestCtrl.GetExchangeRequest)
	r.GET("/exchange-requests", exchangeRequestCtrl.GetExchangeRequestsByRequester)
	r.PUT("/exchange-requests/:id", exchangeRequestCtrl.UpdateExchangeRequestStatus)
	// Run server
	logger.Info("Starting Book Service", zap.String("port", cfg.HTTPPort))
	if err := r.Run(":" + cfg.HTTPPort); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}