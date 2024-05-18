package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

// use godot package to load/read the .env file and
// return the value of the key
func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

// goDotEnvVariable returns the value of the given key from the environment variables
func goDotEnvVariable(key string) string {
	// Load .env file
	err := loadEnv()
	if err != nil {
		fmt.Printf("Completion error: ", err)
	}

	// Get the value of the key
	value := os.Getenv(key)
	if value == "" {
		fmt.Errorf("key %s not found in the environment variables", key)
	}

	return value
}

func resultGen(prompt_ string) {
	// godotenv package
	secretKey := goDotEnvVariable("OPENAI_SECRET_KEY")

	// create a new client
	client := openai.NewClient(secretKey)
	ctx := context.Background()
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt_,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func main() {
	resultGen("How are you?")
}
