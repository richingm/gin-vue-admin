package article

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/article"
	articleReq "github.com/flipped-aurora/gin-vue-admin/server/model/article/request"
)

type ArticlesService struct {
}

// CreateArticles 创建articles表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articlesService *ArticlesService) CreateArticles(articles *article.Articles) (err error) {
	err = global.GVA_DB.Create(articles).Error
	return err
}

// DeleteArticles 删除articles表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articlesService *ArticlesService) DeleteArticles(ID string) (err error) {
	err = global.GVA_DB.Delete(&article.Articles{}, "id = ?", ID).Error
	return err
}

// DeleteArticlesByIds 批量删除articles表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articlesService *ArticlesService) DeleteArticlesByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]article.Articles{}, "id in ?", IDs).Error
	return err
}

// UpdateArticles 更新articles表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articlesService *ArticlesService) UpdateArticles(articles article.Articles) (err error) {
	err = global.GVA_DB.Save(&articles).Error
	return err
}

// GetArticles 根据ID获取articles表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articlesService *ArticlesService) GetArticles(ID string) (articles article.Articles, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&articles).Error
	return
}

// GetArticlesInfoList 分页获取articles表记录
// Author [piexlmax](https://github.com/piexlmax)
func (articlesService *ArticlesService) GetArticlesInfoList(info articleReq.ArticlesSearch) (list []article.Articles, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&article.Articles{})
	var articless []article.Articles
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.KnowledgeId != nil {
		db = db.Where("knowledge_id = ?", info.KnowledgeId)
	}
	if info.ImportanceLevel != "" {
		db = db.Where("importance_level = ?", info.ImportanceLevel)
	}
	if info.UnderstandLevel != "" {
		db = db.Where("understand_level = ?", info.UnderstandLevel)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["knowledge_id"] = true
	orderMap["pid"] = true
	orderMap["importance_level"] = true
	orderMap["understand_level"] = true
	orderMap["last_viewed_at"] = true
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

	err = db.Find(&articless).Error
	return articless, total, err
}
