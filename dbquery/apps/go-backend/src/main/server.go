package resultGen

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func resultGen(prompt_ string) {
	// godotenv package
	secretKey := goDotEnvVariable("OPENAI-SECRET-KEY")

	// create a new client
	c := openai.NewClient(secretKey)
	ctx := context.Background()

	// create a completion request
	req := openai.CompletionRequest{
		Model:  openai.GPT3Ada,
		Prompt: prompt_,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
