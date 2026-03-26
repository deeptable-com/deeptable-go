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

	"github.com/deeptable-com/deeptable-go/internal/apijson"
	"github.com/deeptable-com/deeptable-go/internal/apiquery"
	"github.com/deeptable-com/deeptable-go/internal/requestconfig"
	"github.com/deeptable-com/deeptable-go/option"
	"github.com/deeptable-com/deeptable-go/packages/pagination"
	"github.com/deeptable-com/deeptable-go/packages/param"
	"github.com/deeptable-com/deeptable-go/packages/respjson"
	"github.com/deeptable-com/deeptable-go/shared/constant"
)

// Convert uploaded spreadsheets into structured data. Creates relational tables
// from messy spreadsheet data.
//
// StructuredSheetTableService contains methods and other services that help with
// interacting with the deeptable API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStructuredSheetTableService] method instead.
type StructuredSheetTableService struct {
	options []option.RequestOption
}

// NewStructuredSheetTableService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStructuredSheetTableService(opts ...option.RequestOption) (r StructuredSheetTableService) {
	r = StructuredSheetTableService{}
	r.options = opts
	return
}

// Get details of a specific table extracted from the structured sheet. Only
// available when conversion status is 'completed'.
func (r *StructuredSheetTableService) Get(ctx context.Context, tableID string, query StructuredSheetTableGetParams, opts ...option.RequestOption) (res *TableResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.StructuredSheetID == "" {
		err = errors.New("missing required structured_sheet_id parameter")
		return nil, err
	}
	if tableID == "" {
		err = errors.New("missing required table_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/structured-sheets/%s/tables/%s", url.PathEscape(query.StructuredSheetID), url.PathEscape(tableID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all tables extracted from the structured sheet. Only available when
// conversion status is 'completed'.
func (r *StructuredSheetTableService) List(ctx context.Context, structuredSheetID string, query StructuredSheetTableListParams, opts ...option.RequestOption) (res *pagination.CursorIDPage[TableResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if structuredSheetID == "" {
		err = errors.New("missing required structured_sheet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/structured-sheets/%s/tables", url.PathEscape(structuredSheetID))
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

// List all tables extracted from the structured sheet. Only available when
// conversion status is 'completed'.
func (r *StructuredSheetTableService) ListAutoPaging(ctx context.Context, structuredSheetID string, query StructuredSheetTableListParams, opts ...option.RequestOption) *pagination.CursorIDPageAutoPager[TableResponse] {
	return pagination.NewCursorIDPageAutoPager(r.List(ctx, structuredSheetID, query, opts...))
}

// Download the table data in the specified format.
//
// Available formats:
//
// - `parquet`: Apache Parquet columnar format (recommended for data analysis)
// - `csv`: Comma-separated values (compatible with any spreadsheet application)
func (r *StructuredSheetTableService) Download(ctx context.Context, tableID string, params StructuredSheetTableDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.apache.parquet")}, opts...)
	if params.StructuredSheetID == "" {
		err = errors.New("missing required structured_sheet_id parameter")
		return nil, err
	}
	if tableID == "" {
		err = errors.New("missing required table_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/structured-sheets/%s/tables/%s/download", url.PathEscape(params.StructuredSheetID), url.PathEscape(tableID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Response representing a table extracted from a structured sheet.
//
// This is returned from GET (retrieve) and list table endpoints. Table names use a
// composite format: {normalized_sheet_name}\_\_{table_name}.
type TableResponse struct {
	// The unique identifier for this table.
	ID string `json:"id" api:"required"`
	// The timestamp when this table was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Composite table name: {normalized_sheet_name}**{table_name}. Uses lowercase
	// snake_case. Aggregation tables end with '**aggregations'. Two special metadata
	// tables exist per structured sheet: '**deeptable_workbook_metadata' (workbook
	// provenance info) and '**deeptable_table_overview' (summary of all tables).
	// Example: 'staffing**head_count' or 'staffing**head_count\_\_aggregations'.
	Name string `json:"name" api:"required"`
	// The object type, which is always 'table'.
	Object constant.Table `json:"object" default:"table"`
	// The original Excel sheet name this table came from.
	SheetName string `json:"sheet_name" api:"required"`
	// The ID of the structured sheet this table belongs to.
	StructuredSheetID string `json:"structured_sheet_id" api:"required"`
	// The type of table (relational, aggregation, tableless, or metadata).
	//
	// Any of "relational", "aggregation", "tableless", "metadata".
	Type TableResponseType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Name              respjson.Field
		Object            respjson.Field
		SheetName         respjson.Field
		StructuredSheetID respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TableResponse) RawJSON() string { return r.JSON.raw }
func (r *TableResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of table (relational, aggregation, tableless, or metadata).
type TableResponseType string

const (
	TableResponseTypeRelational  TableResponseType = "relational"
	TableResponseTypeAggregation TableResponseType = "aggregation"
	TableResponseTypeTableless   TableResponseType = "tableless"
	TableResponseTypeMetadata    TableResponseType = "metadata"
)

type StructuredSheetTableGetParams struct {
	// The unique identifier of the structured sheet conversion.
	StructuredSheetID string `path:"structured_sheet_id" api:"required" json:"-"`
	paramObj
}

type StructuredSheetTableListParams struct {
	// A cursor for pagination. Use the `last_id` from a previous response to fetch the
	// next page of results.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of tables to return per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StructuredSheetTableListParams]'s query parameters as
// `url.Values`.
func (r StructuredSheetTableListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StructuredSheetTableDownloadParams struct {
	// The unique identifier of the structured sheet conversion.
	StructuredSheetID string `path:"structured_sheet_id" api:"required" json:"-"`
	// The format to download the table data in.
	//
	// Any of "parquet", "csv".
	Format StructuredSheetTableDownloadParamsFormat `query:"format,omitzero" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [StructuredSheetTableDownloadParams]'s query parameters as
// `url.Values`.
func (r StructuredSheetTableDownloadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The format to download the table data in.
type StructuredSheetTableDownloadParamsFormat string

const (
	StructuredSheetTableDownloadParamsFormatParquet StructuredSheetTableDownloadParamsFormat = "parquet"
	StructuredSheetTableDownloadParamsFormatCsv     StructuredSheetTableDownloadParamsFormat = "csv"
)
