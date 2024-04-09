package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ArticlesSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	KnowledgeId     *int   `json:"knowledgeId" form:"knowledgeId" `
	ImportanceLevel string `json:"importanceLevel" form:"importanceLevel" `
	UnderstandLevel string `json:"understandLevel" form:"understandLevel" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
