package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/article"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/knowledge"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/workitem_detail"
)

type RouterGroup struct {
	System          system.RouterGroup
	Example         example.RouterGroup
	Knowledge       knowledge.RouterGroup
	Article         article.RouterGroup
	Workitem_detail workitem_detail.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
