package go_manager_web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "go_manager_db"
	utils "go_manager_utils"
	"net/http"
)

// 定义路由组
// 组中组(嵌套路由组)
func DefineRouteGroup(fatherGroup *gin.RouterGroup, groupName string, r *gin.Engine) *gin.RouterGroup {
	var group *gin.RouterGroup
	if fatherGroup != nil {
		// v1/groupName
		group = fatherGroup.Group(groupName)
	} else {
		// /groupName
		group = r.Group(groupName)
	}
	// 返回路由组
	return group
}

// 存放 token (不同ip不同token)
var TokenMap = make(map[string]string, 10)

// 定时销毁token
func timeDT() {
	// 两小时后销毁
	t := utils.NewMyTimer(2*60*60, func() error {
		utils.DestoryTokenMap(TokenMap)
		return nil
	})
	t.Start()
	fmt.Println(TokenMap)
}

// 路由和处理函数放在不同文件好像会使中间件失效
func Login(c *gin.Context) {
	// 从表单取数据
	// uname := c.PostForm("uname")
	// upass := c.PostForm("upass")
	user := db.MalUser{}
	//绑定json和结构体
	if err := c.BindJSON(&user); err != nil {
		return
	}
	uname := user.Uname
	upass := user.Upass
	// secret := c.Request.Header["Secret"]
	// for k, v := range c.Request.Header {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println(secret[0][9:15])
	userModel, err := db.GetUserByName(uname, upass)
	if err != nil || &userModel == nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"status": 500,
			"msg":    "登录失败",
		})
		return
	}
	// malred := secret[0][8:11] + secret[0][19:22]
	token := utils.SignJWT("malred", uname, upass)
	// 存入map
	// fmt.Println(c.ClientIP(),c.RemoteIP())
	TokenMap[c.ClientIP()] = token
	fmt.Println(TokenMap)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "登录成功",
		// 返回jwt令牌(密码因为前端md5加密过,所以直接放入jwt)
		"token": token,
	})
	go timeDT()
}

// 路由器
// 启动默认的路由
var r = gin.Default()

// user路由组
var v1 *gin.RouterGroup

func Run() {
	// 使用中间件
	// 日志
	r.Use(gin.Logger())
	// 错误恢复
	r.Use(gin.Recovery())
	// 跨域
	r.Use(Core())
	// 阻止缓存响应
	r.Use(NoCache())
	// 安全设置
	r.Use(Secure())
	// 创建路由组v1
	v1 = DefineRouteGroup(nil, "v1", r)
	v1.POST("login", Login)
	// 注册user的路由
	registerUser(Token(), Core())
	// 注册role的路由
	registerRole(Token(), Core())
	// 启动webserver,监听本地127.0.0.1(默认)端口
	r.Run(":10101")
}
