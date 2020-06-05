package model

import (
    "time"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "helloGin/config"
)

type UserProfile struct {
    Id          int         `json:"id"`
    Name        string      `json:"name" binding:"required"`
    Company     string      `json:"company" binding:"required"`
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

// 插入一条个人信息
func Insert(user UserProfile) int {
    // 建立数据库连接
    db, err := sql.Open("mysql", config.Source)
    if err != nil {
        log.Fatal(err.Error())
    }
    defer db.Close()

    // 执行
    statement := "insert into " + tblName + " (name, company, create_time, update_time) values (?, ?, ?, ?)"
    result, _ := db.Exec(statement, user.Name, user.Company, time.Now().Unix(), time.Now().Unix())

    id, _ := result.LastInsertId()
    return int(id)
}