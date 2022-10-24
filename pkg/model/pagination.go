package model

import (
	"github.com/consolelabs/indexer-api/pkg/utils"
)

var maxPageSize int64 = 50

type Pagination struct {
	Page         int64  `json:"page" form:"page,default=0"`            // page index
	Size         int64  `json:"size" form:"size"`                      // page size
	Sort         string `json:"sort" form:"sort" swaggerignore:"true"` // sort by field
	Standardized bool   `json:"-" form:"-" swaggerignore:"true"`
}

func (p *Pagination) Standardize() {
	if p.Standardized {
		return
	}

	if p.Page < 0 {
		p.Page = 0
	}

	if p.Size <= 0 || p.Size >= maxPageSize {
		p.Size = maxPageSize
	}

	if p.Sort == "" {
		return
	}

	p.Sort = utils.StandardizeSortQuery(p.Sort)
	p.Standardized = true
}
