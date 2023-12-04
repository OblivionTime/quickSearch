package quickSearch

import (
	v1 "quickSearch/api/v1"

	"github.com/gin-gonic/gin"
)

type CategoryRouter struct{}

func (a *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	categoryRouter := Router.Group("category")
	categoryApi := v1.ApiGroupApp.ArticleApiGroup.CategoryApi
	{
		//添加栏目
		categoryRouter.POST("/add_category", categoryApi.AddCategory)
		//更新栏目
		categoryRouter.POST("/update_category", categoryApi.UpdateCategory)
		//删除栏目
		categoryRouter.GET("/delete_category", categoryApi.DeleteCategory)
	}
	return categoryRouter
}
