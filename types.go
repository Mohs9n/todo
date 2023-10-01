package main

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

type TodoList []Todo

type Users map[string]User