package ai

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func GenerateEmail(emailContext string) (string, error) {
	// Load the.env file to get the API key
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env; %v", err)
	}

	// Get the API key from environment variable and check if it's set
	AI_KEY := os.Getenv("FUNGUO")
	if AI_KEY == "" {
		log.Fatal("FUNGUO environment variable not set")
	}

	// Create the Gemini client
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(AI_KEY))
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	// Call the Gemini API
	resp, err := model.GenerateContent(context.Background(), genai.Text(emailContext))
	if err != nil {
		return "", err
	}

	// Check if the response contains any content, return approproately
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		if textContent, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
			return string(textContent), nil
		}
	}
	// Return error
	err = errors.New("failed to generate email content")
	return "", err
}
