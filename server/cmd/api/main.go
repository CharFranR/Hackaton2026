package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"

	usecases "github.com/CharFranR/Hackaton2026/aplication/use-cases"
	"github.com/CharFranR/Hackaton2026/infrastructure/adapters/primary/api"
	"github.com/CharFranR/Hackaton2026/infrastructure/adapters/primary/api/handler"
	"github.com/CharFranR/Hackaton2026/infrastructure/adapters/primary/api/middleware"
	"github.com/CharFranR/Hackaton2026/infrastructure/adapters/secondary/auth"
	repo "github.com/CharFranR/Hackaton2026/infrastructure/adapters/secondary/repository"
	timepkg "github.com/CharFranR/Hackaton2026/infrastructure/adapters/secondary/time"
	"github.com/CharFranR/Hackaton2026/infrastructure/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system env")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("DB_SSLMODE"),
	)

	pool, err := database.CreatePool(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.CloseConnection(pool)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	hasher := auth.NewBcryptHasher(0)
	jwtProvider := auth.NewJWTProvider(jwtSecret, 24*time.Hour)
	clock := timepkg.NewClock()

	userRepo := repo.NewUserRepository(pool)
	companyRepo := repo.NewCompanyRepository(pool)
	offeringRepo := repo.NewOfferingRepository(pool)
	reviewRepo := repo.NewReviewRepository(pool)
	categoryRepo := repo.NewCategoryRepository(pool)
	inquiryRepo := repo.NewInquiryRepository(pool)

	userUC := usecases.NewUserUseCase(userRepo, hasher, jwtProvider, clock)
	companyUC := usecases.NewCompanyUseCase(companyRepo, userRepo, categoryRepo, clock)
	offeringUC := usecases.NewOfferingUseCase(offeringRepo, clock)
	reviewUC := usecases.NewReviewUseCase(reviewRepo, clock)
	categoryUC := usecases.NewCategoryUseCase(categoryRepo)
	inquiryUC := usecases.NewInquiryUseCase(inquiryRepo, clock)

	userHandler := handler.NewUserHandler(userUC)
	companyHandler := handler.NewCompanyHandler(companyUC)
	offeringHandler := handler.NewOfferingHandler(offeringUC)
	reviewHandler := handler.NewReviewHandler(reviewUC)
	categoryHandler := handler.NewCategoryHandler(categoryUC)
	inquiryHandler := handler.NewInquiryHandler(inquiryUC)

	authMW := middleware.NewAuthMiddleware(jwtProvider)

	r := api.NewRouter(userHandler, companyHandler, offeringHandler, reviewHandler, categoryHandler, inquiryHandler, authMW)

	srv := &http.Server{
		Addr:         ":" + serverPort,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("server starting on port %s", serverPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	shutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("shutting down server...")
	if err := srv.Shutdown(shutdown); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
}
