// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package deeptable_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/deeptable-go"
	"github.com/stainless-sdks/deeptable-go/internal/testutil"
	"github.com/stainless-sdks/deeptable-go/option"
)

func TestManualPagination(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := deeptable.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	page, err := client.Files.List(context.TODO(), deeptable.FileListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, file := range page.Data {
		t.Logf("%+v\n", file.ID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, file := range page.Data {
			t.Logf("%+v\n", file.ID)
		}
	}
}
