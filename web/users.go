package go_manager_web

import (
	"fmt"
	db "go_manager_db"
	utils "go_manager_utils"
	"strconv" 

	"github.com/gin-gonic/gin"
)

func GetAllUserHandler(c *gin.Context) {
	users, err := db.GetAllUser()
	// 通用响应
	utils.R(c, err, "查询角色失败", users)
}
func AddUserHandler(c *gin.Context) {
	// uname := c.PostForm("uname")
	// upass := c.PostForm("upass")
	// idStr := c.PostForm("rid")
	user := db.MalUser{}
	//绑定json和结构体
	if err := c.BindJSON(&user); err != nil {
		return
	}
	uname := user.Uname
	upass := user.Upass
	rid := user.Rid
	fmt.Println(user)
	// rid, err := strconv.ParseInt(idStr, 10, 64)
	err := db.Insert("mal_user", []string{"uname", "upass", "rid"}, uname, upass, rid)
	// 通用响应
	utils.R(c, err, "添加角色失败", "添加角色成功")
}
func DelUserHandler(c *gin.Context) {
	// 从url获取参数
	idStr := c.Query("uid")
	// fmt.Println(idStr)
	uid, err := strconv.ParseInt(idStr, 10, 64)
	err = db.Delete("mal_user", uid)
	// 通用响应
	utils.R(c, err, "删除角色失败", "删除角色成功")
}
func GetOneUserHandler(c *gin.Context) {
	// 从url获取参数
	idStr := c.Query("uid")
	fmt.Println(idStr)
	uid, _ := strconv.ParseInt(idStr, 10, 64)
	one, err2 := db.GetUserById(uid)
	// 通用响应
	utils.R(c, err2, "查询角色失败", one)
}
func UptUserHandler(c *gin.Context) {
	// 从url获取参数
	// uid := c.PostForm("uid")
	// uname := c.PostForm("uname")
	// upass := c.PostForm("upass")
	// ridStr := c.PostForm("rid")
	user := db.MalUser{}
	//绑定json和结构体
	if err := c.BindJSON(&user); err != nil {
		return
	}
	uname := user.Uname
	upass := user.Upass
	rid := user.Rid
	uid := user.Id
	// fmt.Println(idStr, UserName)
	// rid, _ := strconv.ParseInt(ridStr, 10, 64)
	err := db.UptUserById(strconv.FormatInt(uid, 10), []string{"uname", "upass", "rid"}, uname, upass, rid)
	// 通用响应
	utils.R(c, err, "修改角色失败", "修改角色成功")
}
func registerUser(middles ...gin.HandlerFunc) {
	// 创建路由组v1/user
	user := DefineRouteGroup(v1, "user", r)
	// 添加中间件
	if middles != nil {
		user.Use(middles...)
	}
	user.GET("all", GetAllUserHandler)
	// 添加
	user.POST("add", AddUserHandler)
	// 删除
	user.DELETE("del", DelUserHandler)
	// 根据id获取
	user.GET("id", GetOneUserHandler)
	// 根据id修改
	user.PUT("upt", UptUserHandler)
}
