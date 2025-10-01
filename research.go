// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package caesar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/caesar-data/go-sdk/internal/apijson"
	"github.com/caesar-data/go-sdk/internal/apiquery"
	"github.com/caesar-data/go-sdk/internal/requestconfig"
	"github.com/caesar-data/go-sdk/option"
	"github.com/caesar-data/go-sdk/packages/pagination"
	"github.com/caesar-data/go-sdk/packages/param"
	"github.com/caesar-data/go-sdk/packages/respjson"
)

// ResearchService contains methods and other services that help with interacting
// with the caesar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResearchService] method instead.
type ResearchService struct {
	Options []option.RequestOption
	Files   ResearchFileService
	Results ResearchResultService
}

// NewResearchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResearchService(opts ...option.RequestOption) (r ResearchService) {
	r = ResearchService{}
	r.Options = opts
	r.Files = NewResearchFileService(opts...)
	r.Results = NewResearchResultService(opts...)
	return
}

// Start a new research job using a query and optional file IDs.
func (r *ResearchService) New(ctx context.Context, body ResearchNewParams, opts ...option.RequestOption) (res *ResearchNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "research"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a single research object by ID.
func (r *ResearchService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ResearchGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("research/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a paginated list of research objects.
func (r *ResearchService) List(ctx context.Context, query ResearchListParams, opts ...option.RequestOption) (res *pagination.Pagination[ResearchListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "research"
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

// Returns a paginated list of research objects.
func (r *ResearchService) ListAutoPaging(ctx context.Context, query ResearchListParams, opts ...option.RequestOption) *pagination.PaginationAutoPager[ResearchListResponse] {
	return pagination.NewPaginationAutoPager(r.List(ctx, query, opts...))
}

type ResearchNewResponse struct {
	// Research job identifier.
	ID string `json:"id,required" format:"uuid"`
	// Current status of the research job.
	//
	// Any of "queued", "searching", "summarizing", "analyzing", "completed", "failed",
	// "researching".
	Status ResearchNewResponseStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ResearchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the research job.
type ResearchNewResponseStatus string

const (
	ResearchNewResponseStatusQueued      ResearchNewResponseStatus = "queued"
	ResearchNewResponseStatusSearching   ResearchNewResponseStatus = "searching"
	ResearchNewResponseStatusSummarizing ResearchNewResponseStatus = "summarizing"
	ResearchNewResponseStatusAnalyzing   ResearchNewResponseStatus = "analyzing"
	ResearchNewResponseStatusCompleted   ResearchNewResponseStatus = "completed"
	ResearchNewResponseStatusFailed      ResearchNewResponseStatus = "failed"
	ResearchNewResponseStatusResearching ResearchNewResponseStatus = "researching"
)

type ResearchGetResponse struct {
	// Research job identifier.
	ID string `json:"id,required" format:"uuid"`
	// ISO 8601 timestamp when the job was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Original query.
	Query string `json:"query,required"`
	// Ranked retrieval results and citations.
	Results []ResearchGetResponseResult `json:"results,required"`
	// Current status of the research job.
	//
	// Any of "queued", "searching", "summarizing", "analyzing", "completed", "failed",
	// "researching".
	Status ResearchGetResponseStatus `json:"status,required"`
	// Final content/synthesis (null until available).
	Content string `json:"content,nullable"`
	// Post-processed content (e.g., formatted/converted).
	TransformedContent string `json:"transformed_content,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Query              respjson.Field
		Results            respjson.Field
		Status             respjson.Field
		Content            respjson.Field
		TransformedContent respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ResearchGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchGetResponseResult struct {
	// Result object identifier.
	ID string `json:"id,required" format:"uuid"`
	// Relevance score (0–1).
	Score float64 `json:"score,required"`
	// Result title.
	Title string `json:"title,required"`
	// Canonical URL of the result.
	URL string `json:"url,required" format:"uri"`
	// Index used for inline citations (if present).
	CitationIndex int64 `json:"citation_index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Score         respjson.Field
		Title         respjson.Field
		URL           respjson.Field
		CitationIndex respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchGetResponseResult) RawJSON() string { return r.JSON.raw }
func (r *ResearchGetResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the research job.
type ResearchGetResponseStatus string

const (
	ResearchGetResponseStatusQueued      ResearchGetResponseStatus = "queued"
	ResearchGetResponseStatusSearching   ResearchGetResponseStatus = "searching"
	ResearchGetResponseStatusSummarizing ResearchGetResponseStatus = "summarizing"
	ResearchGetResponseStatusAnalyzing   ResearchGetResponseStatus = "analyzing"
	ResearchGetResponseStatusCompleted   ResearchGetResponseStatus = "completed"
	ResearchGetResponseStatusFailed      ResearchGetResponseStatus = "failed"
	ResearchGetResponseStatusResearching ResearchGetResponseStatus = "researching"
)

type ResearchListResponse struct {
	// Research job identifier.
	ID string `json:"id,required" format:"uuid"`
	// ISO 8601 timestamp when the job was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Original query.
	Query string `json:"query,required"`
	// Ranked retrieval results and citations.
	Results []ResearchListResponseResult `json:"results,required"`
	// Current status of the research job.
	//
	// Any of "queued", "searching", "summarizing", "analyzing", "completed", "failed",
	// "researching".
	Status ResearchListResponseStatus `json:"status,required"`
	// Final content/synthesis (null until available).
	Content string `json:"content,nullable"`
	// Post-processed content (e.g., formatted/converted).
	TransformedContent string `json:"transformed_content,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Query              respjson.Field
		Results            respjson.Field
		Status             respjson.Field
		Content            respjson.Field
		TransformedContent respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchListResponse) RawJSON() string { return r.JSON.raw }
func (r *ResearchListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchListResponseResult struct {
	// Result object identifier.
	ID string `json:"id,required" format:"uuid"`
	// Relevance score (0–1).
	Score float64 `json:"score,required"`
	// Result title.
	Title string `json:"title,required"`
	// Canonical URL of the result.
	URL string `json:"url,required" format:"uri"`
	// Index used for inline citations (if present).
	CitationIndex int64 `json:"citation_index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Score         respjson.Field
		Title         respjson.Field
		URL           respjson.Field
		CitationIndex respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *ResearchListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the research job.
type ResearchListResponseStatus string

const (
	ResearchListResponseStatusQueued      ResearchListResponseStatus = "queued"
	ResearchListResponseStatusSearching   ResearchListResponseStatus = "searching"
	ResearchListResponseStatusSummarizing ResearchListResponseStatus = "summarizing"
	ResearchListResponseStatusAnalyzing   ResearchListResponseStatus = "analyzing"
	ResearchListResponseStatusCompleted   ResearchListResponseStatus = "completed"
	ResearchListResponseStatusFailed      ResearchListResponseStatus = "failed"
	ResearchListResponseStatusResearching ResearchListResponseStatus = "researching"
)

type ResearchNewParams struct {
	// Primary research question or instruction.
	Query string `json:"query,required"`
	// Optional compute budget for the job. Defaults to 1.
	ComputeUnits param.Opt[int64] `json:"compute_units,omitzero"`
	// Optional system prompt to steer the assistant.
	SystemPrompt param.Opt[string] `json:"system_prompt,omitzero"`
	// IDs of previously uploaded files to include.
	Files []string `json:"files,omitzero" format:"uuid"`
	paramObj
}

func (r ResearchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ResearchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResearchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchListParams struct {
	// Page size (items per page).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// 1-based page index.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ResearchListParams]'s query parameters as `url.Values`.
func (r ResearchListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
