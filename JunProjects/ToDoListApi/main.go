package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

type Task struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Status   string  `json:"status"`
}

var db *sql.DB

func main () {
    var err error

	db, err = sql.Open("sqlite", "./tasks.db")

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTask)
	router.POST("/task", createTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	router.Run(":8080")
}

func createTable() {
	createTableSQL := `
	CREATE TABLE IS NOT EXISTS tasks (
	"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"title" TEXT,
	"status" TEXT
	);
	`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
}

func getTasks(context *gin.Context) {
	rows, err := db.Query("SELECT id, title, status FROM tasks")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Query is not correct!"})
		return
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task

		err := rows.Scan(&task.ID, &task.Title, &task.Status)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error":"could send each task"})
		}
		tasks = append(tasks, task)
	}
}

func getTask(context *gin.Context) {
	id := context.Param("id")

	var task Task
	err := db.QueryRow("SELECT id, title, status FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Title, &task.Status)
    
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"message":"Task row not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "row not available"})
		}
		return
	}

	context.JSON(http.StatusOK, task)

}

func createTask(context *gin.Context) {
	var task Task
	
	err := context.ShouldBindJSON(&task)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request creating task"})
		return
	}

	insertTaskSQL := `INSERT INTO tasks(title, status) VALUES (?, ?)`
	result, err := db.Exec(insertTaskSQL, task.Title, task.Status)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Can't create in SQL"})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error id"})
		return
	}
	task.ID = int(id)

	context.JSON(http.StatusOK, task)

}

func updateTask(context *gin.Context) {
	id := context.Param("id")

	var task Task

	err := context.ShouldBindJSON(&task)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request creating task"})
		return
	}
	
	updateTaskSQL := `UPDATE tasks SET title = ?, status = ? WHERE id = ?`
	_, err = db.Exec(updateTaskSQL, task.Title, task.Status, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not updated in sqlite"})
		return	
	}

	task.ID, _ = strconv.Atoi(id)
    context.JSON(http.StatusOK, task)

}

func deleteTask(context *gin.Context) {
    id := context.Param("id")

	deleteTaskSql := `DELETE FROM tasks WHERE id = ?`
	_, err := db.Exec(deleteTaskSql, id)
if err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete in sqlite"})
	return
}
context.JSON(http.StatusOK, gin.H{"message": "Task deleted"})

}