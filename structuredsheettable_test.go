// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package deeptable_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/deeptable-com/deeptable-go"
	"github.com/deeptable-com/deeptable-go/internal/testutil"
	"github.com/deeptable-com/deeptable-go/option"
)

func TestStructuredSheetTableGet(t *testing.T) {
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
	_, err := client.StructuredSheets.Tables.Get(
		context.TODO(),
		"tbl_01kfxgjd94fn9stqm45rqr2pnz",
		deeptable.StructuredSheetTableGetParams{
			StructuredSheetID: "ss_01kfxgjd94fn9stqm42nejb627",
		},
	)
	if err != nil {
		var apierr *deeptable.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStructuredSheetTableListWithOptionalParams(t *testing.T) {
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
	_, err := client.StructuredSheets.Tables.List(
		context.TODO(),
		"ss_01kfxgjd94fn9stqm42nejb627",
		deeptable.StructuredSheetTableListParams{
			After: deeptable.String("tbl_01kfxgjd94fn9stqm45rqr2pnz"),
			Limit: deeptable.Int(20),
		},
	)
	if err != nil {
		var apierr *deeptable.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStructuredSheetTableDownload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := deeptable.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.StructuredSheets.Tables.Download(
		context.TODO(),
		"tbl_01kfxgjd94fn9stqm45rqr2pnz",
		deeptable.StructuredSheetTableDownloadParams{
			StructuredSheetID: "ss_01kfxgjd94fn9stqm42nejb627",
			Format:            deeptable.StructuredSheetTableDownloadParamsFormatParquet,
		},
	)
	if err != nil {
		var apierr *deeptable.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *deeptable.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
