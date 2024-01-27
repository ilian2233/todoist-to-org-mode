package orgmode

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Task struct {
	Title             string
	Description       string
	Date              time.Time
	Priority          rune
	RecurringDuration time.Duration //TODO: Add recurring tasks functionality to date to date
	Subtasks          []*Task
	HasParent         bool
}

func (t Task) ToString() (string, error) {
	if t.Title == "" {
		return "", fmt.Errorf("task title must not be empty")
	}
	titleWithOrgLinks, err := markdownLinksToOrgmodeLinks(t.Title)
	if err != nil {
		return "", err
	}
	subtasks, err := t.getSubtasks()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("* TODO %s%s%s%s\n%s", t.getPriority(), titleWithOrgLinks, t.getDateTime(), t.getDescription(), subtasks), nil //TODO(optional): Add [%%] for percentage of completed subtask https://orgmode.org/manual/Breaking-Down-Tasks.html#Breaking-Down-Tasks-into-Subtasks
}

func markdownLinksToOrgmodeLinks(title string) (string, error) {
	markdownLinkRegex, err := regexp.Compile(`\[(.*?)\]\((.*?)\)`)
	if err != nil {
		return "", err
	}
	return markdownLinkRegex.ReplaceAllString(title, "[[$2]]"), nil
}

func (t Task) getPriority() string {
	if t.Priority == 0 {
		return ""
	}
	return fmt.Sprintf("[#%c] ", t.Priority)
}

func (t Task) getDateTime() string {
	if t.Date.IsZero() {
		return ""
	}
	taskTime := ""
	if t.Date.Hour() != 0 || t.Date.Minute() != 0 {
		taskTime = fmt.Sprintf(" %02d:%02d", t.Date.Hour(), t.Date.Minute())
	}
	return fmt.Sprintf("\n\t<%d-%d-%d %.3s%s>", t.Date.Year(), t.Date.Month(), t.Date.Day(), t.Date.Weekday().String(), taskTime)
}

func (t Task) getDescription() string {
	if t.Description == "" {
		return ""
	}
	sanitizedDescription := strings.ReplaceAll(t.Description, "\n", ",\t")
	return fmt.Sprintf("\n\t%s", sanitizedDescription)
}

func (t Task) getSubtasks() (string, error) {
	subtasks := ""
	for _, v := range t.Subtasks {
		subtaskString, err := v.ToString()
		if err != nil {
			return "", err
		}
		subtasks = fmt.Sprintf("%s*%s", subtasks, subtaskString)
	}
	return subtasks, nil
}
