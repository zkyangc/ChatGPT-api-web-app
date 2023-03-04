package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

type RequestBody struct {
	Message string `json:"message"`
}

type BotResponseString struct {
	UserInput   string `json:"userInput"`
	BotResponse string `json:"botResponse"`
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	app := fiber.New()
	// todo: ban cors and use react proxy
	app.Use(cors.New())
	api_key := goDotEnvVariable("OPENAI_API_KEY")
	openaiclient := openai.NewClient(api_key)

	app.Post("/", func(c *fiber.Ctx) error {
		payload := RequestBody{}
		if err := c.BodyParser(&payload); err != nil {
			return err
		}
		resp, err := openaiclient.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: payload.Message,
					},
				},
			},
		)
		if err != nil {
			log.Fatal(err)
			return c.SendStatus(500)
		}
		response := BotResponseString{
			UserInput:   payload.Message,
			BotResponse: resp.Choices[0].Message.Content,
		}
		return c.JSON(response)
	})

	app.Listen("localhost:4000")
}
