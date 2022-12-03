package main

import (
	"database/sql" //标准库
	"fmt"
	_ "github.com/go-sql-driver/mysql" //我们使用的mysql，需要导入相应驱动包，否则会报错
	"log"
	"std/model"
)

// 定义一个全局对象db
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
func InsertStudent(st model.Student) {
	sqlStr := "insert into student(name,age,sex ) values (?,?,?)"
	_, err := db.Exec(sqlStr, st.Name, st.Age, st.Sex)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	
}

func queryMultiRowDemo(id int) {
	sqlStr := "select id, name, age from student where id > ?"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u model.Student
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
	}
}
func main() {
	//初始化连接
	initDB()
	for i := 1; i <= 10; i++ {
		st := model.Student{
			Name: "小泽",
			Age:  i,
			Sex:  "男",
		}
		InsertStudent(st)
	}

	queryMultiRowDemo(1)
}
