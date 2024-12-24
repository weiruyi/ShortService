package router

import "ShortService/src/internal/router/admin"

type RouterGroup struct {
	admin.ServerRouter
}

var AllRouter = new(RouterGroup)
