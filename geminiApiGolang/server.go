package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Enter your apiKey on there
const apiKey = "YOUR_API_KEY"

var gClient *genai.Client

type Data struct {
	Request string
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	gClient = client

	if err != nil {
		log.Fatal(err)
	}

	app.Post("/", postHandle)

	app.Listen(":3000")
}

func postHandle(c fiber.Ctx) error {
	// Initialize the chat
	user := new(Data)

	err := c.Bind().Body(user)
	if err != nil {
		return err
	}

	//Enter your Gemini Model on there
	model := gClient.GenerativeModel("YOUR_GEMINI_MODEL")

	cs := model.StartChat()
	resp, err := cs.SendMessage(c.Context(), genai.Text(user.Request))

	if err != nil {
		return err
	}
	return c.SendString(fmt.Sprint(resp.Candidates[0].Content.Parts[0]))
}
