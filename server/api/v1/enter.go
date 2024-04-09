package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/article"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/knowledge"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	KnowledgeApiGroup knowledge.ApiGroup
	ArticleApiGroup   article.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
