package router

import (
	"github.com/gin-gonic/gin"
	"project/account"
	"project/information"
	"project/logout"
	"project/private_message"
	"project/team"
)

func Router() {
	r := gin.Default()
	v := r.Group("/v1")
	{
		v.POST("/register", account.Register) //注册
		v.POST("/login", account.Login)       //登录
	}
	i := r.Group("/team_up_information")
	{
		i.POST("/release", information.Release) //组队信息发表
		i.POST("/examine", information.Examine) //组队信息审核
	}
	t := r.Group("/team")
	{
		t.POST("/apply", team.Apply)             //申请加入
		t.POST("/select", team.Select)           //选成员
		t.POST("/checkmember", team.Checkmember) //看成员消息
		t.POST("/notice", team.Notice)           //发送团队消息
	}
	p := r.Group("/private_message")
	{
		p.POST("/send", private_message.Send)   //发私信
		p.POST("/check", private_message.Check) //看私信
	}
	r.GET("/deletetest", logout.Logout) //输入名字退出
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
