// 自动生成模板Articles
package article

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// articles表 结构体  Articles
type Articles struct {
	global.GVA_MODEL
	KnowledgeId     *int       `json:"knowledgeId" form:"knowledgeId" gorm:"column:knowledge_id;comment:知识库id;size:10;" binding:"required"`            //知识库
	Pid             *int       `json:"pid" form:"pid" gorm:"column:pid;comment:父id;size:10;"`                                                          //父id
	Title           string     `json:"title" form:"title" gorm:"column:title;comment:标题;size:255;" binding:"required"`                                 //标题
	ImportanceLevel string     `json:"importanceLevel" form:"importanceLevel" gorm:"column:importance_level;comment:重要程度;size:50;" binding:"required"` //重要程度
	UnderstandLevel string     `json:"understandLevel" form:"understandLevel" gorm:"column:understand_level;comment:理解程度;size:50;" binding:"required"` //理解程度
	LastViewedAt    *time.Time `json:"lastViewedAt" form:"lastViewedAt" gorm:"column:last_viewed_at;comment:最后查看时间;"`                                  //最后查看时间
	Content         string     `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;" binding:"required"`                          //内容
}

// TableName articles表 Articles自定义表名 articles
func (Articles) TableName() string {
	return "articles"
}
