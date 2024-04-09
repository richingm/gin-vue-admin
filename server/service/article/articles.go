package article

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/article"
	articleReq "github.com/flipped-aurora/gin-vue-admin/server/model/article/request"
	knowledge2 "github.com/flipped-aurora/gin-vue-admin/server/model/knowledge"
	"github.com/flipped-aurora/gin-vue-admin/server/service/knowledge"
	"github.com/samber/lo"
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

func (articlesService *ArticlesService) GetArticlesByIds(ids []int) (articles []article.Articles, err error) {
	err = global.GVA_DB.Where("id in ?", ids).First(&articles).Error
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
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
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

	// 获取父id
	pids := lo.Map(articless, func(item article.Articles, index int) int {
		return *item.Pid
	})

	// 获取父id文章
	articlesList, err := articlesService.GetArticlesByIds(pids)
	if err != nil {
		return nil, 0, err
	}
	articlesMap := lo.KeyBy(articlesList, func(item article.Articles) int {
		return int(item.ID)
	})
	// 获取知识库id
	knowledgeIds := lo.Map(articless, func(item article.Articles, index int) int {
		return *item.KnowledgeId
	})

	// 获取知识库
	knowledgesService := knowledge.NewKnowledgesService()
	knowledges, err := knowledgesService.GetKnowledgesByIds(knowledgeIds)
	if err != nil {
		return nil, 0, err
	}
	knowledgeMap := lo.KeyBy(knowledges, func(item knowledge2.Knowledges) int {
		return *item.Id
	})

	articless = lo.Map(articless, func(item article.Articles, index int) article.Articles {
		if _, ok := articlesMap[*item.Pid]; ok {
			item.PTitle = articlesMap[*item.Pid].Title
		}
		if _, ok := knowledgeMap[*item.KnowledgeId]; ok {
			item.KnowledgeName = knowledgeMap[*item.KnowledgeId].Name
		}
		return item
	})

	return articless, total, err
}
