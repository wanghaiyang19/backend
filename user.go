package dao

import (
	"database/sql"
	"fmt"
	"learn/model"
	"log"
)

var database = map[string]string{
	"yxh": "123456",
	"wx":  "654321",
}

func AddUser(username, password string) {
	database[username] = password
}

func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	return database[username]
}

var db *sql.DB

func initDB() {
	var err error
	dsn := "root:Why729831@mysql@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
	return
}
func InsertStudent(st model.User) {
	sqlStr := "insert into user(username,password ) values (?,?)"
	_, err := db.Exec(sqlStr, st.Username, st.Password)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

}
