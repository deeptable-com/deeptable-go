// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package deeptable

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/deeptable-go/internal/apijson"
	"github.com/stainless-sdks/deeptable-go/internal/apiquery"
	"github.com/stainless-sdks/deeptable-go/internal/requestconfig"
	"github.com/stainless-sdks/deeptable-go/option"
	"github.com/stainless-sdks/deeptable-go/packages/pagination"
	"github.com/stainless-sdks/deeptable-go/packages/param"
	"github.com/stainless-sdks/deeptable-go/packages/respjson"
	"github.com/stainless-sdks/deeptable-go/shared/constant"
)

// Convert uploaded spreadsheets into structured data. Creates relational tables
// from messy spreadsheet data.
//
// StructuredSheetService contains methods and other services that help with
// interacting with the deeptable API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStructuredSheetService] method instead.
type StructuredSheetService struct {
	options []option.RequestOption
	// Convert uploaded spreadsheets into structured data. Creates relational tables
	// from messy spreadsheet data.
	Tables StructuredSheetTableService
}

// NewStructuredSheetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStructuredSheetService(opts ...option.RequestOption) (r StructuredSheetService) {
	r = StructuredSheetService{}
	r.options = opts
	r.Tables = NewStructuredSheetTableService(opts...)
	return
}

// Start converting a spreadsheet workbook into structured data. This initiates an
// asynchronous conversion process. Poll the returned resource using the `id` to
// check completion status.
func (r *StructuredSheetService) New(ctx context.Context, body StructuredSheetNewParams, opts ...option.RequestOption) (res *StructuredSheetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/structured-sheets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get the status and details of a structured sheet conversion.
func (r *StructuredSheetService) Get(ctx context.Context, structuredSheetID string, opts ...option.RequestOption) (res *StructuredSheetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if structuredSheetID == "" {
		err = errors.New("missing required structured_sheet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/structured-sheets/%s", url.PathEscape(structuredSheetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all structured sheets conversions for the authenticated user. Results are
// paginated using cursor-based pagination.
func (r *StructuredSheetService) List(ctx context.Context, query StructuredSheetListParams, opts ...option.RequestOption) (res *pagination.CursorIDPage[StructuredSheetResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/structured-sheets"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all structured sheets conversions for the authenticated user. Results are
// paginated using cursor-based pagination.
func (r *StructuredSheetService) ListAutoPaging(ctx context.Context, query StructuredSheetListParams, opts ...option.RequestOption) *pagination.CursorIDPageAutoPager[StructuredSheetResponse] {
	return pagination.NewCursorIDPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a structured sheet conversion and its associated exports. This action
// cannot be undone.
func (r *StructuredSheetService) Delete(ctx context.Context, structuredSheetID string, opts ...option.RequestOption) (res *StructuredSheetDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if structuredSheetID == "" {
		err = errors.New("missing required structured_sheet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/structured-sheets/%s", url.PathEscape(structuredSheetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Cancel a structured sheet conversion that is in progress. Only jobs with status
// 'queued' or 'in_progress' can be cancelled.
func (r *StructuredSheetService) Cancel(ctx context.Context, structuredSheetID string, opts ...option.RequestOption) (res *StructuredSheetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if structuredSheetID == "" {
		err = errors.New("missing required structured_sheet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/structured-sheets/%s/cancel", url.PathEscape(structuredSheetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Download the structured data in the specified format. Only available when
// conversion status is 'completed'.
//
// Available formats:
//
// - `sqlite`: SQLite database containing all extracted tables
// - `cell_labels`: CSV file with cell-level semantic labels
func (r *StructuredSheetService) Download(ctx context.Context, structuredSheetID string, query StructuredSheetDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/x-sqlite3")}, opts...)
	if structuredSheetID == "" {
		err = errors.New("missing required structured_sheet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/structured-sheets/%s/download", url.PathEscape(structuredSheetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Response representing a structured sheet conversion job.
//
// This is returned from POST (create), GET (retrieve), and list endpoints.
type StructuredSheetResponse struct {
	// The unique identifier for this structured sheet conversion.
	ID string `json:"id" api:"required"`
	// The timestamp when the conversion was started.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The unique identifier for the source file.
	FileID string `json:"file_id" api:"required"`
	// The object type, which is always 'structured_sheet'.
	Object constant.StructuredSheet `json:"object" api:"required"`
	// The current processing status.
	//
	// Any of "queued", "in_progress", "completed", "failed", "cancelled".
	Status StructuredSheetResponseStatus `json:"status" api:"required"`
	// The timestamp when the conversion was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Error information when processing fails.
	LastError StructuredSheetResponseLastError `json:"last_error" api:"nullable"`
	// List of sheet names included in this conversion.
	SheetNames []string `json:"sheet_names"`
	// Number of tables extracted from the workbook. Only present when status is
	// 'completed'.
	TableCount int64 `json:"table_count" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		FileID      respjson.Field
		Object      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		LastError   respjson.Field
		SheetNames  respjson.Field
		TableCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StructuredSheetResponse) RawJSON() string { return r.JSON.raw }
func (r *StructuredSheetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current processing status.
type StructuredSheetResponseStatus string

const (
	StructuredSheetResponseStatusQueued     StructuredSheetResponseStatus = "queued"
	StructuredSheetResponseStatusInProgress StructuredSheetResponseStatus = "in_progress"
	StructuredSheetResponseStatusCompleted  StructuredSheetResponseStatus = "completed"
	StructuredSheetResponseStatusFailed     StructuredSheetResponseStatus = "failed"
	StructuredSheetResponseStatusCancelled  StructuredSheetResponseStatus = "cancelled"
)

// Error information when processing fails.
type StructuredSheetResponseLastError struct {
	// A machine-readable error code.
	Code string `json:"code" api:"required"`
	// A human-readable description of the error.
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StructuredSheetResponseLastError) RawJSON() string { return r.JSON.raw }
func (r *StructuredSheetResponseLastError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from deleting a structured sheet.
//
// Following the OpenAI API convention for delete responses.
type StructuredSheetDeleteResponse struct {
	// The unique identifier of the deleted structured sheet.
	ID string `json:"id" api:"required"`
	// Whether the structured sheet was successfully deleted.
	Deleted bool `json:"deleted" api:"required"`
	// The object type, which is always 'structured_sheet'.
	Object constant.StructuredSheet `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StructuredSheetDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *StructuredSheetDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StructuredSheetNewParams struct {
	// The unique identifier of the file to convert.
	FileID string `json:"file_id" api:"required"`
	// List of sheet names to convert. If None, all sheets will be converted.
	SheetNames []string `json:"sheet_names,omitzero"`
	paramObj
}

func (r StructuredSheetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StructuredSheetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StructuredSheetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StructuredSheetListParams struct {
	// A cursor for pagination. Use the `last_id` from a previous response to fetch the
	// next page of results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of results to return per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StructuredSheetListParams]'s query parameters as
// `url.Values`.
func (r StructuredSheetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StructuredSheetDownloadParams struct {
	// The export format to download.
	//
	// Any of "sqlite", "cell_labels".
	Format StructuredSheetDownloadParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StructuredSheetDownloadParams]'s query parameters as
// `url.Values`.
func (r StructuredSheetDownloadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The export format to download.
type StructuredSheetDownloadParamsFormat string

const (
	StructuredSheetDownloadParamsFormatSqlite     StructuredSheetDownloadParamsFormat = "sqlite"
	StructuredSheetDownloadParamsFormatCellLabels StructuredSheetDownloadParamsFormat = "cell_labels"
)
