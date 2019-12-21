package xzq

import (
	"github.com/4color/SiShenCloudDev/skdp-admin-service/api"
	"github.com/4color/SiShenCloudDev/skdp-admin-service/cmd/xzq"
	"github.com/gin-gonic/gin"
	"net/http"
)

type XzqList struct {
	Where     string `json:"where"`
	Pagesize  int32  `json:"pagesize"`
	Pageindex int32  `json:"pageindex"`
	Level     int32  `json:"level"`
}

func GetXzqList(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	var param XzqList
	err := gc.BindJSON(&param)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	//检查用户状态，密码是否正确
	data, err := xzq.XzqList(param.Pagesize, param.Pageindex, param.Where, param.Level)

	if (err != nil) {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	res.Data = data
	res.Status = http.StatusOK
	gc.JSON(200, res)
}
