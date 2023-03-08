package go_manager_db

/*
	创建数据库
		// 进入sqlite命令行
		sqlite3 xxx.db
		// 创建数据库
		.open xxx.db
*/

// 专门定义与数据库交互的结构体
type MalUser struct {
	Id    int64  `db:"id" json:"Rd"`
	Uname string `db:"uname" json:"Uname"`
	Upass string `db:"upass" json:"Upass"`
	Rid   int64  `db:"rid" json:"Rid"`
}

type MalRole struct {
	Id   int64  `db:"id" json:"Id"`
	Role string `db:"role" json:"Role"`
}
