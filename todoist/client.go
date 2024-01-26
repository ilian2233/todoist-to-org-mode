package todoist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	ApiKey string
}

func (c Client) GetProjects() ([]Project, error) {
	body, err := makeTodoistRequest("projects", c.ApiKey)
	if err != nil {
		return nil, err
	}

	var projects []Project
	if err = json.Unmarshal(body, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (c Client) GetTasks() ([]Task, error) {
	body, err := makeTodoistRequest("tasks", c.ApiKey)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err = json.Unmarshal(body, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func makeTodoistRequest(object, apiKey string) ([]byte, error) {
	url := fmt.Sprintf("https://api.todoist.com/rest/v2/%s", object)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			panic(err)
		}
	}(res.Body)

	return io.ReadAll(res.Body)
}
