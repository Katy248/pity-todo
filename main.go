package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskListType int

const (
	AllTasks TaskListType = iota
	CompletedTasks
	UncompletedTasks
)

func main() {
	if _, err := load(); err != nil {
		panic(err)
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", ListTasks(AllTasks))
	router.GET("/completed", ListTasks(CompletedTasks))
	router.GET("/uncompleted", ListTasks(UncompletedTasks))
	router.POST("/delete/:id", DeleteTaskHandler)
	router.POST("/complete/:id", CompleteTaskHandler)
	router.POST("/add", AddTask)

	if err := router.Run(":5900"); err != nil {
		panic(err)
	}
}
func ListTasks(t TaskListType) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := GetTasks(t)
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Tasks": result,
		})
	}
}

func IndexView(ctx *gin.Context, err error) {
	tasks := GetTasks(AllTasks)
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Tasks": tasks,
		"Error": err,
	})
}

func DeleteTaskHandler(ctx *gin.Context) {

	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		IndexView(ctx, err)
		return
	}
	if err := DeleteTask(id); err != nil {
		IndexView(ctx, err)
		return
	}
	IndexView(ctx, nil)
}
func CompleteTaskHandler(ctx *gin.Context) {

	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		IndexView(ctx, err)
		return
	}
	if err := CompleteTask(id); err != nil {
		IndexView(ctx, err)
		return
	}
	IndexView(ctx, nil)
}

func AddTask(ctx *gin.Context) {
	description := ctx.PostForm("description")
	if description == "" {
		IndexView(ctx, fmt.Errorf("description is required"))
		return
	}
	if err := NewTask(description); err != nil {
		IndexView(ctx, err)
		return
	}
	IndexView(ctx, nil)
}
