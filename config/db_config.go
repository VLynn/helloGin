package config

import "fmt"

const (
	IP       = "192.168.3.66"
	Port     = "3306"
	User     = "intsig"
	Password = "*501AA83B185BC219F61FC7866B755A56198B78E5"
)

var Source = fmt.Sprintf("%s:%s@tcp(%s:%s)/db_user?charset=utf8", User, Password, IP, Port)

// 表名
const (
	TblNameUserProfile = "t_user_profile"
	TblNameUserAccount = "t_user_account"
)
