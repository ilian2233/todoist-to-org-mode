package orgmode

import (
	"fmt"
	"strings"
	"time"
)

type Task struct {
	Title             string //TODO: Titles may contain markdown style links
	Description       string
	Date              time.Time
	Priority          rune
	RecurringDuration time.Duration
	Subtasks          []*Task
	HasParent         bool
}

func (t Task) ToString() string {
	subtasks := ""
	for _, v := range t.Subtasks {
		subtasks = fmt.Sprintf("%s*%s", subtasks, v.ToString())
	}
	return fmt.Sprintf("* TODO %s%s%s%s\n%s", t.getPriority(), t.Title, t.getDate(), t.getDescription(), subtasks) //TODO(optional): Add [%%] for percentage of completed subtask https://orgmode.org/manual/Breaking-Down-Tasks.html#Breaking-Down-Tasks-into-Subtasks
}

func (t Task) getPriority() string {
	if t.Priority == 0 {
		return ""
	}
	return fmt.Sprintf("[#%c] ", t.Priority)
}

func (t Task) getDate() string {
	if t.Date.IsZero() {
		return ""
	}
	return fmt.Sprintf("\n\t<%d-%d-%d %3s>", t.Date.Year(), t.Date.Month(), t.Date.Day(), t.Date.Weekday().String())
}

func (t Task) getDescription() string {
	if t.Description == "" {
		return ""
	}
	sanitizedDescription := strings.Replace(t.Description, "\n", ",\t", -1)
	return fmt.Sprintf("\n\t%s", sanitizedDescription)
}
