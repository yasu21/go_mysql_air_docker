package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	UserID    uint      `gorm:"primary_key;AUTO_INCREMENT"`
	UserName  string    `gorm:"type:varchar(36);NOT NULL"`
	CreatedAt time.Time `gorm:"type:datetime"`
}

var db *gorm.DB
var err error

func init() {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	HOST := os.Getenv("MYSQL_HOST")
	PROTOCOL := "tcp(" + HOST + ":3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")
	QUERY := "charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	connect := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + QUERY
	db, err = gorm.Open("mysql", connect)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&User{})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		var users []User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(500, gin.H{"error": "Error fetching users from database"})
			return
		}
		fmt.Println(users)
		c.JSON(200, users)
	})
	defer db.Close()
	r.Run()
}
