package api

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Number   int    `json:"number" binding:"required"`
	Position string `json:"position" binding:"required"`
}

var (
	filePath = "users.json"
	mu       sync.Mutex
)

func createUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	users := loadUsers()

	users = append(users, user)

	saveUsers(users)

	c.JSON(http.StatusOK, gin.H{"status": "user created", "user": user})
}

func loadUsers() []User {
	var users []User

	data, err := os.ReadFile(filePath)
	if err != nil {
		// File doesn't exist or is empty
		return users
	}

	_ = json.Unmarshal(data, &users)
	return users
}

func saveUsers(users []User) {
	data, _ := json.MarshalIndent(users, "", "  ")
	_ = os.WriteFile(filePath, data, 0644)
}

func GinExampleThree() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ping"})
	})

	router.POST("/user", createUser)

	router.Run(":8080")
}

// curl -X POST http://localhost:8080/user \
//   -H "Content-Type: application/json" \
//   -d '{
//     "name": "John",
//     "last_name": "Doe",
//     "number": 123,
//     "position": "Developer"
//   }'
