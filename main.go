package main

import (
	"net/http"
  "errors"
	"github.com/gin-gonic/gin"
)

type todo struct { // adding struct
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{ // craeting array for todos to insert values
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Reading books", Completed: false},
	{ID: "3", Item: "Record Videos", Completed: false},
	// this values are not json types in order to retunrn values to client we have to make it json values, so wae go to the struct and add `json: " name(we want to show the client)"`
}

func getTodos(context *gin.Context) { // fn that will convert the above values to json format
	context.IndentedJSON(http.StatusOK, todos)

}

func addTodo(context *gin.Context){  //creating new todo for POST
	var newTodo todo // creating new variable name newTodo with type=todo
  
	if err :=context.BindJSON(&newTodo); err!= nil{ //we are creating a new var and using BindJSON(&newTodo).Bcz,it will take values from client and make it in json format to understandable to go lang // and if the newTodo doesnt hvae the struct format above it will show error, if therer's no error it will return.
		return
	}
  
	
todos = append(todos, newTodo)// taking the todos array & adding the newTodo into it.

context.IndentedJSON(http.StatusCreated, newTodo) // we will do it as before & and this time we will add StatusCreated bcz it has alrady created before, and add newTodo with it..
// it will convert the new todo to json and return.. 
}


// Now we are going to create handellars that will control the getTodoByid()
func getTodo(context *gin.Context){      //handdelar and its type is context as before, it will get the id 
	id := context.Param("id")              //it will bring the, to implemet we need to add new line in main function 	router.GET("todos/id", getTodo)..
	todo, err :=getTodoById(id)            // it will give a todo or error by givern id..

	if err != nil { // if there is an error it will give
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"}) // status not found status with a coustom message..
		return
	}
	context.IndentedJSON(http.StatusOK, todo) // if there is no error it will return the todo
 }


 func toggleTodoStatus(context *gin.Context){    // for changing boolean values,this fn will reverse the value
  
	id := context.Param("id")             
	todo, err :=getTodoById(id)

	if err != nil {             // if there is an error it will give
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"}) // status not found status with a coustom message..
		return
	}

	todo.Completed = !todo.Completed // it will reverse the bool, if true give false,if false give true
  
	context.IndentedJSON(http.StatusOK, todo)

 }




 func getTodoById(id string) (*todo, error){  //here we are getting todo by mentiong only id with localhost:9090 address,and it will gaive us two valu(*todo, error)a todo or a error, if give todo error is nil,or give error todo is nil..
	for i, t := range todos { // i=id & t=todos.id, 
		if t.ID == id { // if todos id == our given id here ,
		 return &todos[i], nil // then it will return a todos and error will be nil
		}
	}
	return nil, errors.New("todo not found") // and if we don't found any todos, then it will be goin to nil, and error.New() will show us the not found massage, and we have to imoport "error" package..
	//after that go to postman and search with id like""localhost:todo/3" and it will give us the todo we want..
 }


func main() {
	router := gin.Default()        // craeting routern var which goin to act like defalult fn.
	router.GET("/todos", getTodos) // adding endpoints of our server, we need a function now who will get values from /todos in future,so we are creating a fn before our main fn..after that we cmae back to fn and call it..
	router.GET("/todos/:id", getTodo) // we can get the specific todo using this, and the function is getTodo which will give us the specific id todo..
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo) // it will post the new added todo into it..
	router.Run("localhost:9090")   // adding the server localhost number
}
