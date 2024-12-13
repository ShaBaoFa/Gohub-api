package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/event"
	"gohub/pkg/helpers"
	"gohub/pkg/ip2region"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	LoginEvent = "user.login"

	LoginStatusSuccess = 1
	LoginStatusFailed  = 2
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseAPIController
}

type LoginLog struct {
	LoginId string `json:"login_id"`
}

// LoginByPhone 手机登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		event.FireC(LoginEvent, CreateLoginInfo(c, request.Phone, LoginStatusFailed))
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在")
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		event.FireC(LoginEvent, CreateLoginInfo(c, request.Phone, LoginStatusSuccess))
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

// LoginByPassword 多种方法登录，支持手机号、email 和用户名
func (lc *LoginController) LoginByPassword(c *gin.Context) {
	// 1. 验证表单
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(c, "账号不存在或密码错误")
		event.FireC(LoginEvent, CreateLoginInfo(c, request.LoginID, LoginStatusFailed))
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		event.FireC(LoginEvent, CreateLoginInfo(c, request.LoginID, LoginStatusSuccess))
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

// RefreshToken 刷新 Access Token
func (lc *LoginController) RefreshToken(c *gin.Context) {

	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

func CreateLoginInfo(c *gin.Context, name string, status int) map[string]interface{} {
	message := "登录失败"
	if status == LoginStatusSuccess {
		message = "登录成功"
	}
	ip := c.ClientIP()
	userAgent := c.Request.UserAgent()
	return map[string]interface{}{
		"name":        name,
		"ip":          ip,
		"ip_location": ip2region.Search(ip),
		"os":          helpers.AgentToOs(userAgent),
		"browser":     helpers.AgentToBrowser(userAgent),
		"status":      status,
		"message":     message,
		"login_time":  time.Now(),
	}
}
