package tuser

import (
	"github.com/4color/SiShenCloudDev/skdp-admin-service/api"
	"github.com/4color/SiShenCloudDev/skdp-admin-service/cmd/tuser"
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户列表
func ControllerUserList(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	var param tuser.QueryUsersList
	err := gc.BindJSON(&param)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	//检查用户状态，密码是否正确
	data, err := tuser.UsersList(param)

	if (err != nil) {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	res.Data = data
	res.Status = http.StatusOK
	gc.JSON(200, res)
}

