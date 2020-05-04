package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin-demo/models"
	"go-admin-demo/tools"
	"go-admin-demo/tools/app"
	"go-admin-demo/tools/app/msg"
)

func GetArticleList(c *gin.Context) {
	var data models.Article
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.Title = c.Request.FormValue("title")
	data.Author = c.Request.FormValue("author")


	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetArticle(c *gin.Context) {
	var data models.Article
	data.ArticleId, _ = tools.StringToInt(c.Param("articleId"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func InsertArticle(c *gin.Context) {
	var data models.Article
	err := c.ShouldBindJSON(&data)
	//data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateArticle(c *gin.Context) {
	var data models.Article
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	//data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.Update(data.ArticleId)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func DeleteArticle(c *gin.Context) {
	var data models.Article
	//data.UpdateBy = tools.GetUserIdStr(c)

	IDS := tools.IdsStrToIdsIntGroup("articleId", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}