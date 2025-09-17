package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskListType int

const (
	AllTasks TaskListType = iota
	CompletedTask
	UncompletedTasks
)

func main() {
	load()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", ListTasks(AllTasks))
	router.GET("/completed", ListTasks(CompletedTask))
	router.GET("/uncompleted", ListTasks(UncompletedTasks))
	router.POST("/delete/:id", DeleteTask)
	router.POST("/complete/:id", CompleteTask)
	router.POST("/add", AddTask)

	router.Run(":5900")
}
func ListTasks(t TaskListType) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var result []*Task

		switch t {
		case CompletedTask:
			for _, task := range tasks {
				if task.Completed {
					result = append(result, task)
				}
			}
		case UncompletedTasks:
			for _, task := range tasks {
				if !task.Completed {
					result = append(result, task)
				}
			}
		default:
			result = tasks
		}

		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Tasks": result,
		})
	}
}

func DeleteTask(ctx *gin.Context) {

	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Redirect(http.StatusFound, "/")
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			save()
			ctx.Redirect(http.StatusFound, "/")
			return
		}
	}
	ctx.Redirect(http.StatusOK, "/")
}
func CompleteTask(ctx *gin.Context) {

	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Redirect(http.StatusFound, "/")
		return
	}
	for _, task := range tasks {
		if task.ID == id {
			task.Completed = true
			save()
			ctx.Redirect(http.StatusFound, "/")
			return
		}
	}
	ctx.Redirect(http.StatusOK, "/")
}

type Task struct {
	ID          int
	Description string
	Completed   bool
}

func AddTask(ctx *gin.Context) {
	description := ctx.PostForm("description")
	if description == "" {
		ctx.Redirect(http.StatusFound, "/")
		return
	}
	NewTask(description)
	ctx.Redirect(http.StatusFound, "/")
}

func NewTask(description string) {
	task := &Task{ID: len(tasks) + 1, Description: description, Completed: false}
	tasks = append(tasks, task)
	save()
}

var tasks = []*Task{}

func GetTasks() []*Task {
	load()
	return tasks
}

func load() {
	data, err := os.ReadFile("tasks.json")
	json.Unmarshal(data, &tasks)
	if err != nil {
		panic(err)
	}
}
func save() {
	data, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}
	os.WriteFile("tasks.json", data, 0644)
}
