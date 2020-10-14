package api

import (
	"github.com/allentom/youcrawl"
	"github.com/gin-gonic/gin"
)

var engine *youcrawl.Engine
var taskItemSerializer ItemSerializer
type DashboardAPI struct {
	TaskItemSerializer ItemSerializer
}

func (p *DashboardAPI) Run(e *youcrawl.Engine) {
	engine = e
	taskItemSerializer = p.TaskItemSerializer
	r := gin.Default()
	r.GET("/ping", PingHandler)
	r.POST("/pool/urls", AddUrlHandler)
	r.POST("/pool/tasks", AddTaskHandler)
	r.Run()
}
