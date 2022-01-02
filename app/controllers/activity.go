package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"app/models"
	"time"
)

type CreateActivityGroupsInput struct {
	Title	string `json:"title"`
	Email	string `json:"email"`
}

type UpdateActivityGroupsInput struct {
	Title	string `json:"title"`
}

type CreatedResponse struct {
	CreatedAt	time.Time	`json:"created_at"`
    UpdatedAt 	time.Time 	`json:"updated_at"`
    Id 			uint 		`json:"id"`
    Title 		string		`json:"title"`
	Email		string	    `json:"email"`
}
type Response struct {
	Status  string `json:"status" gorm:"default:"Not Found""`
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}

type NullResponse struct {}

var responses Response
func FindActivityGroups(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var activity []models.Activity
	db.Find(&activity)

	c.JSON(http.StatusOK, Response{
		Status:"Success",
		Message:"Success",
		Data:activity,
	})
}

func FindActivityGroup(c *gin.Context) {
	var activity models.Activity
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&activity).Error;
	 
	err != nil {
        c.JSON(http.StatusNotFound, Response{
			Status: "Not Found",
			Message: "Activity with ID "+c.Param("id")+" Not Found",
			Data: NullResponse{},
		})
        return
    }	
    c.JSON(http.StatusOK, Response{
		Status:"Success",
		Message:"Success",
		Data:activity,
	})
}

func CreateActivityGroups(c *gin.Context) {
	var input CreateActivityGroupsInput
	if err := c.ShouldBindJSON(&input);
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	activity := models.Activity{Title: input.Title, Email: input.Email}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&activity)
	

	c.JSON(http.StatusOK, Response{
		Status:"Success",
		Message:"Success",
		Data:CreatedResponse{
			CreatedAt: activity.CreateAt,
			UpdatedAt: activity.UpdateAt,
			Id: activity.Id,
			Title: activity.Title,
			Email: activity.Email,
		},
	})
}

func UpdateActivityGroups(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var activity models.Activity
    if err := db.Where("id = ?", c.Param("id")).First(&activity).Error; err != nil {
        c.JSON(http.StatusNotFound, Response{
			Status: "Not Found",
			Message: "Activity with ID "+c.Param("id")+" Not Found",
			Data: NullResponse{},
		})
        return
    }

    // Validate input
    var input UpdateActivityGroupsInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if input.Title == "" {
		c.JSON(http.StatusBadRequest, Response{
			Status:"Bad request",
			Message:"title cannot be null",
			Data:NullResponse{},
		})
		return
	}

    var updatedInput models.Activity
    updatedInput.Title = input.Title
	updatedInput.UpdateAt = time.Now()

    db.Model(&activity).Updates(updatedInput)

    c.JSON(http.StatusCreated, Response{
		Status:"Success",
		Message:"Success",
		Data:activity,
	})
}

func DeleteActivityGroups(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var activity models.Activity
    if err := db.Where("id = ?", c.Param("id")).First(&activity).Error; err != nil {
        c.JSON(http.StatusNotFound, Response{
			Status: "Not Found",
			Message: "Activity with ID "+c.Param("id")+" Not Found",
			Data: NullResponse{},})
        return
    }

    db.Delete(&activity)

    c.JSON(http.StatusOK, Response{
		Status: "Success",
		Message: "Success",
		Data: NullResponse{},
	})
}