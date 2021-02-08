package model

import (
	"helloGin/config"
	"helloGin/lib"
)

type UserAccount struct {
	username string
	password string
}

// 根据用户名查询账号信息
func QueryAccount(username string) (account UserAccount) {
	tblName := config.TblNameUserAccount
	statement := "select username, password from " + tblName + " where username = ?"
	rows, _ := lib.Query(statement, username)
	if rows.Next() {
		rows.Scan(&account.username, &account.password)
	}
	return
}
