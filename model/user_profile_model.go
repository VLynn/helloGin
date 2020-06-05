package model

import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "helloGin/config"
)

type UserProfile struct {
    Id          int         `json:"id"`
    Name        string      `json:"name"`
    Company     string      `json:"compnay"`
    Position    string      `json:"position"`
    Avatar      string      `json:"avatar"`
}

const tblName = "t_user_profile"

// 批量获取个人信息
func GetList(offset, num int) []UserProfile {
    // 建立数据库连接
    db, err := sql.Open("mysql", config.Source)
    if err != nil {
        log.Fatal(err.Error())
    }
    // 推迟断开数据库连接
    defer db.Close()

    // 查询
    statement := "select id, name, company, position, avatar from " + tblName + " limit ?, ?"
    rows, err := db.Query(statement, offset, num)
    if err != nil {
        log.Fatal(err.Error())
    }
    defer rows.Close()

    // 遍历rows，并添加到结果集中
    users := make([]UserProfile, 0, num)
    for rows.Next() {
        var user UserProfile
        rows.Scan(&user.Id, &user.Name, &user.Company, &user.Position, &user.Avatar)
        users = append(users, user)
    }
    return users
}