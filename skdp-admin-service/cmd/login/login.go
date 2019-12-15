package login

import (
	"github.com/4color/SiShenCloudDev/skdp-admin-service/cmd"
	"github.com/4color/SiShenCloudDev/skdp-admin-service/cmd/checkcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginModel struct {
	U           string `json:"u"`
	P           string `json:"p"`
	CaptchaId   string `json:"captchaId"`
	VerifyValue string `json:"VerifyValue"`
}

type LoginResult struct {
	Uuid  string `json:"uuid"`
	Token string `json:"token"`
	Name  string `json:"name"`
}

func Login(gc *gin.Context) {
	res := cmd.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	var param LoginModel
	err := gc.BindJSON(&param)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	//判断验证码
	var Codeparam checkcode.ConfigJsonBody
	Codeparam.VerifyValue = param.VerifyValue
	Codeparam.Id = param.CaptchaId
	codeResult := checkcode.CheckCodeVerify(Codeparam)

	if (codeResult["code"] == 0) {
		res.Message = codeResult["msg"].(string)
		gc.JSON(200, res)
		return
	}

	result := LoginResult{"uuid", "token", param.U}
	res.Data = result

	res.Status = http.StatusOK
	gc.JSON(200, res)
}
