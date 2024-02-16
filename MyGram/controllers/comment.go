package controllers

import (
	"Mygram/database"
	"Mygram/helpers"
	"Mygram/models"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CommentRepository struct {
	DB *gorm.DB
}

func CreateComment(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	commentRequest := models.CreateCommentRequest{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	comment := models.Comment{
		PhotoId: commentRequest.PhotoId,
		Message: commentRequest.Message,
		UserId:  userID,
	}

	photos := models.Photo{}

	if comment.ID != photos.ID {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "photo not found",
		})
		return
	}

	err := db.Debug().Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	commentString, _ := json.Marshal(comment)
	commentResponse := models.CreateCommentResponse{}
	json.Unmarshal(commentString, &commentResponse)

	c.JSON(http.StatusCreated, commentResponse)
}

func GetComment(c *gin.Context) {
	db := config.GetDB()

	comments := []models.Comment{}

	err := db.Debug().Preload("User").Preload("Photo").Order("id asc").Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	commentsString, _ := json.Marshal(comments)
	commentsResponse := []models.CommentResponse{}
	json.Unmarshal(commentsString, &commentsResponse)

	c.JSON(http.StatusOK, commentsResponse)
}

func GetCommentByCommentId(c *gin.Context) {
	db := config.GetDB()

	commentID, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": "Invalid comment ID",
		})
		return
	}

	comment := models.Comment{}
	err = db.Preload("User").Preload("Photo").First(&comment, commentID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Comment not found",
			"error_message": fmt.Sprintf("Comment with ID %v not found", commentID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}

func UpdateComment(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	commentRequest := models.UpdateCommentRequest{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	comment := models.Comment{}
	comment.ID = uint(commentId)
	comment.UserId = userID

	updateString, _ := json.Marshal(commentRequest)
	updateData := models.Comment{}
	json.Unmarshal(updateString, &updateData)

	err := db.Model(&comment).Updates(updateData).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	_ = db.First(&comment, comment.ID).Error

	commentString, _ := json.Marshal(comment)
	commentResponse := models.UpdateCommentResponse{}
	json.Unmarshal(commentString, &commentResponse)

	c.JSON(http.StatusOK, commentResponse)
}

func DeleteComment(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	comment := models.Comment{}
	comment.ID = uint(commentId)
	comment.UserId = userID

	err := db.Delete(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
