package quickSearch

import (
	v1 "quickSearch/api/v1"

	"github.com/gin-gonic/gin"
)

type FileInfoRouter struct{}

func (a *FileInfoRouter) InitFileInfoRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	fileInfoRouter := Router.Group("file_info")
	fileInfoApi := v1.ApiGroupApp.ArticleApiGroup.FileInfoApi
	{
		//添加文件
		fileInfoRouter.POST("/add_fileinfo", fileInfoApi.AddFileInfo)
		//更新文件
		fileInfoRouter.POST("/update_fileinfo", fileInfoApi.UpdateFileInfo)
		//删除文件
		fileInfoRouter.GET("/delete_fileinfo", fileInfoApi.DeleteFileInfo)
	}
	return fileInfoRouter
}
