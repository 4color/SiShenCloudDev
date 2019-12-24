package tuser

import "time"

func (MODEL_T_USER) TableName() string {
	return "T_USER"
}

type MODEL_T_USER struct {
	Userid    string    `json:"userid" gorm:"column:USERID;primary_key"`
	Pwd       string    `json:"pwd" gorm:"column:PWD"`
	Email     string    `json:"email" gorm:"column:EMAIL"`
	Mobile    string    `json:"mobile" gorm:"column:MOBILE"`
	Weixin    string    `json:"weixin" gorm:"column:WEIXIN"`
	Qq        string    `json:"qq" gorm:"column:QQ"`
	Nickname  string    `json:"nickname" gorm:"column:NICKNAME"`
	CreatedAt time.Time `json:"createdat" gorm:"column:CREATEDAT"`
	UpdatedAt time.Time `json:"updatedat" gorm:"column:UPDATEDAT"`
	DeletedAt time.Time `json:"deletedat" gorm:"column:DELETEDAT"`
	Enable    int       `json:"enable" gorm:"column:ENABLE"`
	Isdel     int       `json:"isdel" gorm:"column:ISDEL"`
	Xzqid     int       `json:"xzqid" gorm:"column:XZQID"`
	Source    string    `json:"source" gorm:"column:SOURCE"`
}
