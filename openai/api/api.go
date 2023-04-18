package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func GetOpenAIResponse(prompt string, apiKey string, temperature float32, maxTokens int, stopSequences []string, previousResponse string) (string, error) {
	endpoint := "https://api.openai.com/v1/engines/davinci-codex/completions"

	fmt.Printf("%s\n", previousResponse)
	payload := struct {
		Prompt           string   `json:"prompt"`
		Temperature      float32  `json:"temperature"`
		MaxTokens        int      `json:"max_tokens"`
		Stop             []string `json:"stop,omitempty"`
		PreviousResponse string   `json:"previous_response,omitempty"`
	}{
		Prompt:           prompt,
		Temperature:      temperature,
		MaxTokens:        maxTokens,
		Stop:             stopSequences,
		PreviousResponse: previousResponse,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var openAIResp OpenAIResponse
	err = json.Unmarshal(body, &openAIResp)
	if err != nil {
		return "", err
	}

	fmt.Printf("%+v\n", string(body))
	if len(openAIResp.Choices) > 0 {
		return openAIResp.Choices[0].Text, nil
	}

	return "", nil
}
