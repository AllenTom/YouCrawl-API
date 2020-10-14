package api

import (
	"github.com/allentom/youcrawl"
	"testing"
)

func TestDashboardAPI_Run(t *testing.T) {
	e :=  youcrawl.NewEngine(&youcrawl.EngineOption{
		MaxRequest: 3,
		Daemon:     true,
	})
	//e.AddURLs("http://www.example.com")
	api := DashboardAPI{}
	e.AddPlugins(&api)
	e.RunAndWait()
}