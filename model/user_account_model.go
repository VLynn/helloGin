package model

import (
	"helloGin/config"
	"helloGin/lib"
	"time"
)

type UserAccount struct {
	Username string
	Password string
}

// 根据用户名查询账号信息
func QueryAccount(username string) (account UserAccount) {
	tblName := config.TblNameUserAccount
	statement := "select username, password from " + tblName + " where username = ?"
	rows, _ := lib.Query(statement, username)
	if rows.Next() {
		rows.Scan(&account.Username, &account.Password)
	}
	return
}

// 创建新账号
func NewAccount(account UserAccount) {
	tblName := config.TblNameUserAccount
	statement := "insert into " + tblName + " (username, password, create_time) values (?, ?, ?)"
	lib.Modify(statement, account.Username, account.Password, time.Now().Unix())
}
