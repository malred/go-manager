package go_manager_db

import (
	"fmt" 
	utils "go_manager_utils" 
	_ "github.com/mattn/go-sqlite3"
)

// 查数据
func GetAllUser() (users []*MalUser, err error) {
	sqlStr := `select * from mal_user`
	// 查询,记录到booklist
	err = db.Select(&users, sqlStr)
	if err != nil {
		fmt.Println("查询信息失败")
		fmt.Println(err)
		return
	}
	return
}

// 根据id查数据
func GetUserById(id int64) (user MalUser, err error) {
	// 如果返回的是指针,需要初始化
	//book=&Book{}
	sqlStr := "select * from mal_user where id=?"
	err = db.Get(&user, sqlStr, id)
	if err != nil {
		fmt.Println("查询信息失败")
		return
	}
	return
}

// 根据name查数据
func GetUserByName(uname string, upass string) (user MalUser, err error) {
	sqlStr := "select * from mal_user where uname=? and upass=?"
	err = db.Get(&user, sqlStr, uname, upass)
	if err != nil {
		fmt.Println("查询信息失败")
		return
	}
	return
}

// 根据id改
func UptUserById(uid string, params []string, datas ...interface{}) (err error) {
	// 拼接参数列表 xxx=?,xxx=?
	paramsStr := utils.UptParamsStr(params)
	// uid直接传字符串拼接
	sqlStr := "update mal_role set " + paramsStr + " where id=" + uid
	_, err = db.Exec(sqlStr, datas...)
	if err != nil {
		fmt.Println("修改信息失败")
		return
	}
	return
}
