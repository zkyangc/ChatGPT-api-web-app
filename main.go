package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
	"github.com/zkyangc/ChatGPT-api-web-app/internal/routes"
	"github.com/zkyangc/ChatGPT-api-web-app/pkg/utils"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	// todo: ban cors and use react proxy
	app.Use(cors.New())
	apiKey := utils.GoDotEnvVariable("OPENAI_API_KEY")
	openaiclient := openai.NewClient(apiKey)

	// Secret key for JWT signing
	//jwtSecretkey := goDotEnvVariable("JWT_SECRET_KEY")

	app.Post("/", routes.IndexHandler(openaiclient))
	app.Listen("localhost:4000")
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
