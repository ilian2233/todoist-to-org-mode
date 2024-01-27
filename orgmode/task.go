package orgmode

import (
	"fmt"
	"time"
)

type Task struct {
	Title             string    //TODO: Titles may contain markdown style links
	Description       string    //TODO: Add to ToString
	Date              time.Time //TODO: Add to ToString
	Priority          rune      //TODO: Add to ToString
	RecurringDuration time.Duration
	Subtasks          []*Task
	HasParent         bool
}

func (t Task) ToString() string {
	subtasks := ""
	for _, v := range t.Subtasks {
		subtasks = fmt.Sprintf("%s*%s", subtasks, v.ToString())
	}
	return fmt.Sprintf("* TODO %s\n%s", t.Title, subtasks) //TODO(optional): Add [%%] for percentage of completed subtask https://orgmode.org/manual/Breaking-Down-Tasks.html#Breaking-Down-Tasks-into-Subtasks
}
