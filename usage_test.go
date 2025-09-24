// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package caesar_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/caesar-go"
	"github.com/stainless-sdks/caesar-go/internal/testutil"
	"github.com/stainless-sdks/caesar-go/option"
)

func TestUsage(t *testing.T) {
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
	research, err := client.Research.New(context.TODO(), caesar.ResearchNewParams{
		Query: "REPLACE_ME",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", research.ID)
}
