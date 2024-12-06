package team

import (
	"github.com/gin-gonic/gin"
	"project/db"
	"project/dto"
	"project/model"
)

func Apply(c *gin.Context) {
	db := db.Dbfrom()
	var apply dto.ApplyDto
	c.ShouldBindJSON(&apply)
	var n int64
	error := db.Table("groups").Where("title=?", apply.GroupTitle).Count(&n).Error
	if error != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": error.Error(),
		})
		return
	}
	if n == 0 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "队伍不存在",
		})
		return
	}
	error = db.Table("users").Where("name=?", apply.Name).Count(&n).Error
	if error != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": error.Error(),
		})
		return
	}
	if n == 0 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "请输入正确的用户名",
		})
		return
	}
	newApply := model.Apply{
		Name:       apply.Name,
		GroupTitle: apply.GroupTitle,
		Reason:     apply.Reason,
		Advantage:  apply.Advantage,
		State:      0,
	}
	error = db.Create(&newApply).Error
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
