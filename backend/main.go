package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	Title       string `gorm:"unique;not null" json:"title"`
	Description string `gorm:"size:500" json:"description"`
	Priority    int8   `json:"priority"`
}

type taskList struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int8   `json:"priority"`
}

func getTasks(c *gin.Context) {
	var tasks []taskList

	DB.Model(&Tasks{}).Scan(&tasks)

	c.JSON(http.StatusOK, tasks)
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

	dsn := "root:1234@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

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
	r.GET("/tasks", getTasks)
	r.POST("/tasks", postTask)

	r.Run(":8080")
}
