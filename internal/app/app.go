package app

import (
	"context"
	"fmt"
	"github.com/RicliZz/aidentity/internal/api"
	"github.com/RicliZz/aidentity/internal/repositories/authRepository"
	"github.com/RicliZz/aidentity/internal/repositories/qualityRepository"
	"github.com/RicliZz/aidentity/internal/server"
	"github.com/RicliZz/aidentity/internal/services/authService"
	"github.com/RicliZz/aidentity/internal/services/qualityService"
	"github.com/RicliZz/avito-internship-pvz-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	logger.InitLogger()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := pgxpool.New(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Проверка соединения
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("Ping failed: %v\n", err)
	}
	fmt.Println("Successfully connected to the database")

	//repo
	qualityRepo := qualityRepository.NewQualityRepository(conn)
	authRepo := authRepository.NewAuthenticationRepository(conn)
	//service
	qualityServ := qualityService.NewQualityService(qualityRepo)
	authServ := authService.NewAuthenticationService(authRepo)
	//handler
	qualityHand := api.NewQualityHandlers(qualityServ)
	authHand := api.NewAuthenticationHandler(authServ)
	//default route
	router := gin.Default()
	API := router.Group("/api/v1")

	//swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = "/api/v1"

	//REGISTER_ROUTES
	qualityHand.InitQualityHandlers(API)
	authHand.InitAuthenticationHandlers(API)

	//Инициализация и конфигурация HTTP сервера
	srv := server.NewAPIServer(router)

	//Старт сервера
	go srv.Start()

	//Выключение
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		logger.Logger.Fatalw("Shutdown error",
			"error", err)
	}
}
