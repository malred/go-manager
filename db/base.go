package go_manager_db

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	utils "go_manager_utils"
)

// 插入数据
func Insert(tableName string, params []string, datas ...interface{}) (err error) {
	// 拼接 表名(参数1,参数2,...)
	paramStr := utils.ParamsStr(params)
	// 拼接 values(?,?,...)
	values := utils.ValueStr(len(params))
	var sqlStr = "insert into " + tableName + paramStr + " values" + values
	fmt.Println(sqlStr)
	_, err = db.Exec(sqlStr, datas...) // 要用...展开
	if err != nil {
		fmt.Println(err)
		fmt.Println("插入数据失败")
		return
	}
	return
}

// 删除数据
func Delete(tableName string, id int64) (err error) {
	sqlStr := "delete from " + tableName + " where id=?"
	fmt.Println(sqlStr)
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除数据失败")
		return
	}
	return
}
