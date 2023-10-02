package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Mohs9n/todo/internal/db"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	if session, _ := store.Get(c.Request, "logged_in"); session.Values["username"] != nil {
		fmt.Printf("\n%v\n", session.Values["username"])
		if usr, err := queries.GetUserByName(c, session.Values["username"].(string)); err == nil {
			todos, err := queries.GetTodos(c, usr.ID)
			if err != nil {
				fmt.Printf("\n%v\n", err)
			}
			user := User{
				ID:   uint64(usr.ID),
				Name: usr.Name,
				TodoList: todos,
			}
			c.HTML(http.StatusOK, "loggedIn.html", user)
		}
	} else {
		c.HTML(http.StatusOK, "index.html", nil)
	}
}

func login(c *gin.Context) {
	// db
	if exists, _ :=queries.CheckUserExists(c, c.PostForm("username")); exists {
		usr, err :=queries.GetUserByName(c, c.PostForm("username"))
		if err != nil {
			fmt.Printf("\n%v\n", err)
		}
		fmt.Printf("\n%v\n", usr)

		todos, err := queries.GetTodos(c, usr.ID)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		}
		user := User{
			ID:   uint64(usr.ID),
			Name: usr.Name,
			TodoList: todos,
		}

		// session
		session, _ := store.Get(c.Request, "logged_in")
		session.Values["username"] = usr.Name

		session.Save(c.Request, c.Writer)
		fmt.Printf("\nafter login coockie: %v\n", session.Values["username"])

		c.HTML(http.StatusOK, "loggedIn.html", user)
	}
}

func signup(c *gin.Context) {
	// db
	if exists, _ :=queries.CheckUserExists(c, c.PostForm("username")); !exists {
		usr, err :=queries.CreateUser(c, c.PostForm("username"))
		if err != nil {
			fmt.Printf("\n%v\n", err)
		}
		fmt.Printf("\n%v\n", usr)

		todos, err := queries.GetTodos(c, usr.ID)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		}
		user := User{
			ID:   uint64(usr.ID),
			Name: usr.Name,
			TodoList: todos,
		}

		// session
		session, _ := store.Get(c.Request, "logged_in")
		session.Values["username"] = usr.Name

		fmt.Printf("\nafter signup cookie: %v\n", session.Values["username"])
		session.Save(c.Request, c.Writer)

		c.HTML(http.StatusOK, "loggedIn.html", user)
	}
}

func addTodoItem(c *gin.Context) {
	// db
	usr, err := queries.GetUserByName(c, c.PostForm("username"))
	if err != nil {
		fmt.Printf("\n%v\n", err)
	}
	queries.CreateTodo(c, db.CreateTodoParams{
		Text:   c.PostForm("todoText"),
		Done:   false,
		UserID: int32(usr.ID),
	})
	todos, err := queries.GetTodos(c, int32(usr.ID))
	if err != nil {
		fmt.Printf("\n%v\n", err)
	}
	user:= User{
		ID:   uint64(usr.ID),
		Name: c.PostForm("username"),
		TodoList: todos,
	}
	c.HTML(http.StatusOK, "loggedIn.html", user)
}

func toggleTodoItem(c *gin.Context) {
	// db
	todoID, err := strconv.ParseInt(c.PostForm("todoID"), 10, 32)
	if err != nil {
		fmt.Printf("\n%v\n", err)
	}
	queries.UpdateTodo(c, int32(todoID))
	usr, err := queries.GetUserByName(c, c.PostForm("username"))
	if err != nil {
		fmt.Printf("\n%v\n", err)
	}
	todos, err := queries.GetTodos(c, usr.ID)
	if err != nil {
		fmt.Printf("\n%v\n", err)
	}
	user := User{
		ID:   uint64(usr.ID),
		Name: usr.Name,
		TodoList: todos,
	}
	c.HTML(http.StatusOK, "loggedIn.html", user)
}