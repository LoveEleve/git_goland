package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddleWareBuilder struct {
	paths []string
}

func NewLoginMiddleWareBuilder() *LoginMiddleWareBuilder {
	return &LoginMiddleWareBuilder{}
}

func (l *LoginMiddleWareBuilder) IgnorePath(path string) *LoginMiddleWareBuilder {
	l.paths = append(l.paths, path)
	return l
}
func (l *LoginMiddleWareBuilder) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		//@1 另外一种写法 不需要校验的
		for _, path := range l.paths {
			if context.Request.URL.Path == path {
				return
			}

		}
		//@1登录和注册不需要做权限校验(登录)
		if context.Request.URL.Path == "/users/signup" || context.Request.URL.Path == "/users/login" {
			return
		}
		session := sessions.Default(context)
		id := session.Get("userId")
		if id == nil {
			context.AbortWithStatus(http.StatusUnauthorized) //没有登录
			return
		}
	}
}
