package todoist

import (
	"fmt"
	"strings"
	"time"
	"todoist-to-org-mode/orgmode"
)

type Task struct {
	CreatorID    string    `json:"creator_id"`
	CreatedAt    time.Time `json:"created_at"`
	AssigneeID   string    `json:"assignee_id"`
	AssignerID   string    `json:"assigner_id"`
	CommentCount int       `json:"comment_count"`
	IsCompleted  bool      `json:"is_completed"`
	Content      string    `json:"content"`
	Description  string    `json:"description"`
	Due          *struct {
		Date        string `json:"date"`
		IsRecurring bool   `json:"is_recurring"`
		Datetime    string `json:"datetime"`
		String      string `json:"string"`
		Timezone    string `json:"timezone"`
	} `json:"due"`
	Duration  any      `json:"duration"`
	ID        string   `json:"id"`
	Labels    []string `json:"labels"`
	Order     int      `json:"order"`
	Priority  int      `json:"priority"`
	ProjectID string   `json:"project_id"`
	SectionID string   `json:"section_id"`
	ParentID  string   `json:"parent_id"`
	URL       string   `json:"url"`
}

func (t Task) GetID() string {
	return t.ID
}

func (t Task) GetParentID() string {
	return t.ParentID
}

func (t Task) ToTask() (*orgmode.Task, error) {
	taskDate, err := t.getDate()
	if err != nil {
		return nil, err
	}

	return &orgmode.Task{
		Title:       t.Content,
		Description: t.Description,
		Date:        taskDate,
		Priority:    t.getPriority(),
		//RecurringDuration: t.getRecurringPeriod(),
		HasParent: t.hasParent(),
	}, nil
}

func (t Task) getDate() (time.Time, error) {
	if t.Due == nil {
		return time.Time{}, nil
	}
	if t.Due.Datetime != "" {
		formattedDate := strings.Replace(t.Due.Datetime, "T", " ", 1)
		return time.Parse(time.DateTime, formattedDate)
	}
	if t.Due.Date != "" { //This check should not be needed as far as I can see if Due is not null date is always set, but just to be sure.
		return time.Parse(time.DateOnly, t.Due.Date)
	}
	return time.Time{}, fmt.Errorf("failed to get date")
}

func (t Task) getPriority() rune {
	//Priority in todoist is 1-4 where 1 is normal and 4 is urgent
	switch t.Priority {
	case 4:
		return 'A'
	case 3:
		return 'B'
	case 2:
		return 'C'
	default:
		return 0
	}
}

func (t Task) getRecurringPeriod() time.Duration {
	panic("implement")
}

func (t Task) hasParent() bool {
	if t.ParentID == "" {
		return false
	}
	return true
}
