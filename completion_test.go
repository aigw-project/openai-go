// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openai_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/internal/testutil"
	"github.com/openai/openai-go/option"
)

func TestCompletionNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Completions.New(context.TODO(), openai.CompletionNewParams{
		Model: openai.CompletionNewParamsModelGPT3_5TurboInstruct,
		Prompt: openai.CompletionNewParamsPromptUnion{
			OfString: openai.String("This is a test."),
		},
		BestOf:           openai.Int(0),
		Echo:             openai.Bool(true),
		FrequencyPenalty: openai.Float(-2),
		LogitBias: map[string]int64{
			"foo": 0,
		},
		Logprobs:        openai.Int(0),
		MaxTokens:       openai.Int(16),
		N:               openai.Int(1),
		PresencePenalty: openai.Float(-2),
		Seed:            openai.Int(0),
		Stop: openai.CompletionNewParamsStopUnion{
			OfString: openai.String("\n"),
		},
		StreamOptions: openai.ChatCompletionStreamOptionsParam{
			IncludeUsage: openai.Bool(true),
		},
		Suffix:      openai.String("test."),
		Temperature: openai.Float(1),
		TopP:        openai.Float(1),
		User:        openai.String("user-1234"),
	})
	if err != nil {
		var apierr *openai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUnmarshal_CompletionNewParams(t *testing.T) {
	tests := []struct {
		input string
		check func(t *testing.T, msg openai.CompletionNewParams)
	}{
		{
			input: `{"model": "gpt-3.5-turbo-instruct"}`,
			check: func(t *testing.T, params openai.CompletionNewParams) {
				assert.Equal(t, "", params.Prompt.OfString.Value)
				assert.Equal(t, []string(nil), params.Prompt.OfArrayOfStrings)
				assert.Equal(t, []int64(nil), params.Prompt.OfArrayOfTokens)
				assert.Equal(t, [][]int64(nil), params.Prompt.OfArrayOfTokenArrays)
				assert.Equal(t, "", params.Stop.OfString.Value)
				assert.Equal(t, []string(nil), params.Stop.OfCompletionNewsStopArray)

				assert.Equal(t, openai.CompletionNewParamsModelGPT3_5TurboInstruct, params.Model)
			},
		},
		{
			input: `{"prompt": "chatgpt", "model": "gpt-3.5-turbo-instruct", "stop": "test-case1"}`,
			check: func(t *testing.T, params openai.CompletionNewParams) {
				assert.Equal(t, "chatgpt", params.Prompt.OfString.Value)
				assert.Equal(t, []string(nil), params.Prompt.OfArrayOfStrings)
				assert.Equal(t, []int64(nil), params.Prompt.OfArrayOfTokens)
				assert.Equal(t, [][]int64(nil), params.Prompt.OfArrayOfTokenArrays)
				assert.Equal(t, "test-case1", params.Stop.OfString.Value)
				assert.Equal(t, []string(nil), params.Stop.OfCompletionNewsStopArray)
			},
		},
		{
			input: `{"prompt": ["chatgpt", "deepseek_v1"], "model": "gpt-3.5-turbo-instruct", "stop": ["test-case1", "stop"]}`,
			check: func(t *testing.T, params openai.CompletionNewParams) {
				assert.Equal(t, "", params.Prompt.OfString.Value)
				assert.Equal(t, []string{"chatgpt", "deepseek_v1"}, params.Prompt.OfArrayOfStrings)
				assert.Equal(t, []int64(nil), params.Prompt.OfArrayOfTokens)
				assert.Equal(t, [][]int64(nil), params.Prompt.OfArrayOfTokenArrays)
				assert.Equal(t, "", params.Stop.OfString.Value)
				assert.Equal(t, []string{"test-case1", "stop"}, params.Stop.OfCompletionNewsStopArray)
				assert.Equal(t, false, params.StreamOptions.IncludeUsage.Value)
			},
		},

		{
			input: `{"stream": true, "model": "gpt-3.5-turbo-instruct", "temperature": 9.9, "frequency_penalty": -2.0, "presence_penalty": -2.0, "max_tokens": 99, "top_k": 9, "top_p":10 }`,
			check: func(t *testing.T, params openai.CompletionNewParams) {
				assert.Equal(t, true, params.Stream)
				assert.Equal(t, 9.9, params.Temperature.Value)
				assert.Equal(t, -2.0, params.FrequencyPenalty.Value)
				assert.Equal(t, -2.0, params.PresencePenalty.Value)
				assert.Equal(t, int64(99), params.MaxTokens.Value)
				assert.Equal(t, float64(10), params.TopP.Value)
				assert.Equal(t, int64(9), params.TopK.Value)
			},
		},
		{
			input: `{"model": "gpt-3.5-turbo-instruct"}`,
			check: func(t *testing.T, params openai.CompletionNewParams) {
				assert.Equal(t, false, params.Stream)
				assert.Equal(t, 0.0, params.Temperature.Value)
				assert.Equal(t, 0.0, params.FrequencyPenalty.Value)
				assert.Equal(t, 0.0, params.PresencePenalty.Value)
				assert.Equal(t, int64(0), params.MaxTokens.Value)
				assert.Equal(t, float64(0), params.TopP.Value)
				assert.Equal(t, int64(0), params.TopK.Value)
			},
		},
		{
			input: `{"model": "gpt-3.5-turbo-instruct", "frequency_penalty": null}`,
			check: func(t *testing.T, params openai.CompletionNewParams) {
				assert.Equal(t, false, params.FrequencyPenalty.IsPresent())
				assert.Equal(t, true, params.FrequencyPenalty.IsNull())

				buf, err := json.Marshal(params)
				assert.NoError(t, err)
				assert.Contains(t, string(buf), `"frequency_penalty":null`)
				assert.Equal(t, string(buf), `{"model":"gpt-3.5-turbo-instruct","frequency_penalty":null,"stream":false}`)
			},
		},
		{
			input: `{"model": "gpt-3.5-turbo-instruct"}`,
			check: func(t *testing.T, params openai.CompletionNewParams) {
				assert.Equal(t, false, params.FrequencyPenalty.IsPresent())
				assert.Equal(t, false, params.FrequencyPenalty.IsNull())

				buf, err := json.Marshal(params)
				assert.NoError(t, err)
				assert.NotContains(t, string(buf), `"frequency_penalty":null`)
				assert.Equal(t, string(buf), `{"model":"gpt-3.5-turbo-instruct","stream":false}`)
				assert.NotContains(t, string(buf), `"multi_modal_data":`)
			},
		},
		{
			input: `{"model":"gpt-3.5-turbo-instruct","multi_modal_data":[{"type":"image_url","image_url":{"url":"blah"}}]}`,
			check: func(t *testing.T, params openai.CompletionNewParams) {
				assert.Equal(t, false, params.FrequencyPenalty.IsPresent())
				assert.Equal(t, false, params.FrequencyPenalty.IsNull())

				buf, err := json.Marshal(params)
				assert.NoError(t, err)
				assert.Contains(t, string(buf), `"multi_modal_data":[{`)
				assert.Contains(t, string(buf), `"image_url":{"url":"blah"}`)
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			params := openai.CompletionNewParams{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)
			tt.check(t, params)
		})
	}
}

func TestUnmarshal_CompletionNewParamsPromptUnion(t *testing.T) {
	tests := []struct {
		input string
		check func(t *testing.T, msg openai.CompletionNewParamsPromptUnion)
	}{
		{
			input: `"chatgpt"`,
			check: func(t *testing.T, params openai.CompletionNewParamsPromptUnion) {
				assert.Equal(t, "chatgpt", params.OfString.Value)
				assert.Equal(t, []string(nil), params.OfArrayOfStrings)
				assert.Equal(t, []int64(nil), params.OfArrayOfTokens)
				assert.Equal(t, [][]int64(nil), params.OfArrayOfTokenArrays)
			},
		},
		{
			input: `["chatgpt","deepseek"]`,
			check: func(t *testing.T, params openai.CompletionNewParamsPromptUnion) {
				assert.Equal(t, "", params.OfString.Value)
				assert.Equal(t, []string{"chatgpt", "deepseek"}, params.OfArrayOfStrings)
				assert.Equal(t, []int64(nil), params.OfArrayOfTokens)
				assert.Equal(t, [][]int64(nil), params.OfArrayOfTokenArrays)
			},
		},

		{
			input: `[12,19]`,
			check: func(t *testing.T, params openai.CompletionNewParamsPromptUnion) {
				assert.Equal(t, "", params.OfString.Value)
				assert.Equal(t, []string(nil), params.OfArrayOfStrings)
				assert.Equal(t, []int64{12, 19}, params.OfArrayOfTokens)
				assert.Equal(t, [][]int64(nil), params.OfArrayOfTokenArrays)
			},
		},

		{
			input: `[[99,1],[1,2,3],[4]]`,
			check: func(t *testing.T, params openai.CompletionNewParamsPromptUnion) {
				assert.Equal(t, "", params.OfString.Value)
				assert.Equal(t, []string(nil), params.OfArrayOfStrings)
				assert.Equal(t, []int64(nil), params.OfArrayOfTokens)
				assert.Equal(t, [][]int64{{99, 1}, {1, 2, 3}, {4}}, params.OfArrayOfTokenArrays)
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			params := openai.CompletionNewParamsPromptUnion{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)
			tt.check(t, params)
		})
	}
}

func TestUnmarshal_CompletionNewParamsStopUnion(t *testing.T) {
	tests := []struct {
		input string
		check func(t *testing.T, msg openai.CompletionNewParamsStopUnion)
	}{
		{
			input: `"chatgpt"`,
			check: func(t *testing.T, params openai.CompletionNewParamsStopUnion) {
				assert.Equal(t, "chatgpt", params.OfString.Value)
				assert.Equal(t, []string(nil), params.OfCompletionNewsStopArray)
			},
		},
		{
			input: `["baidu", "chatgpt"]`,
			check: func(t *testing.T, params openai.CompletionNewParamsStopUnion) {
				assert.Equal(t, "", params.OfString.Value)
				assert.Equal(t, []string{"baidu", "chatgpt"}, params.OfCompletionNewsStopArray)
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			params := openai.CompletionNewParamsStopUnion{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)
			tt.check(t, params)
		})
	}
}

func TestUnmarshal_CompletionNullFields(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(*testing.T, openai.CompletionNewParams)
	}{
		{
			name: "null_prompt",
			input: `{
                "model": "gpt-3",
                "prompt": null,
                "max_tokens": 50
            }`,
			check: func(t *testing.T, params openai.CompletionNewParams) {
				assert.True(t, params.Prompt.IsNull())
			},
		},
		{
			name: "null_stop",
			input: `{
                "model": "gpt-3",
                "prompt": "Hello",
                "stop": null
            }`,
			check: func(t *testing.T, params openai.CompletionNewParams) {
				assert.True(t, params.Stop.IsNull())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := openai.CompletionNewParams{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)
			tt.check(t, params)
		})
	}
}

