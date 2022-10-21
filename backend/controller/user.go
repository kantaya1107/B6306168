package controller

import (
	"github.com/kantaya1107/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users

func CreateLeave(c *gin.Context) {

	var leave entity.Leave

	if err := c.ShouldBindJSON(&leave); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&leave).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": leave})

}

// GET /user/:id

func GetLeave(c *gin.Context) {

	var leave entity.Leave

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM leaves WHERE id = ?", id).Scan(&leave).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": leave})

}

// GET /users
func ListDoctor(c *gin.Context) {

	var dt []entity.Doctor

	if err := entity.DB().Raw("SELECT * FROM doctors").Scan(&dt).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": dt})

}

func ListLeave(c *gin.Context) {

	var leave []entity.Leave

	if err := entity.DB().Preload("Doctor").Preload("Evidence").Preload("Type").Raw("SELECT * FROM leaves").Find(&leave).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": leave})

}
func ListType(c *gin.Context) {

	var typ []entity.Type

	if err := entity.DB().Raw("SELECT * FROM Types").Scan(&typ).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": typ})

}
func ListEvidence(c *gin.Context) {

	var ev []entity.Evidence

	if err := entity.DB().Raw("SELECT * FROM evidences").Scan(&ev).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": ev})

}

// DELETE /users/:id

// func DeleteUser(c *gin.Context) {

// 	id := c.Param("id")

// 	if tx := entity.DB().Exec("DELETE FROM users WHERE id = ?", id); tx.RowsAffected == 0 {

// 		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

// 		return

// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": id})

// }

// PATCH /users

// func UpdateUser(c *gin.Context) {

// 	var user entity.User

// 	if err := c.ShouldBindJSON(&user); err != nil {

// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 		return

// 	}

// 	if tx := entity.DB().Where("id = ?", user.ID).First(&user); tx.RowsAffected == 0 {

// 		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

// 		return

// 	}

// 	if err := entity.DB().Save(&user).Error; err != nil {

// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 		return

// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": user})

// }
