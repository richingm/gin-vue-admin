// 自动生成模板WorkitemPersonDetail
package workitem_detail

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// workitemPersonDetail表 结构体  WorkitemPersonDetail
type WorkitemPersonDetail struct {
	global.GVA_MODEL
	BorrowEndTime       *time.Time `json:"borrowEndTime" form:"borrowEndTime" gorm:"column:borrow_end_time;comment:借用结束时间;"`                                       //借用结束时间
	BorrowStartTime     *time.Time `json:"borrowStartTime" form:"borrowStartTime" gorm:"column:borrow_start_time;comment:借用开始时间;"`                                 //借用开始时间
	Email               string     `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:255;"`                                                            //邮箱
	EstimatedCost       *int       `json:"estimatedCost" form:"estimatedCost" gorm:"column:estimated_cost;comment:预估人力成本;size:10;"`                                //预估人力成本
	EstimatedDays       *int       `json:"estimatedDays" form:"estimatedDays" gorm:"column:estimated_days;comment:预估人天;size:10;"`                                  //预估人天
	IsEstimatedFinished *int       `json:"isEstimatedFinished" form:"isEstimatedFinished" gorm:"column:is_estimated_finished;comment:是否结束借用[1:结束,0:不结束];size:10;"` //是否结束借用[1:结束,0:不结束]
	Position            string     `json:"position" form:"position" gorm:"column:position;comment:岗位;size:100;"`                                                   //岗位
	PositionLevel       string     `json:"positionLevel" form:"positionLevel" gorm:"column:position_level;comment:能力等级;size:100;"`                                 //能力等级
	RoleId              *int       `json:"roleId" form:"roleId" gorm:"column:role_id;comment:项目角色ID;size:10;"`                                                     //项目角色ID
	RoleName            string     `json:"roleName" form:"roleName" gorm:"column:role_name;comment:项目角色名称;size:255;"`                                              //项目角色名称
	SettleDays          *int       `json:"settleDays" form:"settleDays" gorm:"column:settle_days;comment:结算人天;size:10;"`                                           //结算人天
	SettleLockDays      *int       `json:"settleLockDays" form:"settleLockDays" gorm:"column:settle_lock_days;comment:正在结算的天数;size:10;"`                           //正在结算的天数
	SettledCost         *int       `json:"settledCost" form:"settledCost" gorm:"column:settled_cost;comment:结算人力成本;size:10;"`                                      //结算人力成本
	SettledDays         *int       `json:"settledDays" form:"settledDays" gorm:"column:settled_days;comment:已经结算的天数;size:10;"`                                     //已经结算的天数
	UserId              *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:19;"`                                                       //用户id
	UserName            string     `json:"userName" form:"userName" gorm:"column:user_name;comment:用户名;size:255;"`                                                 //用户名
	WorkitemId          *int       `json:"workitemId" form:"workitemId" gorm:"column:workitem_id;comment:工作项id;size:19;" binding:"required"`                       //工作项id
}

// TableName workitemPersonDetail表 WorkitemPersonDetail自定义表名 workitem_person_detail
func (WorkitemPersonDetail) TableName() string {
	return "workitem_person_detail"
}