func TestUnmarshal_Completion(t *testing.T) {
	tests := []struct {
		input string
		check func(t *testing.T, msg openai.Completion)
	}{
		{
			input: `{"id":"cmpl-uqkvlQyYK7bGYrRHQ0eXlWi7","object":"text_completion","created":1589478378,"model":"gpt-3.5-turbo-instruct","system_fingerprint":"fp_44709d6fcb","choices":[{"text":"\n\nThis is indeed a test","index":0,"logprobs":null,"finish_reason":"length"}],"usage":{"prompt_tokens":5,"completion_tokens":7,"total_tokens":12}}`,
			check: func(t *testing.T, params openai.Completion) {
				assert.Equal(t, params.Usage.TotalTokens, int64(12))
				assert.Equal(t, params.Usage.PromptTokens, int64(5))
				assert.Equal(t, params.Usage.CompletionTokens, int64(7))
				assert.Equal(t, params.ID, "cmpl-uqkvlQyYK7bGYrRHQ0eXlWi7")

				buf, err := json.Marshal(params)
				assert.NoError(t, err)
				assert.Contains(t, string(buf), `"id":"cmpl-uqkvlQyYK7bGYrRHQ0eXlWi7"`)
				assert.Contains(t, string(buf), `"usage":{"completion_tokens":7,"prompt_tokens":5,"total_tokens":12}`)
				assert.NotContains(t, string(buf), `"logprobs":`)
				//assert.Equal(t, tt.input, string(buf))
			},
		},
		{
			input: `{"id":"cmpl-uqkvlQyYK7bGYrRHQ0eXlWi7","usage":null}`,
			check: func(t *testing.T, params openai.Completion) {
				assert.Equal(t, params.ID, "cmpl-uqkvlQyYK7bGYrRHQ0eXlWi7")
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			params := openai.Completion{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)
			tt.check(t, params)
		})
	}
}
