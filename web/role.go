package go_manager_web

import (
	"fmt"
	db "go_manager_db"
	utils "go_manager_utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRoleHandler(c *gin.Context) {
	roles, err := db.GetAllRole()
	// 通用响应
	utils.R(c, err, "获取角色列表失败", roles)
}
func AddRoleHandler(c *gin.Context) {
	// Role := c.PostForm("Role")
	// fmt.Println(Role)
	role := db.MalRole{}
	//绑定json和结构体
	if err := c.BindJSON(&role); err != nil {
		return
	}
	Role := role.Role
	err := db.Insert("mal_role", []string{"role"}, Role)
	// 通用响应
	utils.R(c, err, "添加角色失败", "添加角色成功")
}
func DelRoleHandler(c *gin.Context) {
	// 从url获取参数
	idStr := c.Query("rid")
	// fmt.Println(idStr)
	rid, err := strconv.ParseInt(idStr, 10, 64)
	err = db.Delete("mal_role", rid)
	// 通用响应
	utils.R(c, err, "删除角色失败", "删除角色成功")
}
func GetOneRoleHandler(c *gin.Context) {
	// 从url获取参数
	idStr := c.Query("rid")
	fmt.Println(idStr)
	rid, _ := strconv.ParseInt(idStr, 10, 64)
	one, err2 := db.GetRoleById(rid)
	// 通用响应
	utils.R(c, err2, "查询角色失败", one)
}
func UptRoleHandler(c *gin.Context) {
	role := db.MalRole{}
	//绑定json和结构体
	if err := c.BindJSON(&role); err != nil {
		return
	}
	// 获取参数
	// idStr := c.PostForm("rid")
	// roleName := c.PostForm("roleName")
	// fmt.Println(idStr, roleName)
	// rid, _ := strconv.ParseInt(idStr, 10, 64)
	rid := role.Id
	roleName := role.Role
	fmt.Println(role)
	err := db.UptRoleById(rid, roleName)
	// 通用响应
	utils.R(c, err, "修改角色失败", "修改角色成功")
}
func registerRole(middles ...gin.HandlerFunc) {
	// 创建路由组v1/user
	role := DefineRouteGroup(v1, "role", r)
	// 添加中间件
	if middles != nil {
		role.Use(middles...)
	}
	// 获取所有
	role.GET("all", GetAllRoleHandler)
	// 添加
	role.POST("add", AddRoleHandler)
	// 删除
	role.DELETE("del", DelRoleHandler)
	// 根据id获取
	role.GET("id", GetOneRoleHandler)
	// 根据id修改
	role.PUT("upt", UptRoleHandler)
}
