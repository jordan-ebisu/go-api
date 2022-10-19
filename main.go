package main

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Wash Dishes", Completed: false},
	{ID: "3", Item: "Wash Clothes", Completed: false},
}

func main() {
	// Create the server
	router := gin.Default()

	// Create the endpoint
	router.GET("/todos", getTodos)

	// Start the server
	router.Run("localhost:9090")

}

func gotTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOk, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
}
