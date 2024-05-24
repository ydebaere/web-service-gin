package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ydebaere/web-service-gin/api/user"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err)
	}

	if !checkUsersTableExists() {
		_, err := db.Exec(`CREATE TABLE users (
            id INTEGER PRIMARY KEY,
            name TEXT,
            email TEXT
        )`)
		if err != nil {
			log.Fatal(err)
		}
	}

	router := gin.Default()

	router.StaticFile("/", "./public/index.html")

	router.POST("/users", user.AddUser)
	router.GET("/users", user.GetUsers)
	router.GET("/users/:id", user.GetUser)

	router.Run("localhost:8080")
}

func checkUsersTableExists() bool {
	_, err := db.Query("SELECT 1 FROM users LIMIT 1")
	if err != nil {
		return false
	}
	return true
}
