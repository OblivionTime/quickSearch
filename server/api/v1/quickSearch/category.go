package quickSearch

import (
	"fmt"
	"quickSearch/global"
	QSResp "quickSearch/model/common"
	"quickSearch/model/quickSearch"

	"github.com/gin-gonic/gin"
)

type CategoryApi struct{}

//添加栏目
func (b *CategoryApi) AddCategory(ctx *gin.Context) {
	var info quickSearch.Category
	if err := ctx.ShouldBindJSON(&info); err != nil {
		QSResp.FailWithMessage("参数错误!!!", ctx)
		return
	}
	// 检查数据库中是否存在相同名称的记录
	var existingCategory quickSearch.Category
	if err := global.DB.Where("name = ?", info.Name).First(&existingCategory).Error; err == nil {
		// 如果存在同名记录，为名称添加编号后缀
		info.Name = generateUniqueName(info.Name)
	}
	if err := global.DB.Create(&info).Error; err != nil {
		QSResp.Fail(ctx)
		return
	}
	QSResp.Ok(ctx)
}
func generateUniqueName(baseName string) string {
	// 递增的编号
	suffix := 1

	// 循环直到找到一个不存在的名称
	for {
		newName := fmt.Sprintf("%s_%d", baseName, suffix)
		// 检查数据库中是否存在相同名称的记录
		var existingCategory quickSearch.Category
		if err := global.DB.Where("name = ?", newName).First(&existingCategory).Error; err != nil {
			// 不存在相同名称的记录，返回新名称
			return newName
		}

		// 如果存在相同名称的记录，递增编号
		suffix++
	}
}

//更新栏目信息
func (b *CategoryApi) UpdateCategory(ctx *gin.Context) {
	var info quickSearch.Category
	if err := ctx.ShouldBindJSON(&info); err != nil {
		QSResp.FailWithMessage("参数错误!!!", ctx)
		return
	}
	// 检查数据库中是否存在相同名称的记录
	var existingCategory quickSearch.Category
	if err := global.DB.Where("name = ?", info.Name).First(&existingCategory).Error; err == nil {
		// 如果存在同名记录，为名称添加编号后缀
		info.Name = generateUniqueName(info.Name)
	}
	if err := global.DB.Model(quickSearch.Category{}).Where("id=?", info.ID).UpdateColumn("name", info.Name).Error; err != nil {
		QSResp.Fail(ctx)
		return
	}
	QSResp.Ok(ctx)
}

//删除栏目信息
func (b *CategoryApi) DeleteCategory(ctx *gin.Context) {
	var id = ctx.Query("id")
	var count int
	//判断是否为最后一个
	global.DB.Model(quickSearch.Category{}).Count(&count)
	if count == 1 {
		QSResp.FailWithMessage("删除失败,无法删除最后一条", ctx)
		return
	}
	//查询栏目下所有文件信息并删除
	var list []quickSearch.FileInfo
	global.DB.Model(quickSearch.FileInfo{}).Where("column_id=?", id).Scan(&list)
	for _, f := range list {
		global.DB.Model(quickSearch.Category{}).Delete(&quickSearch.FileInfo{}, f.ID)
	}
	if err := global.DB.Model(quickSearch.Category{}).Delete(&quickSearch.Category{}, id).Error; err != nil {
		QSResp.Fail(ctx)
		return
	}
	QSResp.Ok(ctx)
}
