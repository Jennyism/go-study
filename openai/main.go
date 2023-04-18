package main

import (
	"bufio"
	"fmt"
	"openai/api"
	"os"
	"strings"
)

func main() {
	apiKey := "_"
	temperature := float32(0.5)
	maxTokens := 50
	stopSequences := []string{"\n"}

	previousResponse := ""
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("You: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}
		response, err := api.GetOpenAIResponse(input, apiKey, temperature, maxTokens, stopSequences, previousResponse)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("OpenAI: " + response)
		previousResponse = response
	}

}
