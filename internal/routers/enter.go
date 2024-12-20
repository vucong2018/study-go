package routers

import (
	"github.com/vucong2018/study-go/internal/routers/manage"
	"github.com/vucong2018/study-go/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
