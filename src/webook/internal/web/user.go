package web

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/webook/internal/domain"
	"test/webook/internal/service"
)

// 邮箱和密码的正则表达式
const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

type UserHandle struct {
	userService *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandle(userService *service.UserService) *UserHandle {
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandle{
		userService: userService,
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

// 路由注册
func (u *UserHandle) RegisterRouters(server *gin.Engine) {
	group := server.Group("/users")
	group.POST("/signup", u.SignUp)
	group.POST("/login", u.Login)
	group.POST("/edit", u.Edit)
	group.GET("/profile", u.Profile)
}

// 用户注册
func (u *UserHandle) SignUp(ctx *gin.Context) {
	//内部私有结构体
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req SignUpReq
	//Bind()方法会根据前端传入的content-type(一般是json)来将数据解析到req结构体中
	//如果解析失败,则直接向前端写回4xx错误码
	if err := ctx.Bind(&req); err != nil {
		return
	}

	/*
		here:前端传过来的用户的数据(邮箱和密码就已经在req对象中了),下一步就是校验是否正确
			1.邮箱是否符合格式
			2.密码和确认密码是否一致
			3.密码是否符合格式

	*/
	//@1 校验邮箱格式
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "您的邮箱格式不对")
		return
	}

	//@2 校验两次输入的密码是否一致
	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "您两次输入的密码不一致")
		return
	}
	//@3 校验密码格式
	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "密码必须包含字母、数字、特殊字符,并且不少于八位")
		return
	}

	//next 数据库操作 -> 委托给userService来完成
	err = u.userService.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.ErrDuplicateEmail {
		ctx.String(http.StatusOK, "注册的邮箱重复")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	ctx.String(http.StatusOK, "注册成功")
	fmt.Println(req)
}

// 用户登录
func (u *UserHandle) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	fmt.Println(req)
	user, err := u.userService.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "用户名或者密码错误")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	//登录成功后设置用户对应的session
	session := sessions.Default(ctx)
	session.Set("userId", user.Id)
	session.Save()
	ctx.String(http.StatusOK, "登录成功")

	return
}

// 用户修改个人信息
func (u *UserHandle) Edit(ctx *gin.Context) {

}

// 用户查看个人信息
func (u *UserHandle) Profile(ctx *gin.Context) {

}
