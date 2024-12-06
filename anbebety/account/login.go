package account

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"project/db"
	"project/model"
	"project/tool"
	"strconv"
)

func Login(c *gin.Context) {
	db := db.Dbfrom()
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	var n int64
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	error := db.Table("users").Where("telephone=?", telephone).Count(&n).Error
	if error != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": error,
		})
		return
	}
	if n == 0 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "用户不存在",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "密码错误",
		})
		return
	}
	f := 1
	for f == 1 {
		value := tool.Randam()
		f = -1
		c.SetCookie(user.Name, strconv.Itoa(value), 3600, "/", "http://127.0.0.1:8080", false, false)
		newsession := model.Session{
			Name:  user.Name,
			Value: strconv.Itoa(value),
		}
		db.Create(&newsession)
		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
		})
		return
	}
	c.JSON(422, gin.H{
		"code":    422,
		"message": "cookie failure",
	})
}
