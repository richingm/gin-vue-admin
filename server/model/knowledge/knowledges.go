// 自动生成模板Knowledges
package knowledge

import (
	"time"
)

// knowledges表 结构体  Knowledges
type Knowledges struct {
	CreatedAt   *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;"`                                  //创建时间
	DeletedAt   *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:删除时间;"`                                  //删除时间
	Id          *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:ID;size:10;"`                                       //ID
	ImportLevel string     `json:"importLevel" form:"importLevel" gorm:"column:import_level;comment:重要程度;size:50;" binding:"required"` //重要程度
	Name        string     `json:"name" form:"name" gorm:"column:name;comment:知识库名称;size:255;" binding:"required"`                     //知识库名称
	Notes       string     `json:"notes" form:"notes" gorm:"column:notes;comment:备注;size:255;"`                                        //备注
	Pid         *int       `json:"pid" form:"pid" gorm:"column:pid;comment:父知识库;size:10;"`                                             //父知识库
	PName       string     `json:"pName" gorm:"-"`                                                                                     //父知识库
	UpdatedAt   *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:更新时间;"`                                  //更新时间
	CreatedBy   uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName knowledges表 Knowledges自定义表名 knowledges
func (Knowledges) TableName() string {
	return "knowledges"
}

type KnowledgesOption struct {
	Id    *int               `json:"id" form:"id" gorm:"primarykey;column:id;comment:ID;size:10;"`                   //ID
	Name  string             `json:"name" form:"name" gorm:"column:name;comment:知识库名称;size:255;" binding:"required"` //知识库名称
	Pid   *int               `json:"pid" form:"pid" gorm:"column:pid;comment:父知识库;size:10;"`                         //父知识库
	Child []KnowledgesOption `json:"child" gorm:"-"`
}
