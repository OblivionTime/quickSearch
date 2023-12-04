package router

import "quickSearch/router/quickSearch"

type RouterGroup struct {
	System quickSearch.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
