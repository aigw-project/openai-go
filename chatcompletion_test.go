// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openai_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"testing"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/internal/testutil"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/shared"

	"github.com/stretchr/testify/assert"
)

func TestChatCompletionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{{
			OfDeveloper: &openai.ChatCompletionDeveloperMessageParam{
				Content: openai.ChatCompletionDeveloperMessageParamContentUnion{
					OfString: openai.String("string"),
				},
				Name: openai.String("name"),
			},
		}},
		Model: shared.ChatModelO3Mini,
		Audio: openai.ChatCompletionAudioParam{
			Format: openai.ChatCompletionAudioParamFormatWAV,
			Voice:  openai.ChatCompletionAudioParamVoiceAlloy,
		},
		FrequencyPenalty: openai.Float(-2),
		FunctionCall: openai.ChatCompletionNewParamsFunctionCallUnion{
			OfFunctionCallMode: openai.String("none"),
		},
		Functions: []openai.ChatCompletionNewParamsFunction{{
			Name:        "name",
			Description: openai.String("description"),
			Parameters: shared.FunctionParameters{
				"foo": "bar",
			},
		}},
		LogitBias: map[string]int64{
			"foo": 0,
		},
		Logprobs:            openai.Bool(true),
		MaxCompletionTokens: openai.Int(0),
		MaxTokens:           openai.Int(0),
		Metadata: shared.MetadataParam{
			"foo": "string",
		},
		Modalities:        []string{"text"},
		N:                 openai.Int(1),
		ParallelToolCalls: openai.Bool(true),
		Prediction: openai.ChatCompletionPredictionContentParam{
			Content: openai.ChatCompletionPredictionContentContentUnionParam{
				OfString: openai.String("string"),
			},
		},
		PresencePenalty: openai.Float(-2),
		ReasoningEffort: shared.ReasoningEffortLow,
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfText: &shared.ResponseFormatTextParam{},
		},
		Seed:        openai.Int(-9007199254740991),
		ServiceTier: openai.ChatCompletionNewParamsServiceTierAuto,
		Stop: openai.ChatCompletionNewParamsStopUnion{
			OfString: openai.String("\n"),
		},
		Store: openai.Bool(true),
		StreamOptions: openai.ChatCompletionStreamOptionsParam{
			IncludeUsage: openai.Bool(true),
		},
		Temperature: openai.Float(1),
		ToolChoice: openai.ChatCompletionToolChoiceOptionUnionParam{
			OfAuto: openai.String("none"),
		},
		Tools: []openai.ChatCompletionToolParam{{
			Function: shared.FunctionDefinitionParam{
				Name:        "name",
				Description: openai.String("description"),
				Parameters: shared.FunctionParameters{
					"foo": "bar",
				},
				Strict: openai.Bool(true),
			},
		}},
		TopLogprobs: openai.Int(0),
		TopP:        openai.Float(1),
		User:        openai.String("user-1234"),
		WebSearchOptions: openai.ChatCompletionNewParamsWebSearchOptions{
			SearchContextSize: "low",
			UserLocation: openai.ChatCompletionNewParamsWebSearchOptionsUserLocation{
				Approximate: openai.ChatCompletionNewParamsWebSearchOptionsUserLocationApproximate{
					City:     openai.String("city"),
					Country:  openai.String("country"),
					Region:   openai.String("region"),
					Timezone: openai.String("timezone"),
				},
			},
		},
	})
	if err != nil {
		var apierr *openai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatCompletionGet(t *testing.T) {
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
	_, err := client.Chat.Completions.Get(context.TODO(), "completion_id")
	if err != nil {
		var apierr *openai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatCompletionUpdate(t *testing.T) {
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
	_, err := client.Chat.Completions.Update(
		context.TODO(),
		"completion_id",
		openai.ChatCompletionUpdateParams{
			Metadata: shared.MetadataParam{
				"foo": "string",
			},
		},
	)
	if err != nil {
		var apierr *openai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatCompletionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Chat.Completions.List(context.TODO(), openai.ChatCompletionListParams{
		After: openai.String("after"),
		Limit: openai.Int(0),
		Metadata: shared.MetadataParam{
			"foo": "string",
		},
		Model: openai.String("model"),
		Order: openai.ChatCompletionListParamsOrderAsc,
	})
	if err != nil {
		var apierr *openai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatCompletionDelete(t *testing.T) {
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
	_, err := client.Chat.Completions.Delete(context.TODO(), "completion_id")
	if err != nil {
		var apierr *openai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUnmarshal_ChatCompletion(t *testing.T) {
	tests := []struct {
		input string
		check func(t *testing.T, msg openai.ChatCompletionNewParams)
	}{
		{
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
                "temperature": 0.7,
                "max_tokens": 100,
                "stream": true
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, int64(100), params.MaxTokens.Value)
				assert.Equal(t, float64(0.7), params.Temperature.Value)
				assert.Equal(t, "user", *params.Messages[0].GetRole())
				assert.Equal(t, "Hello!", *(params.Messages[0].GetContent().AsAny().(*string)))
			},
		},
		{
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
				"prediction": {
					"type": "content",
					"content": "Hello!"
				}
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "content", string(params.Prediction.Type))
				assert.Equal(t, "Hello!", params.Prediction.Content.OfString.Value)
			},
		},
		{
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
				"prediction": {
					"type": "content",
					"content": [{"type":"text","text":"Hello!"}]
				}
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "content", string(params.Prediction.Type))
				assert.Equal(t, "Hello!", params.Prediction.Content.OfArrayOfContentParts[0].Text)
			},
		},
		{
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
				"tool_choice": "auto"
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "auto", params.ToolChoice.OfAuto.Value)
			},
		},
		{
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
				"tool_choice": {
					"type": "function",
					"function": {"name": "get_weather"}
				}
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "function", *params.ToolChoice.GetType())
				assert.Equal(t, "get_weather", params.ToolChoice.GetFunction().Name)
			},
		},
		{
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
				"function_call": "auto"
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "auto", params.FunctionCall.OfFunctionCallMode.Value)
			},
		},
		{
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
				"function_call": {
					"name": "get_weather"
				}
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "get_weather", *params.FunctionCall.GetName())
			},
		},
		{
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
				"response_format": {
					"type": "text"
				}
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "text", *params.ResponseFormat.GetType())
			},
		},
		{
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
				"response_format": {
					"type": "json_object"
				}
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "json_object", *params.ResponseFormat.GetType())
			},
		},
		{
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
				"response_format": {
					"type": "json_schema",
					"json_schema": {
						"name": "blah"
					}
				}
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "json_schema", *params.ResponseFormat.GetType())
				assert.Equal(t, "blah", params.ResponseFormat.GetJSONSchema().Name)
			},
		},
		{
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
				"stop": "stop"
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "stop", params.Stop.OfString.Value)
			},
		},
		{
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
				"stop": ["stop", "here"]
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, []string{"stop", "here"}, params.Stop.OfChatCompletionNewsStopArray)
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			params := openai.ChatCompletionNewParams{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)
			tt.check(t, params)
		})
	}
}

