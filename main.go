package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID       string `json:"id"`
	Item     string `json:"item"`
	Complete bool   `json:"complete"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Complete: false},
	{ID: "2", Item: "Read Book", Complete: false},
	{ID: "3", Item: "Record Video", Complete: false},
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
