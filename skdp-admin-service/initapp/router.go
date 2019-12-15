package initapp

import (
	"fmt"
	"github.com/4color/SiShenCloudDev/skdp-admin-service/cmd/checkcode"
	"github.com/4color/SiShenCloudDev/skdp-admin-service/cmd/login"
	"github.com/gin-gonic/gin"
	"net/http"
)

var R *gin.Engine
func InitRouter()  {

	R = gin.Default()

	// 使用跨域中间件
	R.Use(Cors())

	v1 := R.Group("/api/v1")
	{
		//验证码
		v1.GET("/checkcode", checkcode.Checkcode)
		v1.POST("/checkcode", checkcode.CheckCodeVerifyHandle)

		//登陆
		v1.POST("/login", login.Login)
	}
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,x-skdp-token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}

		// 处理请求
		c.Next()
	}
}