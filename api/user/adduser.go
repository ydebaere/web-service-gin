package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	statement, _ := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	statement.Exec(newUser.Name, newUser.Email)
	c.JSON(http.StatusOK, gin.H{"status": "user added"})
}

func GetUsers(c *gin.Context) {
	rows, _ := db.Query("SELECT * FROM users")
	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name, &user.Email)
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	var user User
	row.Scan(&user.ID, &user.Name, &user.Email)
	c.JSON(http.StatusOK, user)
}
