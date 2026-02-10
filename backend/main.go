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

type Tasks struct {
	gorm.Model
	Title       string `gorm:"unique;not null" json:"title"`
	Description string `gorm:"size:500" json:"description"`
	Priority    int8   `json:"priority"`
}

type taskList struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int8   `json:"priority"`
}

func getTasks(c *gin.Context) {
	var tasks []taskList

	DB.Model(&Tasks{}).Scan(&tasks)

	c.JSON(http.StatusOK, tasks)
}

func deleteTaskByID(c *gin.Context) {
	id := c.Param("id")
	DB.Unscoped().Delete(&Tasks{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func updateTaskByID(c *gin.Context) {
	id := c.Param("id")
	var updateTask Tasks
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Model(&Tasks{}).Where("id = ?", id).Update("priority", updateTask)
}

func postTask(c *gin.Context) {
	var newTask Tasks
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := DB.Create(&newTask)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTask)

}

var DB *gorm.DB

func main() {

	dsn := "root:1234@tcp(localhost:3305)/task_manager?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection established")

	err = DB.AutoMigrate(&Tasks{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database migration complete")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		MaxAge:       12 * time.Hour,
	}))

	r.GET("/api/tasks", getTasks)
	r.POST("/tasks", postTask)
	r.PATCH("/tasks/:id", updateTaskByID)
	r.DELETE("/tasks/:id", deleteTaskByID)

	r.Run(":8080")
}
