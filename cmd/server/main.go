package main

import (
	"context"
	"log"

	"avito-tech-fall-2025/internal/app/http"
	"avito-tech-fall-2025/internal/app/services"
	"avito-tech-fall-2025/internal/config"
	"avito-tech-fall-2025/internal/repository/postgres"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := postgres.NewDB(context.Background(), cfg.DBConn)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	userRepo := postgres.NewUserRepository(db)
	teamRepo := postgres.NewTeamRepository(db)
	prRepo := postgres.NewPRRepository(db)

	userService := services.NewUserService(userRepo, teamRepo)
	teamService := services.NewTeamService(teamRepo, userRepo)
	prService := services.NewPRService(prRepo, userRepo, teamRepo)

	router := http.NewRouter(userService, teamService, prService)

	log.Printf("server started on :8080")
	log.Fatal(router.Listen(":8080"))
}
