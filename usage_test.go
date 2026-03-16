// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package deeptable_test

import (
	"context"
	"os"
	"testing"

	"github.com/deeptable-com/deeptable-go"
	"github.com/deeptable-com/deeptable-go/internal/testutil"
	"github.com/deeptable-com/deeptable-go/option"
)

func TestUsage(t *testing.T) {
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
	structuredSheetResponse, err := client.StructuredSheets.New(context.TODO(), deeptable.StructuredSheetNewParams{
		FileID: "file_01h45ytscbebyvny4gc8cr8ma2",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", structuredSheetResponse.ID)
}
