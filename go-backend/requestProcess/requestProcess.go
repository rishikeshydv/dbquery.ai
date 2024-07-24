package requestProcess

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

// RequestBody defines the structure of the expected JSON body
type RequestBody struct {
	Database string `json:"database"`
	Prompt   string `json:"prompt"`
}

// use godot package to load/read the .env file and
// return the value of the key
func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

// goDotEnvVariable returns the value of the given key from the environment variables
func GoDotEnvVariable(key string) string {
	// Load .env file
	err := LoadEnv()
	if err != nil {
		fmt.Printf("Completion error: %v", err)
	}

	// Get the value of the key
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("key %s not found in the environment variables", key)
	}

	return value
}

func ResultGen(db string, prompt_ string) (string, error) {
	// godotenv package
	secretKey := GoDotEnvVariable("OPENAI_SECRET_KEY")

	// create a new client
	client := openai.NewClient(secretKey)
	ctx := context.Background()

	//fine-tuning the model
	// file, err := client.CreateFile(ctx, openai.FileRequest{
	// 	FilePath: "training_datasets.json",
	// 	Purpose:  "fine-tune",
	// })
	// if err != nil {
	// 	fmt.Printf("Upload JSONL file error: %v\n", err)
	// 	return "", err
	// }

	// fineTuningJob, err := client.CreateFineTuningJob(ctx, openai.FineTuningJobRequest{
	// 	TrainingFile: file.ID,
	// 	Model:        "gpt-3.5-turbo-0613", // gpt-3.5-turbo-0613, babbage-002.
	// })
	// if err != nil {
	// 	fmt.Printf("Creating new fine tune model error: %v\n", err)
	// 	return "", err
	// }

	// fineTuningJob, err = client.RetrieveFineTuningJob(ctx, fineTuningJob.ID)
	// if err != nil {
	// 	fmt.Printf("Getting fine tune model error: %v\n", err)
	// 	return "", err
	// }
	// fmt.Println(fineTuningJob.FineTunedModel)

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Your are a helpful assistant who is helping users to get database queries based on their prompts.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Use the database: " + db + " to get query for the following prompt: " + prompt_,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content, nil
}
