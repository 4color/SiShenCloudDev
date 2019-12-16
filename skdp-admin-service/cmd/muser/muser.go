package muser

import (
	"github.com/4color/SiShenCloudDev/skdp-admin-service/cmd/db"
	"strings"
)

type Muser struct {
	Muserid    string `json:"muserid" form:"muserid"`
	Username   string `json:"username" form:"username"`
	pwd        string `json:"pwd" form:"pwd"`
	Truename   string `json:"truename" form:"truename"`
	Created_at string `json:"created_at" form:"created_at"`
	Updated_at string `json:"updated_at" form:"updated_at"`
	Login_time string `json:"login_time" form:"login_time"`
	Enable     int32  `json:"enable" form:"enable"`
}

//用户登录
func (p *Muser) GetUserLogin(username string, pwd string) (user Muser, err error) {

	row := db.SqlDB.QueryRow("SELECT muserid,username,truename,Created_at,Updated_at,Login_time,`Enable` FROM M_USER where username=? and pwd=? and Enable=1 ", username, pwd)
	err = row.Scan(&user.Muserid, &user.Username, &user.Truename, &user.Created_at, &user.Updated_at, &user.Login_time, &user.Enable)

	//如果包含了无行数据的消息，则不为错误
	if (err != nil && strings.Contains(err.Error(), "no rows in result set")) {
		err = nil
	}

	return
}
