package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/article"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/knowledge"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/workitem_detail"
)

type ApiGroup struct {
	SystemApiGroup          system.ApiGroup
	ExampleApiGroup         example.ApiGroup
	KnowledgeApiGroup       knowledge.ApiGroup
	ArticleApiGroup         article.ApiGroup
	Workitem_detailApiGroup workitem_detail.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
