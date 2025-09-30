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

func TestAutoPagination(t *testing.T) {
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
	iter := client.Research.ListAutoPaging(context.TODO(), caesar.ResearchListParams{
		Limit: caesar.Int(30),
		Page:  caesar.Int(2),
	})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		research := iter.Current()
		t.Logf("%+v\n", research.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
