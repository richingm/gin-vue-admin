package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/article"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/knowledge"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/workitem_detail"
)

type ServiceGroup struct {
	SystemServiceGroup          system.ServiceGroup
	ExampleServiceGroup         example.ServiceGroup
	KnowledgeServiceGroup       knowledge.ServiceGroup
	ArticleServiceGroup         article.ServiceGroup
	Workitem_detailServiceGroup workitem_detail.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
