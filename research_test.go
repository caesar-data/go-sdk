// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package caesar_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/caesar-go"
	"github.com/stainless-sdks/caesar-go/internal/testutil"
	"github.com/stainless-sdks/caesar-go/option"
)

func TestResearchNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := caesar.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Research.New(context.TODO(), caesar.ResearchNewParams{
		Query:        "Is lithium supply a bottleneck for EV adoption?",
		ComputeUnits: caesar.Int(1),
		Files:        []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		SystemPrompt: caesar.String("system_prompt"),
	})
	if err != nil {
		var apierr *caesar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResearchGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := caesar.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Research.Get(context.TODO(), "id")
	if err != nil {
		var apierr *caesar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResearchListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := caesar.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Research.List(context.TODO(), caesar.ResearchListParams{
		Limit: caesar.Int(1),
		Page:  caesar.Int(1),
	})
	if err != nil {
		var apierr *caesar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
