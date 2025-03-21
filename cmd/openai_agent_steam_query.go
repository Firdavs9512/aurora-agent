package cmd

import (
	"aurora-agent/config"
	"aurora-agent/utils"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/sashabaranov/go-openai"
)

// StreamQuery sends a prompt to OpenAI and streams the response to the writer
func (a *OpenAIAgent) StreamQuery(prompt string, writer io.Writer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Add user message to history
	a.messages = append(a.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})

	// Create a streaming request
	stream, err := a.client.CreateChatCompletionStream(
		ctx,
		openai.ChatCompletionRequest{
			Model:    a.model,
			Messages: a.messages,
			Stream:   true,
		},
	)
	if err != nil {
		return fmt.Errorf("OpenAI API stream error: %v", err)
	}
	defer stream.Close()

	// Variable to collect the full response
	fullResponse := ""

	// Ansi buffer
	ansiBuffer := ""

	// Stream the response
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("stream error: %v", err)
		}

		// Get the content delta
		content := response.Choices[0].Delta.Content
		if content != "" {
			// Collect the full response
			fullResponse += content

			// add to ansi buffer
			ansiBuffer += content

			// If buffer contains ANSI code
			if config.AnsiPattern.MatchString(ansiBuffer) {
				// Buffer has ANSI code, process it
				processedBuffer := utils.ProcessANSICodes(ansiBuffer)
				fmt.Print(processedBuffer)
				ansiBuffer = ""
			} else if config.AnsiStartPattern.MatchString(ansiBuffer) && len(ansiBuffer) > 100 {
				// If buffer contains the start of an ANSI code, but not the end
				// and buffer length is more than 100, process it
				// This can happen when ANSI code is in incorrect format
				processedBuffer := utils.ProcessANSICodes(ansiBuffer)
				fmt.Print(processedBuffer)
				ansiBuffer = ""
			} else if len(ansiBuffer) > 80 && !config.AnsiStartPattern.MatchString(ansiBuffer) {
				// If buffer length is more than 80 and no ANSI code start is found,
				// process it
				processedBuffer := utils.ProcessANSICodes(ansiBuffer)
				fmt.Print(processedBuffer)
				ansiBuffer = ""
			}
		}
	}

	// Process remaining buffer
	if ansiBuffer != "" {
		processedBuffer := utils.ProcessANSICodes(ansiBuffer)
		fmt.Print(processedBuffer)
	}

	// Add assistant response to history
	a.messages = append(a.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: fullResponse,
	})

	return nil
}
