package xzq

import (
	"github.com/4color/SiShenCloudDev/skdp-admin-service/api"
	"github.com/4color/SiShenCloudDev/skdp-admin-service/cmd/xzq"
	"github.com/gin-gonic/gin"
	"net/http"
)

type XzqList struct {
	Where             string `json:"where"`
	Pagesize          int  `json:"pagesize"`
	Pageindex         int  `json:"pageindex"`
	Level             int  `json:"level"`
	Parentid          string `json:"parentid"`
	ParentlikeOrEqual string `json:"parentlikeorequal"`
}

//行政区列表
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
	data, err := xzq.XzqList(param.Pagesize, param.Pageindex, param.Where, param.Level, param.Parentid, param.ParentlikeOrEqual)

	if (err != nil) {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	res.Data = data
	res.Status = http.StatusOK
	gc.JSON(200, res)
}

//行政区新增或编辑
func XzqAdd(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	var param xzq.MODEL_XZQ_ENUM
	err := gc.BindJSON(&param)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	data, err := xzq.XzqAdd(param)

	if (err != nil) {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	res.Data = data
	res.Message = "编辑成功"
	res.Status = http.StatusOK
	gc.JSON(200, res)
}

//取一个行政区信息
func GetXzqEntity(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	xzqid := gc.Param("id")
	xzqvalue := gc.Param("xzqid")

	var param xzq.MODEL_XZQ_ENUM
	param.Xzqid = xzqid;
	param.Xzqvalue = xzqvalue;
	data, err := xzq.XzqAddEntity(param)
	if (err != nil) {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}
	res.Data = data
	res.Message = "获取成功"
	res.Status = http.StatusOK
	gc.JSON(200, res)
}

//删除行政区
func DeleteXzq(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	xzqid := gc.Param("id")

	data, err := xzq.XzqDelete(xzqid)
	if (err != nil) {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}
	res.Data = data
	res.Message = "删除成功"
	res.Status = http.StatusOK
	gc.JSON(200, res)
}

//删除行政区
func EnableXzq(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	xzqid := gc.Param("id")
	enable := gc.Param("enable")


	data, err := xzq.XzqUpdateEnable(xzqid,enable)
	if (err != nil) {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}
	if(data) {
		res.Data = data
		res.Status = http.StatusOK
		res.Message = "修改成功"
	}else {
		res.Message = "修改失败"
	}
	gc.JSON(200, res)
}

