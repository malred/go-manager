package main

import ( 
    db "go_manager_db"
    web "go_manager_web"
)

func main() {
    // 初始化数据库
    db.InitDB()   
    // 开启服务
    web.Run()
}