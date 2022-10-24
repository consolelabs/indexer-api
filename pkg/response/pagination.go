package response

import "github.com/consolelabs/indexer-api/pkg/model"

type PaginationResponse struct {
	model.Pagination
	Total int64 `json:"total"`
}
