// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package caesar

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/caesar-data/go-sdk/internal/apijson"
	"github.com/caesar-data/go-sdk/internal/requestconfig"
	"github.com/caesar-data/go-sdk/option"
	"github.com/caesar-data/go-sdk/packages/param"
	"github.com/caesar-data/go-sdk/packages/respjson"
)

// X402ResearchService contains methods and other services that help with
// interacting with the caesar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewX402ResearchService] method instead.
type X402ResearchService struct {
	Options []option.RequestOption
}

// NewX402ResearchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewX402ResearchService(opts ...option.RequestOption) (r X402ResearchService) {
	r = X402ResearchService{}
	r.Options = opts
	return
}

// Start a new research job using x402 payment. This endpoint mints a temporary API
// key that is returned in the response and is billed via the x402 settlement flow
// instead of your Caesar API credits.
func (r *X402ResearchService) New(ctx context.Context, params X402ResearchNewParams, opts ...option.RequestOption) (res *X402ResearchNewResponse, err error) {
	if !param.IsOmitted(params.XPayment) {
		opts = append(opts, option.WithHeader("X-PAYMENT", fmt.Sprintf("%s", params.XPayment)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "x402/research"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type X402ResearchNewResponse struct {
	// Research job identifier.
	ID string `json:"id,required" format:"uuid"`
	// Current status of the research job.
	//
	// Any of "queued", "searching", "summarizing", "analyzing", "completed", "failed",
	// "researching".
	Status X402ResearchNewResponseStatus `json:"status,required"`
	// Temporary API key secret created for this x402 request.
	APIKeySecret string `json:"api_key_secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Status       respjson.Field
		APIKeySecret respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r X402ResearchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *X402ResearchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the research job.
type X402ResearchNewResponseStatus string

const (
	X402ResearchNewResponseStatusQueued      X402ResearchNewResponseStatus = "queued"
	X402ResearchNewResponseStatusSearching   X402ResearchNewResponseStatus = "searching"
	X402ResearchNewResponseStatusSummarizing X402ResearchNewResponseStatus = "summarizing"
	X402ResearchNewResponseStatusAnalyzing   X402ResearchNewResponseStatus = "analyzing"
	X402ResearchNewResponseStatusCompleted   X402ResearchNewResponseStatus = "completed"
	X402ResearchNewResponseStatusFailed      X402ResearchNewResponseStatus = "failed"
	X402ResearchNewResponseStatusResearching X402ResearchNewResponseStatus = "researching"
)

type X402ResearchNewParams struct {
	// Primary research question or instruction.
	Query    string `json:"query,required"`
	XPayment string `header:"X-PAYMENT,required" json:"-"`
	// Optional compute budget for the job. Defaults to 1.
	ComputeUnits param.Opt[int64] `json:"compute_units,omitzero"`
	// Optional system prompt to steer the assistant.
	SystemPrompt param.Opt[string] `json:"system_prompt,omitzero"`
	paramObj
}

func (r X402ResearchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow X402ResearchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *X402ResearchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
