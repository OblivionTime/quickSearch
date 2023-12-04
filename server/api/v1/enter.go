package v1

import "quickSearch/api/v1/quickSearch"

type ApiGroup struct {
	ArticleApiGroup quickSearch.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
