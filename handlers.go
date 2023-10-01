package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	if session, _ := store.Get(c.Request, "logged_in"); session.Values["username"] != nil {
		if usr, ok := users[session.Values["username"].(string)]; ok {
			c.HTML(http.StatusOK, "loggedIn.html", usr)
		}
	} else {
		c.HTML(http.StatusOK, "index.html", nil)
	}
}

func login(c *gin.Context) {
	if usr, ok := users[c.PostForm("username")]; ok {
		session, _ := store.Get(c.Request, "logged_in")
		session.Values["username"] = usr.Name

		session.Save(c.Request, c.Writer)
		fmt.Printf("\n%v\n", session.Values["username"])
		c.HTML(http.StatusOK, "loggedIn.html", usr)
	}
}

func signup(c *gin.Context) {
	if usrname := c.PostForm("username"); usrname != "" {
		if _, ok := users[usrname]; !ok {
			session, _ := store.Get(c.Request, "logged_in")
			session.Values["username"] = users[usrname].Name
	
			session.Save(c.Request, c.Writer)
			users[usrname] = User{
				ID:   uint64(len(users)+1),
				Name: usrname,
				TodoList: TodoList{},
			}
			c.HTML(http.StatusOK, "loggedIn.html", users[usrname])
		}
	}
}

func addTodoItem(c *gin.Context) {
	fmt.Println(c.PostForm("username"))
	
	// Retrieve the user from the map
	user, ok := users[c.PostForm("username")]
	if !ok {
		// Handle the case where the user doesn't exist
		c.AbortWithStatus(http.StatusNotFound)
		return
	}	

	// Modify the TodoList field of the retrieved user
	user.TodoList = append(user.TodoList, Todo{
		ID:   uint64(len(user.TodoList) + 1),
		Text: c.PostForm("todoText"),
		Done: false,
	})

	users[c.PostForm("username")] = user
	
	fmt.Println(user.TodoList)
	c.HTML(http.StatusOK, "loggedIn.html", user)
}

func toggleTodoItem(c *gin.Context) {
	fmt.Printf("\nindex: %v\n", c.PostForm("todoIdx"))
	user, ok := users[c.PostForm("username")]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	fmt.Printf("\nbefore: %v\n", user.TodoList)
	idx, _ := strconv.ParseInt(c.PostForm("todoIdx"), 10, 64)

	user.TodoList[idx].Done = !user.TodoList[idx].Done
	fmt.Printf("\nconv index: %v\n", idx)
	users[c.PostForm("username")] = user
	fmt.Printf("\nafter: %v\n", users[c.PostForm("username")].TodoList)
	c.HTML(http.StatusOK, "loggedIn.html", user)
}