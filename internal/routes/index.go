package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sashabaranov/go-openai"
	"github.com/zkyangc/ChatGPT-api-web-app/internal/models"
	"log"
)

func IndexHandler(client *openai.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := models.RequestBody{}
		if err := c.BodyParser(&payload); err != nil {
			return err
		}
		resp, err := client.CreateChatCompletion(
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
		response := models.BotResponseString{
			UserInput:   payload.Message,
			BotResponse: resp.Choices[0].Message.Content,
		}
		return c.JSON(response)
	}
}
