package main

import (
	"fmt"
	"os"
	"todoist-to-org-mode/orgmode"
	"todoist-to-org-mode/todoist"
)

func main() {
	apiKey := os.Getenv("TODOIST_API")

	todoistClient := todoist.Client{
		ApiKey: apiKey,
	}

	todoistProjects, err := todoistClient.GetProjects()
	if err != nil {
		panic(err)
	}

	todoistTasks, err := todoistClient.GetTasks()
	if err != nil {
		panic(err)
	}

	stt := orgmode.TaskCreator{}
	for _, v := range todoistProjects {
		stt.Add(v)
	}
	for _, v := range todoistTasks {
		stt.Add(v)
	}

	orgmodeTasks := stt.CreateOrgmodeTasks()

	//TODO: Replace with saving to file
	for _, v := range orgmodeTasks {
		fmt.Print(v.ToString())
	}
}
