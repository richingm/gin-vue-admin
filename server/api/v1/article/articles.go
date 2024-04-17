package article

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/article"
	articleReq "github.com/flipped-aurora/gin-vue-admin/server/model/article/request"
	response2 "github.com/flipped-aurora/gin-vue-admin/server/model/article/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"strconv"
)

type ArticlesApi struct {
}

var articlesService = service.ServiceGroupApp.ArticleServiceGroup.ArticlesService
var knowledgeService = service.ServiceGroupApp.KnowledgeServiceGroup.KnowledgesService

// CreateArticles 创建articles表
// @Tags Articles
// @Summary 创建articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body article.Articles true "创建articles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /articles/createArticles [post]
func (articlesApi *ArticlesApi) CreateArticles(c *gin.Context) {
	var articles article.Articles
	err := c.ShouldBindJSON(&articles)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := articlesService.CreateArticles(&articles); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteArticles 删除articles表
// @Tags Articles
// @Summary 删除articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body article.Articles true "删除articles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /articles/deleteArticles [delete]
func (articlesApi *ArticlesApi) DeleteArticles(c *gin.Context) {
	ID := c.Query("id")
	if err := articlesService.DeleteArticles(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteArticlesByIds 批量删除articles表
// @Tags Articles
// @Summary 批量删除articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /articles/deleteArticlesByIds [delete]
func (articlesApi *ArticlesApi) DeleteArticlesByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := articlesService.DeleteArticlesByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateArticles 更新articles表
// @Tags Articles
// @Summary 更新articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body article.Articles true "更新articles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /articles/updateArticles [put]
func (articlesApi *ArticlesApi) UpdateArticles(c *gin.Context) {
	var articles article.Articles
	err := c.ShouldBindJSON(&articles)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := articlesService.UpdateArticles(articles); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindArticles 用id查询articles表
// @Tags Articles
// @Summary 用id查询articles表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query article.Articles true "用id查询articles表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /articles/findArticles [get]
func (articlesApi *ArticlesApi) FindArticles(c *gin.Context) {
	ID := c.Query("id")
	if rearticles, err := articlesService.GetArticles(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rearticles": rearticles}, c)
	}
}

// GetArticlesList 分页获取articles表列表
// @Tags Articles
// @Summary 分页获取articles表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query articleReq.ArticlesSearch true "分页获取articles表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /articles/getArticlesList [get]
func (articlesApi *ArticlesApi) GetArticlesList(c *gin.Context) {
	var pageInfo articleReq.ArticlesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := articlesService.GetArticlesInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetArticlesByKnowledgeId 根据知识库id获取文章
// @Tags Articles
// @Summary 根据知识库id获取文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query articleReq.ArticlesSearchByKnowledgeId true "根据知识库id获取参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /articles/getArticlesMind [get]
func (articlesApi *ArticlesApi) GetArticlesByKnowledgeId(c *gin.Context) {
	var param articleReq.ArticlesSearchByKnowledgeId
	err := c.ShouldBindQuery(&param)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	knowledges, err := knowledgeService.GetKnowledges(param.KnowledgeId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var pageInfo articleReq.ArticlesSearch
	pageInfo.Page = 1
	pageInfo.KnowledgeId = &param.KnowledgeId
	pageInfo.PageSize = 999999999
	list, _, err := articlesService.GetArticlesInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		mind := response2.Mind{
			Meta: response2.Meta{
				Name:    "demo",
				Author:  "hizzgdev@163.com",
				Version: "0.2",
			},
			Format: "node_array",
		}
		root := response2.MindNode{ID: "root", IsRoot: true, Topic: knowledges.Name, Direction: "right"}
		mind.Data = append(mind.Data, root)
		data := lo.Map(list, func(item article.Articles, index int) response2.MindNode {
			pid := strconv.Itoa(item.Pid)
			if pid == "0" {
				pid = "root"
			}
			return response2.MindNode{
				ID:        strconv.Itoa(int(item.ID)),
				Topic:     item.Title,
				ParentId:  pid,
				Direction: "right",
			}
		})
		mind.Data = append(mind.Data, data...)
		response.OkWithData(gin.H{"rearticles": mind}, c)
	}
}
