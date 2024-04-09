package knowledge

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/knowledge"
	knowledgeReq "github.com/flipped-aurora/gin-vue-admin/server/model/knowledge/request"
	"gorm.io/gorm"
)

type KnowledgesService struct {
}

// CreateKnowledges 创建knowledges表记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgesService *KnowledgesService) CreateKnowledges(knowledges *knowledge.Knowledges) (err error) {
	err = global.GVA_DB.Create(knowledges).Error
	return err
}

// DeleteKnowledges 删除knowledges表记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgesService *KnowledgesService) DeleteKnowledges(id string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&knowledge.Knowledges{}).Where("id = ?", id).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&knowledge.Knowledges{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteKnowledgesByIds 批量删除knowledges表记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgesService *KnowledgesService) DeleteKnowledgesByIds(ids []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&knowledge.Knowledges{}).Where("id in ?", ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids).Delete(&knowledge.Knowledges{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateKnowledges 更新knowledges表记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgesService *KnowledgesService) UpdateKnowledges(knowledges knowledge.Knowledges) (err error) {
	err = global.GVA_DB.Save(&knowledges).Error
	return err
}

// GetKnowledges 根据id获取knowledges表记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgesService *KnowledgesService) GetKnowledges(id string) (knowledges knowledge.Knowledges, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&knowledges).Error
	return
}

// GetKnowledgesInfoList 分页获取knowledges表记录
// Author [piexlmax](https://github.com/piexlmax)
func (knowledgesService *KnowledgesService) GetKnowledgesInfoList(info knowledgeReq.KnowledgesSearch) (list []knowledge.Knowledges, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&knowledge.Knowledges{})
	var knowledgess []knowledge.Knowledges
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ? ", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ImportLevel != "" {
		db = db.Where("import_level = ?", info.ImportLevel)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["updated_at"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&knowledgess).Error
	return knowledgess, total, err
}
