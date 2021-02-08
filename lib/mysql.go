package lib

import (
	"database/sql"
	"helloGin/config"
	"log"
)

var db *sql.DB

/* 思考将db封装到struct中的好处

type MysqlDB struct {
	db *sql.DB
}

// 初始化MysqlDB结构体，连接数据库生成DB的事情交给main程序显式地处理
func InitMysqlDB() *MysqlDB {
	// 建立数据库连接
	db, err := sql.Open("mysql", config.Source)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &MysqlDB{db: db}
}

*/

/*
 * 建立数据库连接
 */
func InitMysql() {
	if db == nil {
		var err error
		db, err = sql.Open("mysql", config.Source)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

/*
 * 封装mysql写操作
 * @param statement sql语句
 * @param args ?占位符对应的变量
 * @return 写操作影响的行数, 错误
 */
func Modify(statement string, args ...interface{}) (affectedRows int64, err error) {
	result, err := db.Exec(statement, args)
	if err != nil {
		log.Println("exec sql failed, err: ", err)
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

/*
 * 封装mysql读操作
 * @param statement sql语句
 * @param args ?占位符对应的变量
 */
func Query(statement string, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = db.Query(statement, args)
	if err != nil {
		log.Println("query sql failed, err: ", err)
	}
	return
}
