package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"test/webook/internal/web"
	"time"
)

type LoginJWTMiddleWareBuilder struct {
	paths []string
}

func NewLoginJWTMiddleWareBuilder() *LoginJWTMiddleWareBuilder {
	return &LoginJWTMiddleWareBuilder{}
}

func (l *LoginJWTMiddleWareBuilder) IgnorePath(path string) *LoginJWTMiddleWareBuilder {
	l.paths = append(l.paths, path)
	return l
}
func (l *LoginJWTMiddleWareBuilder) Build() gin.HandlerFunc {

	return func(context *gin.Context) {
		// @1 拦截不需要校验的路径 (比如登录和注册)
		for _, path := range l.paths {
			if context.Request.URL.Path == path {
				return
			}

		}

		// @2 使用JWT来进行校验
		token := context.GetHeader("token")
		if token == "" {
			// 没有携带token -> 没有登录
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims := &web.UserClaims{}
		tokenStr, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("aB3f9KjL8mNpQrStUvWxYz12345678"), nil
		})
		if err != nil {
			//没有登录
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if tokenStr == nil || !tokenStr.Valid || claims.UID == 0 {
			//没有登录
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims.UserAgent != context.Request.UserAgent() {
			//严重的安全问题 --> 复制token来进行登录操作
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
		//每10s钟刷新一次 这里是用户每次访问的时候,如果已经过了10s,就生成一个新的token返回
		now := time.Now()
		if claims.ExpiresAt.Sub(now) < time.Second*50 {
			newToken, err := tokenStr.SignedString([]byte("aB3f9KjL8mNpQrStUvWxYz12345678"))
			if err != nil {
				// 续约失败
				log.Println("jwt续约失败", err)
			}
			context.Header("x-jwt-token", newToken)
		}

		context.Set("claims", claims)
	}
}
