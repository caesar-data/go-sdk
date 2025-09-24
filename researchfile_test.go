// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package caesar_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/caesar-go"
	"github.com/stainless-sdks/caesar-go/internal/testutil"
	"github.com/stainless-sdks/caesar-go/option"
)

func TestResearchFileNew(t *testing.T) {
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
	_, err := client.Research.Files.New(context.TODO(), caesar.ResearchFileNewParams{
		File: io.Reader(bytes.NewBuffer([]byte("some file contents"))),
	})
	if err != nil {
		var apierr *caesar.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResearchFileListWithOptionalParams(t *testing.T) {
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
	_, err := client.Research.Files.List(context.TODO(), caesar.ResearchFileListParams{
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
