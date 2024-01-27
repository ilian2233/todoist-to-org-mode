package main

import (
	"fmt"
	"os"
	"todoist-to-org-mode/orgmode"
	"todoist-to-org-mode/todoist"

	"github.com/spf13/pflag"
)

func main() {
	var apiKeys []string
	pflag.StringArrayVarP(&apiKeys, "key", "k", nil, "ApiKeys to be used")

	var destination string
	pflag.StringVarP(&destination, "output", "o", "", "Name of file for saving results")
	pflag.Parse()

	var projects []todoist.Project
	var tasks []todoist.Task
	for _, key := range apiKeys {
		client := todoist.Client{
			ApiKey: key,
		}
		currentProjects, err := client.GetProjects()
		if err != nil {
			panic(err)
		}
		projects = append(projects, currentProjects...)

		currentTasks, err := client.GetTasks()
		if err != nil {
			panic(err)
		}
		tasks = append(tasks, currentTasks...)
	}

	taskCreator := orgmode.TaskCreator{}
	for _, v := range projects {
		taskCreator.Add(v)
	}
	for _, v := range tasks {
		taskCreator.Add(v)
	}

	orgmodeTasks := taskCreator.CreateOrgmodeTasks()

	buf, err := orgmode.TasksToBuffer(orgmodeTasks)
	if err != nil {
		panic(err)
	}

	if destination != "" {
		if err = os.WriteFile("destination", buf.Bytes(), 0666); err != nil {
			panic(err)
		}
	} else {
		fmt.Print(buf.String())
	}
}
