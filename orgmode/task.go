package orgmode

import "fmt"

type Task struct {
	Title       string //TODO: Titles may contain markdown style links
	Description string //TODO: Add to ToString
	Date        string //TODO: Add to ToString
	Priority    rune   //TODO: Add to ToString
	IsRecurring bool
	Subtasks    []*Task
	HasParent   bool
}

func (t Task) ToString() string {
	subtasks := ""
	for _, v := range t.Subtasks {
		subtasks = fmt.Sprintf("%s*%s", subtasks, v.ToString())
	}
	return fmt.Sprintf("* TODO %s\n%s", t.Title, subtasks) //TODO(optional): Add [%%] for percentage of completed subtask https://orgmode.org/manual/Breaking-Down-Tasks.html#Breaking-Down-Tasks-into-Subtasks
}
