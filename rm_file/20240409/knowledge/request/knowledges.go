package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type KnowledgesSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	ImportLevel    string     `json:"importLevel" form:"importLevel" `
	Name           string     `json:"name" form:"name" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
