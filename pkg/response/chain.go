package response

import "github.com/consolelabs/indexer-api/pkg/model"

type GetChainsResponse struct {
	*PaginationResponse `json:",omitempty"`
	Data                []model.Chain `json:"data"`
}
