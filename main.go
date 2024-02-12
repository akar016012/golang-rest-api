package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* define data struct */
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

/* define dummy data */
var todos = []todo{
	{
		ID:        "1",
		Item:      "Clean Room",
		Completed: false,
	},
	{
		ID:        "2",
		Item:      "Read Book",
		Completed: false,
	},
	{
		ID:        "3",
		Item:      "Record Video",
		Completed: false,
	},
}

/* convert data to JSON */
func getTodos(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK,todos)
}

/* create another endpoint for post */

func addTodo(c *gin.Context){
	var newTodo todo

	if err:= c.BindJSON(&newTodo); err!=nil {
		return
	}
	todos = append(todos, newTodo)

	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string)(*todo, error){
	for i, t:=range todos{
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(c *gin.Context){
	id := c.Param("id")
	todo, err:= getTodoById(id)
	if err !=nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Todo not found"})
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(c *gin.Context){
	id := c.Param("id")
	todo, err:= getTodoById(id)

	if err !=nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Todo not found"})
	}

	todo.Completed =!todo.Completed
	c.IndentedJSON(http.StatusOK, todo)

}

/* create server and router */
func main(){
	router := gin.Default()
	router.GET("/todos",getTodos)
	router.GET("/todos/:id",getTodo)
	router.PATCH("/todos/:id",toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
