package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"time"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"unique;not null; size:100" json:"title"`
	Description string `gorm:"size:500" json:"description"`
	Priority    int8   `json:"priority"`
}

type taskList struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int8   `json:"priority"`
	CreatedAt   string `json:"created_at"`
}

type TaskHandler struct {
	DB *gorm.DB
}

func (h *TaskHandler) getTasks(c *gin.Context) {
	var tasks []taskList

	if err := h.DB.Model(&Task{}).Scan(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) deleteTaskByID(c *gin.Context) {
	id := c.Param("id")
	result := h.DB.Unscoped().Delete(&Task{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func (h *TaskHandler) updateTaskByID(c *gin.Context) {
	id := c.Param("id")
	var updateTask struct {
		Priority int8 `json:"priority"`
	}

	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.DB.Model(&Task{}).Where("id = ?", id).Update("priority", updateTask.Priority)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func (h *TaskHandler) postTask(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.DB.Create(&newTask)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTask)

}

func main() {

	dsn := "root:1234@tcp(localhost:3305)/task_manager?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Database connection established")

	err = db.AutoMigrate(&Task{})
	if err != nil {
		panic("Database migration failed")
	}
	fmt.Println("Database migration complete")

	taskHandler := &TaskHandler{DB: db}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		MaxAge:       12 * time.Hour,
	}))

	r.GET("/api/tasks", taskHandler.getTasks)
	r.POST("/api/tasks", taskHandler.postTask)
	r.PATCH("/api/tasks/:id", taskHandler.updateTaskByID)
	r.DELETE("/api/tasks/:id", taskHandler.deleteTaskByID)

	r.Run(":8080")
}
