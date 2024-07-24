package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

// Task represents a to-do item
type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite", "./tasks.db")
	if err != nil {
		panic(err)
	}

	// Create the tasks table if it doesn't exist
	createTable()

	r := gin.Default()

	r.GET("/tasks", getTasks)
	r.GET("/tasks/:id", getTask)
	r.POST("/tasks", createTask)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	r.Run(":8080")
}

func createTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"status" TEXT
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
}

// getTasks handles GET requests to /tasks
func getTasks(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, status FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

// getTask handles GET requests to /tasks/:id
func getTask(c *gin.Context) {
	id := c.Param("id")
	var task Task
	err := db.QueryRow("SELECT id, title, status FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Title, &task.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, task)
}

// createTask handles POST requests to /tasks
func createTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertTaskSQL := `INSERT INTO tasks(title, status) VALUES (?, ?)`
	result, err := db.Exec(insertTaskSQL, task.Title, task.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	task.ID = int(id)

	c.JSON(http.StatusOK, task)
}

// updateTask handles PUT requests to /tasks/:id
func updateTask(c *gin.Context) {
	id := c.Param("id")
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTaskSQL := `UPDATE tasks SET title = ?, status = ? WHERE id = ?`
	_, err := db.Exec(updateTaskSQL, task.Title, task.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task.ID, _ = strconv.Atoi(id)
	c.JSON(http.StatusOK, task)
}

// deleteTask handles DELETE requests to /tasks/:id
func deleteTask(c *gin.Context) {
	id := c.Param("id")

	deleteTaskSQL := `DELETE FROM tasks WHERE id = ?`
	_, err := db.Exec(deleteTaskSQL, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
