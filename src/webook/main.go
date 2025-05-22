package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"test/webook/internal/repository"
	"test/webook/internal/repository/dao"
	"test/webook/internal/service"
	"test/webook/internal/web"
	"test/webook/internal/web/middleware"
	"test/webook/pkg/ginx/middleware/ratelimit"
	"time"
)

/*
在win上打包go文件的命令:

	SET CGO_ENABLED=0
	SET GOOS=linux
	SET GOARCH=amd64
	go build main.go //todo go build -o name .（将当前命令下的go文件打包为可执行文件）
*/
func main() {
	//db := initDB()
	//server := initWebServer()
	//userHandle := initUser(db)
	//userHandle.RegisterRouters(server)
	//server.Run(":8080")
	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world~")
	})
	server.Run(":8080")
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	/*
		解决跨域问题 --> 使用middleware(中间件,在Java可以被称为拦截器链,过滤器链等组件)
		跨域问题: 请求的协议,域名,端口有一个不一样,浏览器就会拦截(也即跨域问题,prelight请求)
	*/
	//middleware的处理 先于 业务处理
	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"https://foo.com"}, 允许的请求源
		//AllowMethods:     []string{"PUT", "PATCH"}, 允许跨域的请求方法
		AllowHeaders:     []string{"Content-Type"}, //允许的请求头
		ExposeHeaders:    []string{"x-jwt-token"},  //允许前端能够读取后端写回的响应头(在这里只允许读取x-jwt-token)
		AllowCredentials: true,                     //是否允许携带类型于cookie的东西
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") { //如果请求是以这个为前缀的,则允许
				return true
			}
			return strings.Contains(origin, "xxx_domain") //公司的域名
		},
		MaxAge: 12 * time.Hour,
	}))

	//限流 1分钟 100次
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	server.Use(ratelimit.NewBuilder(redisClient,
		time.Second, 100).Build())

	//store := cookie.NewStore([]byte("secret"))
	//server.Use(sessions.Sessions("mysession", store))
	//server.Use(middleware.NewLoginMiddleWareBuilder().
	//	IgnorePath("/users/login").
	//	IgnorePath("/users/signup").
	//	Build())
	server.Use(middleware.NewLoginJWTMiddleWareBuilder().
		IgnorePath("/users/login").
		IgnorePath("/users/signup").
		Build())

	return server
}

func initUser(db *gorm.DB) *web.UserHandle {
	userDAO := dao.NewUserDAO(db)
	userRepo := repository.NewUserRepository(userDAO)
	userService := service.NewUserService(userRepo)
	userHandle := web.NewUserHandle(userService)
	return userHandle
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
