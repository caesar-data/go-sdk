// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package caesar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/caesar-go/internal/apijson"
	"github.com/stainless-sdks/caesar-go/internal/requestconfig"
	"github.com/stainless-sdks/caesar-go/option"
	"github.com/stainless-sdks/caesar-go/packages/respjson"
)

// ResearchResultService contains methods and other services that help with
// interacting with the caesar API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResearchResultService] method instead.
type ResearchResultService struct {
	Options []option.RequestOption
}

// NewResearchResultService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResearchResultService(opts ...option.RequestOption) (r ResearchResultService) {
	r = ResearchResultService{}
	r.Options = opts
	return
}

// Returns the raw content for a specific result within a research object.
func (r *ResearchResultService) GetContent(ctx context.Context, resultID string, query ResearchResultGetContentParams, opts ...option.RequestOption) (res *ResearchResultGetContentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if resultID == "" {
		err = errors.New("missing required resultId parameter")
		return
	}
	path := fmt.Sprintf("research/%s/results/%s/content", query.ID, resultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ResearchResultGetContentResponse struct {
	// Raw extracted content for this result (may include HTML, markdown, or plain
	// text).
	Content string `json:"content,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchResultGetContentResponse) RawJSON() string { return r.JSON.raw }
func (r *ResearchResultGetContentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchResultGetContentParams struct {
	ID string `path:"id,required" format:"uuid" json:"-"`
	paramObj
}
