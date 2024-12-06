package account

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"project/db"
	"project/model"
)

func Register(c *gin.Context) {
	db := db.Dbfrom()
	var use model.User
	use.Name = c.PostForm("name")
	use.Password = c.PostForm("password")
	use.Telephone = c.PostForm("telephone")
	use.PersonalInformation = c.PostForm("personal_information")
	IsAdministrators := c.PostForm("isAdministrators")
	if IsAdministrators == "T" {
		use.Identity = 1
	}
	if IsAdministrators == "F" {
		use.Identity = 0
	}
	if len(use.Name) == 0 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "用户名不能为空",
		})
		return
	}
	if len(use.Telephone) != 11 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "手机号必须为11位",
		})
		return
	}
	if len(use.Password) <= 6 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}
	var n int64
	error := db.Table("users").Where("telephone=?", use.Telephone).Count(&n).Error
	if error != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": error.Error(),
		})
		return
	}
	if n != 0 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "用户已存在",
		})
		return
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(use.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}
	use.Password = string(hasedPassword)
	error = db.Create(&use).Error
	if error != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
	})

}
