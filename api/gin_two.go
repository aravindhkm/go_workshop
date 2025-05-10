package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-ozzo/ozzo-validation/v4"
)

func Pong(ctx *gin.Context) {

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"data": "pong-",
		},
	)
}

type Payload struct {
	Name string `json:"name"`
}

func (a Payload) Validate() error {
	return validator.ValidateStruct(&a,
		validator.Field(a.Name, validator.Required),
	)
}

func Create(ctx *gin.Context) {
	var payload Payload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		fmt.Println("err", err)
	}
	err := payload.Validate()

	if err != nil {
		fmt.Println("validate err", err)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			map[string]any{
				"message": "invalid data type",
			},
		)

		return
	}
	value := map[string]any{
		"message": "created successfully",
		"data":    payload,
	}

	ctx.JSON(
		http.StatusCreated,
		value,
	)
}

func GinExampleTwo() {
	engine := gin.Default()

	engine.GET("/ping", Pong)

	engine.POST("/create", Create)

	engine.Run(":4000")
}
