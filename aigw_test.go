package openai_test

import (
	"testing"

	"github.com/openai/openai-go"
	"github.com/stretchr/testify/assert"
)

func TestForAIGW_AccumulateDelta(t *testing.T) {
	cc := openai.ChatCompletion{
		Choices: []openai.ChatCompletionChoice{
			{},
		},
	}
	chunk := openai.ChatCompletionChunk{
		ID: "chatcmpl",
		Choices: []openai.ChatCompletionChunkChoice{
			{
				Index: 0,
				Delta: openai.ChatCompletionChunkChoiceDelta{},
			},
			{
				Index: 0,
				Delta: openai.ChatCompletionChunkChoiceDelta{
					Content: strptr("hello"),
				},
			},
		},
	}
	accumulator := openai.ChatCompletionAccumulator{
		ChatCompletion: cc,
	}
	accumulator.AddChunk(chunk)
	assert.Equal(t, "hello", cc.Choices[0].Message.Content)
}
