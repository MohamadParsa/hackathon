package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB

type QuickAccess struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Title  string `json:"title"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/items", getAllQuickAccess)
	r.POST("/items", createQuickAccess)
	r.GET("/items/:id", getQuickAccess)
	r.PUT("/items/:id", updateQuickAccess)
	r.DELETE("/items/:id", deleteQuickAccess)

	return r
}

func initDB() {
	var err error
	dsn := "host=localhost user=your_user password=your_password dbname=your_database port=5432 sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&QuickAccess{})
}

func getAllQuickAccess(c *gin.Context) {
	var items []QuickAccess
	db.Find(&items)

	c.JSON(http.StatusOK, items)
}

func getQuickAccess(c *gin.Context) {
	id := c.Param("id")
	var item QuickAccess

	if err := db.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "QuickAccess not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func createQuickAccess(c *gin.Context) {
	var newQuickAccess QuickAccess
	if err := c.ShouldBindJSON(&newQuickAccess); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&newQuickAccess)
	c.JSON(http.StatusCreated, newQuickAccess)
}

func updateQuickAccess(c *gin.Context) {
	id := c.Param("id")
	var item QuickAccess

	if err := db.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "QuickAccess not found"})
		return
	}

	var updatedItem QuickAccess
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&item).Updates(updatedItem)
	c.JSON(http.StatusOK, item)
}

func deleteQuickAccess(c *gin.Context) {
	id := c.Param("id")
	var item QuickAccess

	if err := db.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "QuickAccess not found"})
		return
	}

	db.Delete(&item)
	c.JSON(http.StatusNoContent, nil)
}

func main() {
	initDB()

	r := setupRouter()
	fmt.Println("Server is running on :8080")
	log.Fatal(r.Run(":8080"))
}
