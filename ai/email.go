package ai

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func GenerateEmail(context string) string {
	err := godotenv.Load("FUNGUO")
	if err != nil {
		log.Fatalf("Failed to load .env; %v", err)
	}

	AI_KEY := os.Getenv("FUNGUO_API_KEY") // Get API key
	client := openai.NewClient(AI_KEY)    // Create new OpenAI client

	// Generate a chat completion response based on the provided context and system message
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "You're a sales representive for a SAAS business. Generate an engaging email, informing the client of a lucrative festive season 50% OFF offer."},
			{Role: "user", Content: fmt.Sprintf("Generate a personalized email template based on the context: %s", context)},
		},
	})
	if err != nil {
		log.Fatalf("Error generating chat completion: %v", err)
	}
	return resp.Choices[0].Message.Content
}
