package knowledge

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/knowledge"
	knowledgeReq "github.com/flipped-aurora/gin-vue-admin/server/model/knowledge/request"
	"gorm.io/gorm"
)

type KnowledgesService struct {
}

func NewKnowledgesService() *KnowledgesService {
	return &KnowledgesService{}
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

func (knowledgesService *KnowledgesService) GetKnowledgesByIds(ids []int) (knowledges []knowledge.Knowledges, err error) {
	err = global.GVA_DB.Where("id in ?", ids).Find(&knowledges).Error
	return
}

// loadChildren recursively loads child options for a given knowledge option.
func (knowledgesService *KnowledgesService) loadChildren(option *knowledge.KnowledgesOption) error {
	children := []knowledge.KnowledgesOption{}
	err := global.GVA_DB.Where("pid = ?", option.Id).Find(&children).Error
	if err != nil {
		return err
	}

	for i := range children {
		err := knowledgesService.loadChildren(&children[i])
		if err != nil {
			return err
		}
	}

	option.Child = children
	return nil
}

// GetKnowledgesOptions retrieves all the KnowledgesOption records from the database.
func (knowledgesService *KnowledgesService) GetKnowledgesOptions() ([]knowledge.KnowledgesOption, error) {
	var topKnowledges []knowledge.KnowledgesOption
	err := global.GVA_DB.Where("pid = 0").Find(&topKnowledges).Error // Assuming top-level knowledge options have a nil `pid`
	if err != nil {
		return nil, err
	}

	for i := range topKnowledges {
		err := knowledgesService.loadChildren(&topKnowledges[i])
		if err != nil {
			return nil, err
		}
	}

	return topKnowledges, nil
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