func TestUnmarshal_ChatCompletionMessage(t *testing.T) {
	tests := []struct {
		input string
		check func(t *testing.T, msg openai.ChatCompletionMessageParamUnion)
	}{
		{
			input: `{"role": "developer", "content": "blah"}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				c := *(msg.GetContent().AsAny().(*string))
				assert.Equal(t, "blah", c)
			},
		},
		{
			input: `{"role": "developer", "content": [{"type":"text","text":"blah"}]}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				c := *(msg.GetContent().AsAny().(*[]openai.ChatCompletionContentPartTextParam))
				assert.Equal(t, 1, len(c))
				assert.Equal(t, "text", string(c[0].Type))
				assert.Equal(t, "blah", c[0].Text)
			},
		},
		{
			input: `{"role": "system", "content": "blah"}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				c := *(msg.GetContent().AsAny().(*string))
				assert.Equal(t, "blah", c)
			},
		},
		{
			input: `{"role": "system", "content": [{"type":"text","text":"blah"}]}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				c := *(msg.GetContent().AsAny().(*[]openai.ChatCompletionContentPartTextParam))
				assert.Equal(t, 1, len(c))
				assert.Equal(t, "text", string(c[0].Type))
				assert.Equal(t, "blah", c[0].Text)
			},
		},
		{
			input: `{"role": "system"}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				c := msg.GetContent().AsAny()
				assert.Nil(t, c)
			},
		},
		{
			input: `{"role": "user", "content": "Hello!"}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				assert.Equal(t, "user", *msg.GetRole())
				assert.Equal(t, "Hello!", *(msg.GetContent().AsAny().(*string)))
			},
		},
		{
			input: `{"role": "user", "content": [{"type":"text","text":"blah"}]}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				c := *(msg.GetContent().AsAny().(*[]openai.ChatCompletionContentPartUnionParam))
				assert.Equal(t, 1, len(c))
				assert.Equal(t, "text", *c[0].GetType())
				assert.Equal(t, "blah", *c[0].GetText())
			},
		},
		{
			input: `{"role": "user", "content": [{"type":"image_url","image_url":{"url":"blah"}}]}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				c := *(msg.GetContent().AsAny().(*[]openai.ChatCompletionContentPartUnionParam))
				assert.Equal(t, 1, len(c))
				assert.Equal(t, "image_url", *c[0].GetType())
				assert.Equal(t, "blah", c[0].GetImageURL().URL)
			},
		},
		{
			input: `{"role": "user", "content": [{"type":"audio_url","audio_url":{"url":"blah"}}]}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				c := *(msg.GetContent().AsAny().(*[]openai.ChatCompletionContentPartUnionParam))
				assert.Equal(t, 1, len(c))
				assert.Equal(t, "audio_url", *c[0].GetType())
				assert.Equal(t, "blah", c[0].GetAudioURL().URL)
			},
		},
		{
			input: `{"role": "user", "content": [{"type":"input_audio","input_audio":{"format":"blah"}}]}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				c := *(msg.GetContent().AsAny().(*[]openai.ChatCompletionContentPartUnionParam))
				assert.Equal(t, 1, len(c))
				assert.Equal(t, "input_audio", *c[0].GetType())
				assert.Equal(t, "blah", c[0].GetInputAudio().Format)
			},
		},
		{
			input: `{"role": "user", "content": [{"type":"file","file":{"filename":"blah"}}]}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				c := *(msg.GetContent().AsAny().(*[]openai.ChatCompletionContentPartUnionParam))
				assert.Equal(t, 1, len(c))
				assert.Equal(t, "file", *c[0].GetType())
				assert.Equal(t, "blah", c[0].GetFile().Filename.Value)
			},
		},
		{
			input: `{"role": "assistant", "content": "Hello!"}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				assert.Equal(t, "assistant", *msg.GetRole())
				assert.Equal(t, "Hello!", *(msg.GetContent().AsAny().(*string)))
			},
		},
		{
			input: `{"role": "assistant", "content": [{"type":"text","text":"Hello!"}]}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				assert.Equal(t, "assistant", *msg.GetRole())
				c := *(msg.GetContent().AsAny().(*[]openai.ChatCompletionAssistantMessageParamContentArrayOfContentPartUnion))
				assert.Equal(t, 1, len(c))
				assert.Equal(t, "Hello!", *c[0].GetText())
			},
		},
		{
			input: `{"role": "assistant", "content": [{"type":"refusal","refusal":"Hello!"}]}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				assert.Equal(t, "assistant", *msg.GetRole())
				c := *(msg.GetContent().AsAny().(*[]openai.ChatCompletionAssistantMessageParamContentArrayOfContentPartUnion))
				assert.Equal(t, 1, len(c))
				assert.Equal(t, "Hello!", *c[0].GetRefusal())
			},
		},
		{
			input: `{"role": "tool", "content": "Hello!", "tool_call_id":"call"}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				assert.Equal(t, "tool", *msg.GetRole())
				assert.Equal(t, "Hello!", *(msg.GetContent().AsAny().(*string)))
			},
		},
		{
			input: `{"role": "tool", "content": [{"type":"text","text":"Hello!"}], "tool_call_id":"call"}`,
			check: func(t *testing.T, msg openai.ChatCompletionMessageParamUnion) {
				c := *(msg.GetContent().AsAny().(*[]openai.ChatCompletionContentPartTextParam))
				assert.Equal(t, 1, len(c))
				assert.Equal(t, "text", string(c[0].Type))
			},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var msg openai.ChatCompletionMessageParamUnion
			err := json.Unmarshal([]byte(tt.input), &msg)
			assert.NoError(t, err)
			tt.check(t, msg)
		})
	}

	// multiple msgs
	input := `[
		{"role": "user", "content": "Hello!"},
		{"role":"user", "content":[{"type":"text","text":"Hello!"}, {"type":"image_url","image_url":{"url":"blah"}}]},
		{"role": "user", "content": "Hi!"}
	]`
	var msg []openai.ChatCompletionMessageParamUnion
	err := json.Unmarshal([]byte(input), &msg)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(msg))
	assert.Equal(t, "user", *msg[0].GetRole())
	assert.Equal(t, "Hello!", *(msg[0].GetContent().AsAny().(*string)))
	assert.Equal(t, "user", *msg[1].GetRole())
	assert.Equal(t, 2, len(*(msg[1].GetContent().AsAny().(*[]openai.ChatCompletionContentPartUnionParam))))
	assert.Equal(t, "text", *(*msg[1].GetContent().AsAny().(*[]openai.ChatCompletionContentPartUnionParam))[0].GetType())
	assert.Equal(t, "Hello!", *(*msg[1].GetContent().AsAny().(*[]openai.ChatCompletionContentPartUnionParam))[0].GetText())
	assert.Equal(t, "image_url", *(*msg[1].GetContent().AsAny().(*[]openai.ChatCompletionContentPartUnionParam))[1].GetType())
	assert.Equal(t, "blah", (*msg[1].GetContent().AsAny().(*[]openai.ChatCompletionContentPartUnionParam))[1].GetImageURL().URL)
	assert.Equal(t, "Hi!", *(msg[2].GetContent().AsAny().(*string)))

	// bad role
	input = `{"role": "unknown", "content": "Hello!"}`
	var param openai.ChatCompletionMessageParamUnion
	err = json.Unmarshal([]byte(input), &param)
	assert.Equal(t, "apijson: can not find the variant has discriminator value unknown", err.Error())

	// bad content
	input = `{"role": "user", "content": 3}`
	param = openai.ChatCompletionMessageParamUnion{}
	err = json.Unmarshal([]byte(input), &param)
	assert.Equal(t, "apijson: was not able to coerce type as union", err.Error())
}

func TestUnmarshal_IncludeUsage(t *testing.T) {
	input := `{
		"model": "gpt-4",
		"messages": [
			{"role": "user", "content": "Hello!"}
		]
		"stream": true
	}`
	params := openai.ChatCompletionNewParams{}
	json.Unmarshal([]byte(input), &params)
	assert.False(t, params.UsageDisabledByUser())
	params.DisableUsageByUser()
	assert.True(t, params.UsageDisabledByUser())
}

func TestUnmarshal_CorrectStatusForOption(t *testing.T) {
	input := `{
		"model": "gpt-4",
		"messages": [
			{"role": "user", "content": "Hello!"}
		],
		"stream_options": {
			"include_usage": false
		},
		"stream": true
	}`
	params := openai.ChatCompletionNewParams{}
	json.Unmarshal([]byte(input), &params)
	assert.False(t, params.StreamOptions.IncludeUsage.IsOmitted())
	assert.True(t, params.StreamOptions.IncludeUsage.IsPresent())

	input = `{
		"model": "gpt-4",
		"messages": [
			{"role": "user", "content": "Hello!"}
		],
		"stream": true
	}`
	params = openai.ChatCompletionNewParams{}
	json.Unmarshal([]byte(input), &params)
	assert.True(t, params.StreamOptions.IncludeUsage.IsOmitted())
	assert.False(t, params.StreamOptions.IncludeUsage.IsPresent())
}

func TestUnmarshal_DecodeStrictly(t *testing.T) {
	input := `{"messages":[{"content":"you are a cool man","role":"system"}],"model":"aigw-s2-test-qwq","temperature":1,"top_p":1,"stop":[123],"stream_options":{"include_usage":true},"stream":true}`
	params := openai.ChatCompletionNewParams{}
	err := json.Unmarshal([]byte(input), &params)
	assert.Equal(t, "apijson: was not able to coerce type as union", err.Error())
}

func TestUnmarshal_ChatCompletionNullField(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(*testing.T, openai.ChatCompletionNewParams)
	}{
		{
			name: "null_stop",
			input: `{
                "model": "gpt-4",
                "messages": [{"role": "user", "content": "Hello"}],
                "stop": null
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.True(t, params.Stop.IsNull())
			},
		},
		{
			name: "null_function_call",
			input: `{
                "model": "gpt-4",
                "messages": [{"role": "user", "content": "Hello"}],
                "function_call": null
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.True(t, params.FunctionCall.IsNull())
			},
		},
		{
			name: "null_tool_choice",
			input: `{
                "model": "gpt-4",
                "messages": [{"role": "user", "content": "Hello"}],
                "tool_choice": null
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.True(t, params.ToolChoice.IsNull())
			},
		},
		{
			name: "null_prediction_content",
			input: `{
                "model": "gpt-4",
                "messages": [{"role": "user", "content": "Hello"}],
                "prediction": {
                    "type": "content",
                    "content": null
                }
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.True(t, params.Prediction.Content.IsNull())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := openai.ChatCompletionNewParams{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)
			tt.check(t, params)
		})
	}
}

func strptr(s string) *string {
	return &s
}

func TestForAIGW_Marshal_ChatCompletionChunkNullContents(t *testing.T) {
	chunk := openai.ChatCompletionChunk{
		ID: "chunk_id",
		Choices: []openai.ChatCompletionChunkChoice{
			{
				Index: 0,
				Delta: openai.ChatCompletionChunkChoiceDelta{
					Role: "assistant",
				},
			},
		},
	}
	data, _ := json.Marshal(chunk)
	assert.Contains(t, string(data), `"content":null`)
	assert.Contains(t, string(data), `"reasoning_content":null`)

	chunk.Choices[0].Delta.Content = strptr("hello")
	data, _ = json.Marshal(chunk)
	assert.Contains(t, string(data), `"content":"hello"`)
	assert.Contains(t, string(data), `"reasoning_content":null`)

	chunk.Choices[0].Delta.ReasoningContent = strptr("hello")
	data, _ = json.Marshal(chunk)
	assert.Contains(t, string(data), `"content":"hello"`)
	assert.Contains(t, string(data), `"reasoning_content":"hello"`)

	chunk.Choices[0].Delta.Content = nil
	data, _ = json.Marshal(chunk)
	assert.Contains(t, string(data), `"content":null`)
	assert.Contains(t, string(data), `"reasoning_content":"hello"`)
}

func TestForAIGW_ChatCompletion_ChatTemplateKwargs(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(*testing.T, openai.ChatCompletionNewParams)
	}{
		{
			name: "not_exists",
			input: `{
                "model": "gpt-4",
                "messages": [{"role": "user", "content": "Hello"}]
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.True(t, params.ChatTemplateKwargs == nil)
			},
		},
		{
			name: "null",
			input: `{
                "model": "gpt-4",
                "messages": [{"role": "user", "content": "Hello"}],
                "chat_template_kwargs": null
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.True(t, params.ChatTemplateKwargs == nil)
			},
		},
		{
			name: "empty",
			input: `{
                "model": "gpt-4",
                "messages": [{"role": "user", "content": "Hello"}],
                "chat_template_kwargs": {}
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.True(t, params.ChatTemplateKwargs != nil)
				assert.True(t, len(params.ChatTemplateKwargs) == 0)
			},
		},
		{
			name: "normal",
			input: `{
                "model": "gpt-4",
                "messages": [{"role": "user", "content": "Hello"}],
                "chat_template_kwargs": {
					"bool": true,
					"key": "value"
				}
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.True(t, params.ChatTemplateKwargs != nil)
				assert.True(t, len(params.ChatTemplateKwargs) > 0)
				assert.Equal(t, true, params.ChatTemplateKwargs["bool"])
				assert.Equal(t, "value", params.ChatTemplateKwargs["key"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := openai.ChatCompletionNewParams{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)
			tt.check(t, params)

			// check if marshalling and unmarshalling works, and got the same value
			data, err := json.Marshal(params)
			assert.NoError(t, err)
			var params2 openai.ChatCompletionNewParams
			err = json.Unmarshal(data, &params2)
			assert.NoError(t, err)
			assert.Equal(t, params.ChatTemplateKwargs, params2.ChatTemplateKwargs)
		})
	}
}

func TestForAIGW_ChatCompletion_MmSamplingKwargs(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(*testing.T, openai.ChatCompletionNewParams)
	}{
		{
			name: "not_exists",
			input: `{
                "model": "gpt-4",
                "messages": [{"role": "user", "content": "Hello"}]
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.True(t, params.ChatTemplateKwargs == nil)
			},
		},
		{
			name: "null",
			input: `{
                "model": "gpt-4",
                "messages": [{"role": "user", "content": "Hello"}],
                "mm_sampling_kwargs": null
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.True(t, params.MmSamplingKwargs == nil)
			},
		},
		{
			name: "empty",
			input: `{
                "model": "gpt-4",
                "messages": [{"role": "user", "content": "Hello"}],
                "mm_sampling_kwargs": {}
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.True(t, params.MmSamplingKwargs != nil)
				assert.True(t, len(params.MmSamplingKwargs) == 0)
			},
		},
		{
			name: "normal",
			input: `{
                "model": "gpt-4",
                "messages": [{"role": "user", "content": "Hello"}],
                "mm_sampling_kwargs": {
					"bool": true,
					"key": "value"
				}
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.True(t, params.MmSamplingKwargs != nil)
				assert.True(t, len(params.MmSamplingKwargs) > 0)
				assert.Equal(t, true, params.MmSamplingKwargs["bool"])
				assert.Equal(t, "value", params.MmSamplingKwargs["key"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := openai.ChatCompletionNewParams{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)
			tt.check(t, params)

			// check if marshalling and unmarshalling works, and got the same value
			data, err := json.Marshal(params)
			assert.NoError(t, err)
			var params2 openai.ChatCompletionNewParams
			err = json.Unmarshal(data, &params2)
			assert.NoError(t, err)
			assert.Equal(t, params.ChatTemplateKwargs, params2.ChatTemplateKwargs)
		})
	}
}

func TestForAIGW_ChatCompletion_UserContent_ImageUrl(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(*testing.T, openai.ChatCompletionNewParams)
	}{
		{
			name: "sanity",
			input: `{
                "model": "gpt-4",
                "messages": [{
					"role": "user",
					"content": [{"type":"image_url","image_url":{"url":"image_url","detail":"string"}}]
				}]
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "image_url", params.Messages[0].OfUser.Content.OfArrayOfContentParts[0].GetImageURL().URL)
			},
		},
		{
			name: "string",
			input: `{
                "model": "gpt-4",
                "messages": [{
					"role": "user",
					"content": [{"type":"image_url","image_url":"urlstring"}]
				}]
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "urlstring", params.Messages[0].OfUser.Content.OfArrayOfContentParts[0].OfImageURL.ImageURL.OfString.String())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := openai.ChatCompletionNewParams{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)
			tt.check(t, params)

			// check if marshalling and unmarshalling works, and got the same value
			data, err := json.Marshal(params)
			assert.NoError(t, err)
			var params2 openai.ChatCompletionNewParams
			err = json.Unmarshal(data, &params2)
			assert.NoError(t, err)
			tt.check(t, params2)
		})
	}
}

func TestForAIGW_ChatCompletion_UserContent_AudioUrl(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(*testing.T, openai.ChatCompletionNewParams)
	}{
		{
			name: "sanity",
			input: `{
                "model": "gpt-4",
                "messages": [{
					"role": "user",
					"content": [{"type":"audio_url","audio_url":{"url":"audio_url"}}]
				}]
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "audio_url", params.Messages[0].OfUser.Content.OfArrayOfContentParts[0].GetAudioURL().URL)
			},
		},
		{
			name: "string",
			input: `{
                "model": "gpt-4",
                "messages": [{
					"role": "user",
					"content": [{"type":"audio_url","audio_url":"urlstring"}]
				}]
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "urlstring", params.Messages[0].OfUser.Content.OfArrayOfContentParts[0].OfAudioURL.AudioURL.OfString.String())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := openai.ChatCompletionNewParams{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)
			tt.check(t, params)

			// check if marshalling and unmarshalling works, and got the same value
			data, err := json.Marshal(params)
			assert.NoError(t, err)
			var params2 openai.ChatCompletionNewParams
			err = json.Unmarshal(data, &params2)
			assert.NoError(t, err)
			tt.check(t, params2)
		})
	}
}

func TestForAIGW_ChatCompletion_UserContent_VideoUrl(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(*testing.T, openai.ChatCompletionNewParams)
	}{
		{
			name: "sanity",
			input: `{
                "model": "gpt-4",
                "messages": [{
					"role": "user",
					"content": [{"type":"video_url","video_url":{"url":"video_url"}}]
				}]
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "video_url", params.Messages[0].OfUser.Content.OfArrayOfContentParts[0].GetVideoURL().URL)
			},
		},
		{
			name: "string",
			input: `{
                "model": "gpt-4",
                "messages": [{
					"role": "user",
					"content": [{"type":"video_url","video_url":"urlstring"}]
				}]
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "urlstring", params.Messages[0].OfUser.Content.OfArrayOfContentParts[0].OfVideoURL.VideoURL.OfString.String())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := openai.ChatCompletionNewParams{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)
			tt.check(t, params)

			// check if marshalling and unmarshalling works, and got the same value
			data, err := json.Marshal(params)
			assert.NoError(t, err)
			var params2 openai.ChatCompletionNewParams
			err = json.Unmarshal(data, &params2)
			assert.NoError(t, err)
			tt.check(t, params2)
		})
	}
}

func TestForAIGW_ChatCompletion_RepetitionPenalty(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		repetitionPenalty float64
		exists            bool
	}{
		{
			name: "non-exists",
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ]
            }`,
			exists: false,
		},
		{
			name: "string",
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
				"repetition_penalty": 1.2
            }`,
			repetitionPenalty: 1.2,
			exists:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := openai.ChatCompletionNewParams{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)

			// check value
			if tt.exists {
				assert.Equal(t, tt.repetitionPenalty, params.RepetitionPenalty.Value)
			}

			// check if marshalling and unmarshalling works, and got the same value
			data, err := json.Marshal(params)
			assert.NoError(t, err)

			if tt.exists {
				assert.Contains(t, string(data), `"repetition_penalty":`)
			} else {
				assert.NotContains(t, string(data), `"repetition_penalty":`)
			}

			var params2 openai.ChatCompletionNewParams
			err = json.Unmarshal(data, &params2)
			assert.NoError(t, err)

			// check value again
			if tt.exists {
				assert.Equal(t, tt.repetitionPenalty, params.RepetitionPenalty.Value)
			}
		})
	}
}

func TestForAIGW_ChatCompletion_SkipSpecialTokens(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		skipSpecialToken bool
		exists           bool
	}{
		{
			name: "non-exists",
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ]
            }`,
			exists: false,
		},
		{
			name: "string",
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "user", "content": "Hello!"}
                ],
				"skip_special_tokens": true
            }`,
			skipSpecialToken: true,
			exists:           true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := openai.ChatCompletionNewParams{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)

			// check value
			if tt.exists {
				assert.Equal(t, tt.skipSpecialToken, params.SkipSpecialTokens.Value)
			}

			// check if marshalling and unmarshalling works, and got the same value
			data, err := json.Marshal(params)
			assert.NoError(t, err)

			if tt.exists {
				assert.Contains(t, string(data), `"skip_special_tokens":`)
			} else {
				assert.NotContains(t, string(data), `"skip_special_tokens":`)
			}

			var params2 openai.ChatCompletionNewParams
			err = json.Unmarshal(data, &params2)
			assert.NoError(t, err)

			// check value again
			if tt.exists {
				assert.Equal(t, tt.skipSpecialToken, params.SkipSpecialTokens.Value)
			}
		})
	}
}

func TestForAIGW_ChatCompletion_Message_Tool(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(t *testing.T, params openai.ChatCompletionNewParams)
	}{
		{
			name: "string",
			input: `{
                "model": "gpt-4",
                "messages": [
                    {"role": "tool", "content": "Hello!"}
                ]
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "tool", string(params.Messages[0].OfTool.Role))
				assert.Equal(t, "Hello!", params.Messages[0].OfTool.Content.OfString.String())
			},
		},
		{
			name: "array",
			input: `{
                "model": "gpt-4",
                "messages": [
                    {
						"role": "tool",
						"content": [{
							"type": "string",
							"text": "Hello!"
						}]
					}
                ]
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "tool", string(params.Messages[0].OfTool.Role))
				assert.Equal(t, "string", string(params.Messages[0].OfTool.Content.OfArrayOfContentParts[0].Type))
				assert.Equal(t, "Hello!", params.Messages[0].OfTool.Content.OfArrayOfContentParts[0].Text)
			},
		},
		{
			name: "empty-array",
			input: `{
                "model": "gpt-4",
                "messages": [
                    {
						"role": "tool",
						"content": []
					}
                ]
            }`,
			check: func(t *testing.T, params openai.ChatCompletionNewParams) {
				assert.Equal(t, "tool", string(params.Messages[0].OfTool.Role))
				assert.Equal(t, 0, len(params.Messages[0].OfTool.Content.OfArrayOfContentParts))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := openai.ChatCompletionNewParams{}
			err := json.Unmarshal([]byte(tt.input), &params)
			assert.NoError(t, err)

			tt.check(t, params)

			// check if marshalling and unmarshalling works, and got the same value
			data, err := json.Marshal(params)
			assert.NoError(t, err)

			var params2 openai.ChatCompletionNewParams
			err = json.Unmarshal(data, &params2)
			assert.NoError(t, err)

			// check value again
			tt.check(t, params2)
		})
	}
}
