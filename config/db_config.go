package config

import "fmt"

const (
    IP = "192.168.28.20"
    Port = "3306"
    User = "intsig"
    Password = "intsig"
)

var Source = fmt.Sprintf("%s:%s@tcp(%s:%s)/db_user?charset=utf8", User, Password, IP, Port)