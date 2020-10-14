package api

import (
	"github.com/allentom/youcrawl"
	"github.com/gin-gonic/gin"
)

var PingHandler gin.HandlerFunc = func(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})

}

type AddUrlRequestBody struct {
	Urls []string `json:"urls"`
}

var AddUrlHandler gin.HandlerFunc = func(context *gin.Context) {
	requestBody := AddUrlRequestBody{}
	err := context.BindJSON(&requestBody)
	if err != nil {
		context.AbortWithError(400, err)
		return
	}
	engine.Pool.AddURLs(requestBody.Urls...)
	context.JSON(200, gin.H{"result": true})
}

type TaskItem struct {
	Item map[string]interface{} `json:"item"`
	Url  string                 `json:"url"`
}
type AddTaskItemRequestBody struct {
	Tasks []TaskItem `json:"tasks"`
}

var AddTaskHandler gin.HandlerFunc = func(context *gin.Context) {
	requestBody := AddTaskItemRequestBody{}
	err := context.BindJSON(&requestBody)
	if err != nil {
		context.AbortWithError(400, err)
		return
	}

	tasks := make([]*youcrawl.Task, 0)
	for _, taskItem := range requestBody.Tasks {
		item, err := taskItemSerializer.Unmarshal(taskItem.Item)
		if err != nil {
			context.AbortWithError(400, err)
			return
		}
		tasks = append(tasks, &youcrawl.Task{
			Url: taskItem.Url,
			Context: youcrawl.Context{
				Item: item,
			},
		})
	}
	engine.Pool.AddTasks(tasks...)
}
