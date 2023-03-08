package go_manager_db

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// 应该id越大,权限越高,比较方便区分权限
// user < admin < super < root
// 查数据
func GetAllRole() (roles []*MalRole, err error) {
	sqlStr := `select * from mal_role`
	// 查询,记录到booklist
	err = db.Select(&roles, sqlStr)
	if err != nil {
		fmt.Println("查询信息失败")
		fmt.Println(err)
		return
	}
	return
}

// 根据id查数据
func GetRoleById(id int64) (role MalRole, err error) {
	// 如果返回的是指针,需要初始化
	//book=&Book{}
	sqlStr := "select * from mal_role where id=?"
	err = db.Get(&role, sqlStr, id)
	if err != nil {
		fmt.Println("查询信息失败")
		return
	}
	return
}

// 根据id改数据
func UptRoleById(id int64, roleName string) (err error) {
	// 如果返回的是指针,需要初始化
	//book=&Book{}
	sqlStr := "update mal_role set role=? where id=?"
	_, err = db.Exec(sqlStr, roleName, id)
	if err != nil {
		fmt.Println("修改信息失败")
		return
	}
	return
}
