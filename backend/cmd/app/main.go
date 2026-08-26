package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/ohm-pkh/local-line-messageAPI/internal/database"
	"github.com/ohm-pkh/local-line-messageAPI/internal/handler"
	"github.com/ohm-pkh/local-line-messageAPI/internal/repo"
	"github.com/ohm-pkh/local-line-messageAPI/internal/routes"
	"github.com/ohm-pkh/local-line-messageAPI/internal/service"
	ws "github.com/ohm-pkh/local-line-messageAPI/internal/websocket"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using environment variables")
	}
	db := database.New()

	defer db.Client().Disconnect(context.Background())

	channelSecret := uuid.New()
	accessToken := uuid.New()

	connection := ws.NewConnection()
	repo := repo.New(db)
	service := service.New(accessToken, channelSecret, connection, repo)
	handler := handler.New(service, connection)
	app := fiber.New()

	routes.Router(app, handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Listen(":" + port)
}
