package main

import (
	"database/sql"
	"fmt"

	"github.com/ydebaere/web-service-gin/user"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
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

	router := gin.Default()

	router.POST("/users", user.AddUser)
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)

	router.Run("localhost:8080")
}
