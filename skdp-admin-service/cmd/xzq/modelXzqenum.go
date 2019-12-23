package xzq

import (
	"github.com/4color/SiShenCloudDev/skdp-admin-service/cmd/db"
	"github.com/pkg/errors"
	"time"
)

func (MODEL_XZQ_ENUM) TableName() string {
	return "XZQ_ENUM"
}

type MODEL_XZQ_ENUM struct {
	Xzqid      string    `json:"xzqid" gorm:"column:XZQ_ID;primary_key"`
	Xzqvalue   string    `json:"xzqvalue" gorm:"column:XZQ_VALUE"`
	Xzqmc      string    `json:"xzqmc" gorm:"column:XZQ_MC"`
	Parentid   string    `json:"parentid" gorm:"column:PARENT_ID"`
	Xzqorder   int32     `json:"xzqorder" gorm:"column:XZQ_ORDER"`
	Enable     int32     `json:"enable" gorm:"column:ENABLE"`
	Xzqnf      int32     `json:"xzqnf" gorm:"column:XZQ_NF"`
	Changetime time.Time `json:"changetime" gorm:"column:CHANGE_TIME"`
	Xzqlevel   int32     `json:"xzqlevel" gorm:"column:XZQ_LEVEL"`
}

type XzqListResult struct {
	Count int32            `json:"count"`
	Data  []MODEL_XZQ_ENUM `json:"data"`
}

//获取行政区列表,第几页
func XzqList(pagesize int32, pageindex int32, where string, level int32, parentid string, ParentlikeOrEqual string) (result XzqListResult, err error) {

	if (pagesize > 100) {
		err = errors.New("每页数量不能超过100")
		return
	}

	result.Data = make([]MODEL_XZQ_ENUM, 0)
	result.Count = 0

	var thisdb = db.SqlDB

	if (level > 0) {
		thisdb = thisdb.Where(&MODEL_XZQ_ENUM{Xzqlevel: level})
	}

	if (where != "") {
		thisdb = thisdb.Where("xzq_mc like ?", "%"+where+"%").Or("xzq_value like ?", "%"+where+"%")
		//db.SqlDB.Debug().Where("xzq_mc like ?", "%"+where+"%").Or("xzq_value like ?", "%"+where+"%").Find(&list)
	}

	if (ParentlikeOrEqual == "equal") {
		thisdb = thisdb.Where("PARENT_ID = ?", parentid)
	} else {
		thisdb = thisdb.Where("PARENT_ID like ?", ""+parentid+"%")
	}

	thisdb.Find(&result.Data).Count(&result.Count)

	if pagesize > 0 && pageindex > 0 {
		thisdb = thisdb.Limit(pagesize).Offset((pageindex - 1) * pagesize)
	}

	thisdb.Find(&result.Data)

	return
}

//获取行政区列表,第几页
func XzqAdd(xzqModel MODEL_XZQ_ENUM) (result MODEL_XZQ_ENUM, err error) {

	if (xzqModel.Xzqid == "") {

		xzqModel.Xzqid = xzqModel.Xzqvalue
		db.SqlDB.Create(&xzqModel)
		db.SqlDB.NewRecord(xzqModel)
		result = xzqModel

	} else {
		db.SqlDB.Where(&MODEL_XZQ_ENUM{Xzqid: xzqModel.Xzqid}).First(&result)
		db.SqlDB.Model(&result).Updates(xzqModel)
	}

	return
}

//获取一个行政区
func XzqAddEntity(xzqModel MODEL_XZQ_ENUM) (result MODEL_XZQ_ENUM, err error) {

	if (xzqModel.Xzqid != "") {
		db.SqlDB.Where(&MODEL_XZQ_ENUM{Xzqid: xzqModel.Xzqid}).First(&result)
	}

	if (xzqModel.Xzqvalue != "") {
		db.SqlDB.Where(&MODEL_XZQ_ENUM{Xzqvalue: xzqModel.Xzqvalue}).First(&result)
	}

	return
}
