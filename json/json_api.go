package json

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type returnValue map[string]any

type Data struct {
	FileName  string    `json:"fileName"`
	Date      time.Time `json:"date"`
	IpAddress string    `json:"ipAddress"`
}

const dataFile = "data.json"

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "pong"})
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, returnValue{"data": "ping"})
}

func GetData(c *gin.Context) {
	result, err := readData()
	if err != nil {
		c.JSON(http.StatusNotFound, returnValue{"data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, returnValue{
		"totalLength": len(result),
		"data":        result,
	})
}

func WriteData(c *gin.Context) {
	result, err := readData()
	if err != nil && !os.IsNotExist(err) {
		c.JSON(http.StatusInternalServerError, returnValue{"data": err.Error()})
		return
	}

	newEntry := Data{
		FileName:  fmt.Sprintf("File %v", len(result)+1),
		Date:      time.Now(),
		IpAddress: getIPAddress(),
	}

	result = append(result, newEntry)

	if err := writeData(result); err != nil {
		c.JSON(http.StatusInternalServerError, returnValue{"data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, returnValue{"data": result})
}

func readData() ([]Data, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []Data{}, nil
	}

	var result []Data
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}
	return result, nil
}

func writeData(data []Data) error {
	byteData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	return os.WriteFile(dataFile, byteData, 0644)
}

func getIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip := ipnet.IP.To4(); ip != nil {
				return ip.String()
			}
		}
	}
	return ""
}

func JsonApi() {
	app := gin.Default()

	app.GET("/ping", Ping)
	app.GET("/pong", Pong)
	app.POST("/create", WriteData)
	app.GET("/get", GetData)

	log.Println("Server running on http://localhost:8080")
	if err := app.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}

// curl -X POST http://localhost:8080/create
