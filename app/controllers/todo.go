package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"app/models"
	"time"
	"strconv"
<<<<<<< HEAD
	"gorm.io/hints"
=======
	
>>>>>>> 24cf106e8429343bc23ff7b53f8f5db7a81aece3
)

type CreateToDoInput struct {
	ActivityGroupId	int `json:"activity_group_id"`
	Title			string `json:"title"`

}

type UpdateToDoInput struct {
	Title		string	`json:"title"`
	IsActive	bool	`json:"is_active"`
}

type CreatedToDoResponse struct {
	CreatedAt		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
	Id				uint		`json:"id"`
	Title			string		`json:"title"`
	ActivityGroupId int			`json:"activity_group_id":`
	IsActive		bool		`json:"is_active"`
	Priority 		string		`json:"priority"`
}

type UpdatedaToDoResponse struct {
	Id					uint		`json:"id"`
	ActivityGroupId		string		`json:"activity_group_id"`
	Title				string		`json:"title"`
	IsActive 			string		`json:"is_active"`
	Priority			string		`json:"priority"`
	CreateAt			time.Time	`json:"created_at"`
	UpdateAt 			time.Time	`json:"updated_at"`
	DeletedAt			*time.Time	`json:"deleted_at"`
}


func FindToDos(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todo []models.Todo
	id := c.Query("activity_group_id")
<<<<<<< HEAD
	// db.Clauses(hints.New("hint")).Find(&todo,id)
	db.Clauses(hints.UseIndex("idx")).Find(&todo,id)
	
=======
	
	db.Find(&todo,id)

>>>>>>> 24cf106e8429343bc23ff7b53f8f5db7a81aece3
	c.JSON(http.StatusOK, Response{
		Status:"Success",
		Message:"Success",
		Data:todo,
	})
}

func FindToDo(c *gin.Context) {
	var todo models.Todo
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	if err := db.First(&todo,id).Error;
	 
	err != nil {
        c.JSON(http.StatusNotFound, Response{
			Status: "Not Found",
			Message: "Todo with ID "+id+" Not Found",
			Data: NullResponse{},
		})
        return
    }	
    c.JSON(http.StatusOK, Response{
		Status:"Success",
		Message:"Success",
		Data:todo,
	})
}

func CreateToDo(c *gin.Context) {
	var input CreateToDoInput
	c.ShouldBindJSON(&input)

	var msg string
	
	if input.ActivityGroupId == 0 {
		msg = "activity_group_id cannot be null"
	}
	if input.Title == "" {
		msg = "title cannot be null"
	}

	if msg != "" {
		c.JSON(http.StatusBadRequest, Response{
			Status:"Bad Request",
			Message:msg,
			Data:NullResponse{},
		})
		return
	}

	todo := models.Todo{ActivityGroupId: input.ActivityGroupId ,Title: input.Title}
	  
	db := c.MustGet("db").(*gorm.DB)
	// db.Create(&todo)
	db.Create(&todo)
	c.JSON(http.StatusCreated, Response{
		Status:"Success",
		Message:"Success",
		Data:CreatedToDoResponse{
			CreatedAt: todo.CreateAt,
			UpdatedAt: todo.UpdateAt,
			Id: todo.Id,
			Title: todo.Title,
			ActivityGroupId: todo.ActivityGroupId,
			IsActive: todo.IsActive,
			Priority: todo.Priority,
		},
	})
}

func UpdateToDo(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var todo models.Todo
	id := c.Param("id")
    if err := db.First(&todo,id).Error; err != nil {
        c.JSON(http.StatusNotFound, Response{
			Status: "Not Found",
			Message: "Todo with ID "+id+" Not Found",
			Data: NullResponse{},
		})
        return
    }

    // Validate input
    var input UpdateToDoInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Todo
    updatedInput.Title = input.Title
	updatedInput.IsActive= input.IsActive
	updatedInput.UpdateAt = time.Now()

    db.Model(&todo).Updates(updatedInput)
	var btoi = func(b bool) int {
		if b {
			return 1
		}else {
			return 0
		}
	}

    c.JSON(http.StatusOK, Response{
		Status:"Success",
		Message:"Success",
		Data:UpdatedaToDoResponse{
			Id					: todo.Id,
			ActivityGroupId		: strconv.Itoa(todo.ActivityGroupId),
			Title				: todo.Title,
			IsActive 			: strconv.Itoa(btoi(input.IsActive)),
			Priority			: todo.Priority,
			CreateAt			: todo.CreateAt,
			UpdateAt 			: todo.UpdateAt,
			DeletedAt			: todo.DeletedAt,
		},
	})
}

func DeleteToDo(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var todo models.Todo
	id := c.Param("id")
    if err := db.First(&todo,id).Error; err != nil {
        c.JSON(http.StatusNotFound, Response{
			Status: "Not Found",
			Message: "Todo with ID "+id+" Not Found",
			Data: NullResponse{},
		})
        return
    }

    db.Delete(&todo)

    c.JSON(http.StatusOK, Response{
		Status: "Success",
		Message: "Success",
		Data: NullResponse{},
	})
}