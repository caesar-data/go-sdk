// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"

	"github.com/caesar-data/go-sdk/internal/apijson"
	"github.com/caesar-data/go-sdk/internal/requestconfig"
	"github.com/caesar-data/go-sdk/packages/param"
	"github.com/caesar-data/go-sdk/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type PaginationPagination struct {
	HasNext bool  `json:"has_next"`
	Page    int64 `json:"page"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNext     respjson.Field
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginationPagination) RawJSON() string { return r.JSON.raw }
func (r *PaginationPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Pagination[T any] struct {
	Data       []T                  `json:"data"`
	Pagination PaginationPagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r Pagination[T]) RawJSON() string { return r.JSON.raw }
func (r *Pagination[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *Pagination[T]) GetNextPage() (res *Pagination[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}

	if r.Pagination.JSON.HasNext.Valid() && r.Pagination.HasNext == false {
		return nil, nil
	}
	currentPage := r.Pagination.Page
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *Pagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &Pagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PaginationAutoPager[T any] struct {
	page *Pagination[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPaginationAutoPager[T any](page *Pagination[T], err error) *PaginationAutoPager[T] {
	return &PaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *PaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *PaginationAutoPager[T]) Index() int {
	return r.run
}
