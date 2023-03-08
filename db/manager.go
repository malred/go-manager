// db\main.go
package go_manager_db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// 数据库相关操作
var db *sqlx.DB

// 初始化数据库连接
func InitDB() (err error) {
	dsn := "./manager.db"
	// 连接
	// Open可能仅校验参数，而没有与db间创建连接，
	// 要确认db是否可用，需要调用Ping。Connect则相当于Open+Ping。
	db, err = sqlx.Connect("sqlite3", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	// 最大连接数
	db.SetMaxOpenConns(100)
	// 最大空闲连接数
	db.SetMaxIdleConns(16)
	CreateRoleTable()
	CreateUserTable()
	return
}

// 创建用户表
func CreateUserTable() error {
	sqlc := `
	CREATE TABLE IF NOT EXISTS "mal_user" (
	  -- sqlite 不能用 comment 添加注释
	  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT , -- '主键'
	  "uname" varchar(20) NOT NULL UNIQUE , -- '用户昵称'
	  "upass" varchar(50) NOT NULL, -- '密码(md5加密)'
	  "rid" INTEGER NOT NULL UNIQUE DEFAULT 1 -- '角色id'
	); 
    `
	_, err := db.Exec(sqlc)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 初始化表
	//因为有unique约束,所以不会重复添加
	// sqlStr := "insert into mal_user(uname,upass,rid) values(?,?,?)"
	Insert("mal_user", []string{"uname", "upass", "rid"}, "admin", "e120012d113ff6ea124a2493453c6dd5", 2)
	return nil
}

// 创建权限表
func CreateRoleTable() error {
	sqlc := `
	CREATE TABLE IF NOT EXISTS "mal_role" (
	  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, -- '主键' 
	  "role" varchar(20) NOT NULL UNIQUE DEFAULT 'user' -- '角色(权限)' 
	); 
    `
	_, err := db.Exec(sqlc)
	if err != nil {
		return err
	}
	// 初始化表
	// 因为有unique约束,所以不会重复添加
	// sqlStr := "insert into mal_role(role) values(?)"
	Insert("mal_role", []string{"role"}, "user")
	Insert("mal_role", []string{"role"}, "admin")
	Insert("mal_role", []string{"role"}, "super")
	Insert("mal_role", []string{"role"}, "root")
	return nil
}
 