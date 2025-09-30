// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package caesar

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/caesar-go/internal/apiform"
	"github.com/stainless-sdks/caesar-go/internal/apijson"
	"github.com/stainless-sdks/caesar-go/internal/apiquery"
	"github.com/stainless-sdks/caesar-go/internal/requestconfig"
	"github.com/stainless-sdks/caesar-go/option"
	"github.com/stainless-sdks/caesar-go/packages/pagination"
	"github.com/stainless-sdks/caesar-go/packages/param"
	"github.com/stainless-sdks/caesar-go/packages/respjson"
)

// ResearchFileService contains methods and other services that help with
// interacting with the caesar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResearchFileService] method instead.
type ResearchFileService struct {
	Options []option.RequestOption
}

// NewResearchFileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResearchFileService(opts ...option.RequestOption) (r ResearchFileService) {
	r = ResearchFileService{}
	r.Options = opts
	return
}

// Upload a file via multipart form and create a Research File object.
func (r *ResearchFileService) New(ctx context.Context, body ResearchFileNewParams, opts ...option.RequestOption) (res *ResearchFileNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "research/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a paginated list of Research File objects.
func (r *ResearchFileService) List(ctx context.Context, query ResearchFileListParams, opts ...option.RequestOption) (res *pagination.Pagination[ResearchFileListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "research/files"
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

// Returns a paginated list of Research File objects.
func (r *ResearchFileService) ListAutoPaging(ctx context.Context, query ResearchFileListParams, opts ...option.RequestOption) *pagination.PaginationAutoPager[ResearchFileListResponse] {
	return pagination.NewPaginationAutoPager(r.List(ctx, query, opts...))
}

type ResearchFileNewResponse struct {
	// Unique identifier for the file.
	ID string `json:"id,required" format:"uuid"`
	// MIME type of the file as detected/stored.
	ContentType string `json:"content_type,required"`
	// Original uploaded filename.
	FileName string `json:"file_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ContentType respjson.Field
		FileName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchFileNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ResearchFileNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchFileListResponse struct {
	// Unique identifier for the file.
	ID string `json:"id,required" format:"uuid"`
	// MIME type of the file as detected/stored.
	ContentType string `json:"content_type,required"`
	// Original uploaded filename.
	FileName string `json:"file_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ContentType respjson.Field
		FileName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchFileListResponse) RawJSON() string { return r.JSON.raw }
func (r *ResearchFileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchFileNewParams struct {
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	paramObj
}

func (r ResearchFileNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type ResearchFileListParams struct {
	// Page size (items per page).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// 1-based page index.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ResearchFileListParams]'s query parameters as `url.Values`.
func (r ResearchFileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
