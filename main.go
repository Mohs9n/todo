package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Mohs9n/todo/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	queries *db.Queries
	store = sessions.NewCookieStore([]byte("secret"))
	users = make(Users)
)

func main() {
	godotenv.Load("config/.env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}

	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatal("$DB_URL must be set")
	}

	conn, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queries = db.New(conn)

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