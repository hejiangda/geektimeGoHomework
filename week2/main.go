package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type User struct {
	Name string
	Id   int
}
type Dao struct {
	DB *sql.DB
}

// 通过用户id获取用户名
func (d *Dao) GetUserById(id int) (user []User, err error) {

	result := d.DB.QueryRow("SELECT name FROM user WHERE id=" + fmt.Sprint(id) + "")
	if err != nil {
		return user, errors.Wrap(err, "Id:"+fmt.Sprint(id))
	}
	var name string
	err = result.Scan(&name)
	// 处理错误
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.Wrap(err, "Query Id:"+fmt.Sprint(id)+" not found!")
		}
		return user, errors.Wrap(err, "Query Id:"+fmt.Sprint(id))
	}
	// 保存用户数据
	var tmp User
	tmp.Id = id
	tmp.Name = name
	user = append(user, tmp)

	return
}
func main() {

	var d Dao
	db, err := sql.Open("mysql", "happy:hh@tcp(192.168.64.3:3306)/gostudy")
	d.DB = db
	// 程序结束后释放资源
	defer d.DB.Close()
	// 连接失败，直接挂掉
	if err != nil {
		panic(err)
	}
	// 获取id=1的用户sam
	user, err := d.GetUserById(1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user)
	// 获取id=2的用户（不存在）
	user, err = d.GetUserById(2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user)
}
