// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package deeptable

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/deeptable-com/deeptable-go/internal/apiform"
	"github.com/deeptable-com/deeptable-go/internal/apijson"
	"github.com/deeptable-com/deeptable-go/internal/apiquery"
	"github.com/deeptable-com/deeptable-go/internal/requestconfig"
	"github.com/deeptable-com/deeptable-go/option"
	"github.com/deeptable-com/deeptable-go/packages/pagination"
	"github.com/deeptable-com/deeptable-go/packages/param"
	"github.com/deeptable-com/deeptable-go/packages/respjson"
	"github.com/deeptable-com/deeptable-go/shared/constant"
)

// Upload and manage spreadsheet files. Files must be Excel (.xlsx) format.
//
// FileService contains methods and other services that help with interacting with
// the deeptable API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.options = opts
	return
}

// Get metadata for a specific file.
func (r *FileService) Get(ctx context.Context, fileID string, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.options, opts)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/files/%s", url.PathEscape(fileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all files uploaded by the current user.
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *pagination.CursorIDPage[File], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/files"
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

// List all files uploaded by the current user.
func (r *FileService) ListAutoPaging(ctx context.Context, query FileListParams, opts ...option.RequestOption) *pagination.CursorIDPageAutoPager[File] {
	return pagination.NewCursorIDPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a file. This cannot be undone.
func (r *FileService) Delete(ctx context.Context, fileID string, opts ...option.RequestOption) (res *FileDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/files/%s", url.PathEscape(fileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Download the original uploaded file content.
func (r *FileService) Download(ctx context.Context, fileID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")}, opts...)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/files/%s/content", url.PathEscape(fileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Upload an Excel spreadsheet file for later processing.
//
// Supported formats:
//
// - Excel (.xlsx)
//
// Maximum file size: 100 MB
func (r *FileService) Upload(ctx context.Context, body FileUploadParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Response representing an uploaded file.
//
// This is returned from POST (upload), GET (retrieve), and list endpoints.
type File struct {
	// The unique identifier for this file.
	ID string `json:"id" api:"required"`
	// The MIME type of the file.
	ContentType string `json:"content_type" api:"required"`
	// The timestamp when the file was uploaded.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The original filename of the uploaded file.
	FileName string `json:"file_name" api:"required"`
	// The object type, which is always 'file'.
	Object constant.File `json:"object" api:"required"`
	// The size of the file in bytes.
	Size int64 `json:"size" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ContentType respjson.Field
		CreatedAt   respjson.Field
		FileName    respjson.Field
		Object      respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r File) RawJSON() string { return r.JSON.raw }
func (r *File) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from deleting a file.
//
// Following the OpenAI API convention for delete responses.
type FileDeleteResponse struct {
	// The unique identifier of the deleted file.
	ID string `json:"id" api:"required"`
	// Whether the file was successfully deleted.
	Deleted bool `json:"deleted" api:"required"`
	// The object type, which is always 'file'.
	Object constant.File `json:"object" api:"required"`
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
func (r FileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListParams struct {
	// A cursor for pagination. Use the `last_id` from a previous response to fetch the
	// next page.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of files to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileUploadParams struct {
	// The spreadsheet file to upload
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	paramObj
}

func (r FileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
