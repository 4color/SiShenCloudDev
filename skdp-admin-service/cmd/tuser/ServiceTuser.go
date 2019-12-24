package tuser

import (
	"github.com/4color/SiShenCloudDev/skdp-admin-service/cmd/db"
	"github.com/pkg/errors"
)

type XzqListResult struct {
	Count int32       `json:"count"`
	Data  []MODEL_T_USER `json:"data"`
}

type QueryUsersList struct {
	Where     string `json:"where"`
	Pagesize  int    `json:"pagesize"`
	Pageindex int    `json:"pageindex"`
}

//获取用户列表,第几页
func UsersList(query QueryUsersList) (result XzqListResult, err error) {

	if (query.Pagesize > 100) {
		err = errors.New("每页数量不能超过100")
		return
	}

	result.Data = make([]MODEL_T_USER, 0)
	result.Count = 0

	var thisdb = db.SqlDB

	if (query.Where != "") {
		thisdb = thisdb.Where("nickname like ?", "%"+query.Where+"%").Or("qq like ?", "%"+query.Where+"%").Or("email like ?", "%"+query.Where+"%")
	}

	thisdb.Find(&result.Data).Count(&result.Count)

	if query.Pagesize > 0 && query.Pagesize > 0 {
		thisdb = thisdb.Limit(query.Pagesize).Offset((query.Pageindex - 1) * query.Pagesize)
	}

	thisdb.Find(&result.Data)

	return
}
