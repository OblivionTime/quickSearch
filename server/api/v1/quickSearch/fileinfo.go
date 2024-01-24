package quickSearch

import (
	"quickSearch/global"
	QSResp "quickSearch/model/common"
	"quickSearch/model/quickSearch"

	"github.com/gin-gonic/gin"
)

type FileInfoApi struct{}

// 添加文件
func (b *FileInfoApi) AddFileInfo(ctx *gin.Context) {
	var info quickSearch.FileInfo
	if err := ctx.ShouldBindJSON(&info); err != nil {
		QSResp.FailWithMessage("参数错误!!!", ctx)
		return
	}
	var existinginfo quickSearch.FileInfo
	if err := global.DB.Where("file_path = ?", info.FilePath).First(&existinginfo).Error; err == nil {
		QSResp.FailWithMessage("文件/文件夹已存在,请勿重复添加", ctx)
		return
	}
	if err := global.DB.Create(&info).Error; err != nil {
		QSResp.Fail(ctx)
		return
	}
	QSResp.Ok(ctx)
}

// 更新文件信息
func (b *FileInfoApi) UpdateFileInfo(ctx *gin.Context) {
	var info quickSearch.FileInfo
	if err := ctx.ShouldBindJSON(&info); err != nil {
		QSResp.FailWithMessage("参数错误!!!", ctx)
		return
	}
	if err := global.DB.Model(quickSearch.FileInfo{}).Where("id=?", info.ID).UpdateColumn("name", info.Name).Error; err != nil {
		QSResp.Fail(ctx)
		return
	}
	QSResp.Ok(ctx)
}

// 删除文件信息
func (b *FileInfoApi) DeleteFileInfo(ctx *gin.Context) {
	var id = ctx.Query("id")
	if err := global.DB.Model(quickSearch.FileInfo{}).Delete(&quickSearch.FileInfo{}, id).Error; err != nil {
		QSResp.Fail(ctx)
		return
	}
	QSResp.Ok(ctx)
}

// //文件预览
// func (b *FileInfoApi) PreviewFIleInfo(ctx *gin.Context) {
// 	var info quickSearch.FileInfo
// 	if err := ctx.ShouldBindJSON(&info); err != nil {
// 		QSResp.FailWithMessage("参数错误!!!", ctx)
// 		return
// 	}
// 	if err := global.DB.Create(&info).Error; err != nil {
// 		QSResp.Fail(ctx)
// 		return
// 	}
// 	QSResp.Ok(ctx)
// }

//
