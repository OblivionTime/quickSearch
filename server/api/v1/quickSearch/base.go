package quickSearch

import (
	"fmt"
	"quickSearch/global"
	QSResp "quickSearch/model/common"
	"quickSearch/model/quickSearch"

	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

//获取文件列表
func (b *BaseApi) GetFileList(ctx *gin.Context) {
	/*
	* 读取fileInfo获取所有文件路径及其分类情况
	 */
	id := ctx.DefaultQuery("id", "0")
	name := ctx.DefaultQuery("name", "")
	var sql string
	if id == "0" && name == "" {
		sql = "select f.*,c.name as category_name from category as c,file_info as f where f.column_id = c.id"
	} else if id == "0" {
		sql = fmt.Sprintf("select f.*,c.name as category_name from category as c,file_info as f where f.column_id = c.id and f.name like '%%%s%%' or  f.file_path like '%%%s%%'", name, name)
	} else if name == "" {
		sql = fmt.Sprintf("select f.*,c.name as category_name from category as c,file_info as f where f.column_id = c.id and c.id=%s", id)
	} else {
		sql = fmt.Sprintf("select f.*,c.name as category_name from category as c,file_info as f where f.column_id = c.id and f.name like '%%%s%%' or  f.file_path like '%%%s%%' and c.id=%s", name, name, id)

	}
	var List []quickSearch.Info = make([]quickSearch.Info, 0)
	global.DB.Raw(sql).Order("id desc").Scan(&List)
	var sr quickSearch.SearchResult
	sr.FileList = List
	var cList []quickSearch.Category = make([]quickSearch.Category, 0)
	global.DB.Table("category").Find(&cList)
	sr.CategoryList = cList
	QSResp.OkWithData(sr, ctx)
}
