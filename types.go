package main

import "github.com/Mohs9n/todo/internal/db"

type User struct {
	ID uint64
	Name string
	TodoList TodoList
}

type Todo struct {
	ID uint64
	Text string
	Done bool
}

type TodoList []db.GetTodosRow

type Users map[string]User