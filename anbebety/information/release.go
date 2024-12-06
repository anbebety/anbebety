package information

import (
	"github.com/gin-gonic/gin"
	"project/db"
	"project/dto"
	"project/model"
)

func Release(c *gin.Context) {
	db := db.Dbfrom()
	var release dto.ReleaseDto
	if err := c.ShouldBindJSON(&release); err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	if release.Number <= 1 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "队伍人数异常",
		})
		return
	}
	var n int64
	err := db.Table("groups").Where("title=?", release.Title).Count(&n).Error
	if err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": err.Error(),
		})
	}
	if n != 0 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "队伍标题重复",
		})
		return
	}
	var user model.User
	errs := db.Where("name = ?", release.Name).First(&user).Error
	if user.ID == 0 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "请输入正确的用户名",
		})
		return
	}
	if errs != nil {
		c.JSON(400, gin.H{
			"code": 400,
		})
		return
	}
	newGroup := model.Group{
		Name:     release.Name,
		Title:    release.Title,
		Aim:      release.Aim,
		Time:     release.Time,
		Location: release.Location,
		Require:  release.Require,
		Number:   release.Number,
		State:    0,
	}
	errs = db.Create(&newGroup).Error
	if errs != nil {
		c.JSON(400, gin.H{
			"code": 400,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
	})
}
