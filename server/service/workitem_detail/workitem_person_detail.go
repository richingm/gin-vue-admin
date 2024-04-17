package workitem_detail

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/workitem_detail"
	workitem_detailReq "github.com/flipped-aurora/gin-vue-admin/server/model/workitem_detail/request"
)

type WorkitemPersonDetailService struct {
}

// CreateWorkitemPersonDetail 创建workitemPersonDetail表记录
// Author [piexlmax](https://github.com/piexlmax)
func (workitemPersonDetailService *WorkitemPersonDetailService) CreateWorkitemPersonDetail(workitemPersonDetail *workitem_detail.WorkitemPersonDetail) (err error) {
	err = global.GVA_DB.Create(workitemPersonDetail).Error
	return err
}

// DeleteWorkitemPersonDetail 删除workitemPersonDetail表记录
// Author [piexlmax](https://github.com/piexlmax)
func (workitemPersonDetailService *WorkitemPersonDetailService) DeleteWorkitemPersonDetail(ID string) (err error) {
	err = global.GVA_DB.Delete(&workitem_detail.WorkitemPersonDetail{}, "id = ?", ID).Error
	return err
}

// DeleteWorkitemPersonDetailByIds 批量删除workitemPersonDetail表记录
// Author [piexlmax](https://github.com/piexlmax)
func (workitemPersonDetailService *WorkitemPersonDetailService) DeleteWorkitemPersonDetailByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]workitem_detail.WorkitemPersonDetail{}, "id in ?", IDs).Error
	return err
}

// UpdateWorkitemPersonDetail 更新workitemPersonDetail表记录
// Author [piexlmax](https://github.com/piexlmax)
func (workitemPersonDetailService *WorkitemPersonDetailService) UpdateWorkitemPersonDetail(workitemPersonDetail workitem_detail.WorkitemPersonDetail) (err error) {
	err = global.GVA_DB.Save(&workitemPersonDetail).Error
	return err
}

// GetWorkitemPersonDetail 根据ID获取workitemPersonDetail表记录
// Author [piexlmax](https://github.com/piexlmax)
func (workitemPersonDetailService *WorkitemPersonDetailService) GetWorkitemPersonDetail(ID string) (workitemPersonDetail workitem_detail.WorkitemPersonDetail, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&workitemPersonDetail).Error
	return
}

// GetWorkitemPersonDetailInfoList 分页获取workitemPersonDetail表记录
// Author [piexlmax](https://github.com/piexlmax)
func (workitemPersonDetailService *WorkitemPersonDetailService) GetWorkitemPersonDetailInfoList(info workitem_detailReq.WorkitemPersonDetailSearch) (list []workitem_detail.WorkitemPersonDetail, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&workitem_detail.WorkitemPersonDetail{})
	var workitemPersonDetails []workitem_detail.WorkitemPersonDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.WorkitemId != nil {
		db = db.Where("workitem_id = ?", info.WorkitemId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&workitemPersonDetails).Error
	return workitemPersonDetails, total, err
}
