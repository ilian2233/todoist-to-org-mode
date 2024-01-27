package orgmode

import (
	"fmt"
	"regexp"
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
	return fmt.Sprintf("* TODO %s%s%s%s\n%s", t.getPriority(), titleWithOrgLinks, t.getDate(), t.getDescription(), subtasks), nil //TODO(optional): Add [%%] for percentage of completed subtask https://orgmode.org/manual/Breaking-Down-Tasks.html#Breaking-Down-Tasks-into-Subtasks
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
