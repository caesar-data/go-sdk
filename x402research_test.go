// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package caesar_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/caesar-data/go-sdk"
	"github.com/caesar-data/go-sdk/internal/testutil"
	"github.com/caesar-data/go-sdk/option"
)

func TestX402ResearchNewWithOptionalParams(t *testing.T) {
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
	_, err := client.X402.Research.New(context.TODO(), caesar.X402ResearchNewParams{
		Query:        "query",
		XPayment:     "X-PAYMENT",
		ComputeUnits: caesar.Int(1),
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
