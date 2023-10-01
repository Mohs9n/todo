package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var (
	users = make(Users)
	store = sessions.NewCookieStore([]byte("secret"))
) 

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.StaticFile("/favicon.ico", "static/favicon.ico")
	r.StaticFile("/output.css", "templates/output.css")
	r.GET("/", index)
	r.POST("/login", login)
	r.POST("/signup", signup)
	r.POST("/addTodo", addTodoItem)
	r.POST("/toggleTodo", toggleTodoItem)
	r.Run(":9091")
}