package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type UserData struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Number   int    `json:"number"`
	Position string `json:"position"`
}

var (
	localFilePath = "./data.json"
	mutexLock       sync.Mutex
)

func loadLocalUsers() []UserData {
	dataByte, err := os.ReadFile(localFilePath)

	if err != nil {
		log.Fatal("Unable to read local file")
	}

	var users []UserData
	err = json.Unmarshal(dataByte, &users)
	if err != nil {
		log.Fatal("Unable to json Unmarshal")
	}

	return users
}

func saveLocalUsers(newUser UserData) error {
	existUsers := loadLocalUsers()

	existUsers = append(existUsers, newUser)

	userByte, err := json.MarshalIndent(existUsers, "", "  ")
	if err != nil {
		return errors.New("unable to json marshal")
	}
	err = os.WriteFile(localFilePath, userByte, 0644)
	if err != nil {
		return errors.New("unable to save file")
	}

	return nil
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	var user UserData

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Basic validation
	if user.Name == "" || user.LastName == "" || user.Number == 0 || user.Position == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	mutexLock.Lock()
	defer mutexLock.Unlock()

	saveLocalUsers(user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "user created",
		"user":   user,
	})
}

func HttpApiOneTwo() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"message": "pong"})
	})

	mux.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"message": "ping"})
	})

	mux.HandleFunc("/user", CreateUsers)

	http.ListenAndServe(":8080", mux)

	fmt.Println("http starts")
}

// curl -X POST http://localhost:8080/user \
//   -H "Content-Type: application/json" \
//   -d '{
//     "name": "John",
//     "last_name": "Doe",
//     "number": 123,
//     "position": "Developer"
//   }'

// curl http://localhost:8080/ping
