package muser

import (
	"github.com/4color/SiShenCloudDev/skdp-admin-service/cmd/db"
	"time"
)

func (Muser) TableName() string {
	return "M_USER"
}

type Muser struct {
	Muserid    string    `json:"muserid" form:"muserid" gorm:"primary_key"`
	Username   string    `json:"username" form:"username"`
	pwd        string    `json:"pwd" form:"pwd"`
	Truename   string    `json:"truename" form:"truename"`
	Created_at time.Time `json:"created_at" form:"created_at"`
	Updated_at time.Time `json:"updated_at" form:"updated_at"`
	Login_time time.Time `json:"login_time" form:"login_time"`
	Enable     int32     `json:"enable" form:"enable"`
}

//用户登录
func (p *Muser) GetUserLogin(username string, pwd string) (user Muser, err error) {

	//row := db.SqlDB.Where("username=? and pwd=? and Enable=1 ", username, pwd).First(&user)
	//err = row.Scan(&user.Muserid, &user.Username, &user.Truename, &user.Created_at, &user.Updated_at, &user.Login_time, &user.Enable)

	db.SqlDB.Where(&Muser{Username: username, pwd: pwd, Enable: 1}).First(&user)
	//如果包含了无行数据的消息，则不为错误
	//if (err != nil && strings.Contains(err.Error(), "no rows in result set")) {
	//	err = nil
	//}
	//
	//如果验证成功，则更新用户信息 最后更新日期 login_time
	db.SqlDB.Model(&user).Update(Muser{Login_time: time.Now()})

	return
}
