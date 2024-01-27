package todoist

import (
	"time"
	"todoist-to-org-mode/orgmode"
)

type todoistTaskDate struct {
	Date        string    `json:"date"`
	IsRecurring bool      `json:"is_recurring"`
	Datetime    time.Time `json:"datetime"`
	String      string    `json:"string"`
	Timezone    string    `json:"timezone"`
}

// TODO: Add dates and recurring dates
func (d todoistTaskDate) toCommonDateTime() time.Time {
	//panic("not implemented")
	return time.Time{}
}

type Task struct {
	CreatorID    string          `json:"creator_id"`
	CreatedAt    time.Time       `json:"created_at"`
	AssigneeID   string          `json:"assignee_id"`
	AssignerID   string          `json:"assigner_id"`
	CommentCount int             `json:"comment_count"`
	IsCompleted  bool            `json:"is_completed"`
	Content      string          `json:"content"`
	Description  string          `json:"description"`
	Due          todoistTaskDate `json:"due"`
	Duration     any             `json:"duration"`
	ID           string          `json:"id"`
	Labels       []string        `json:"labels"`
	Order        int             `json:"order"`
	Priority     int             `json:"priority"`
	ProjectID    string          `json:"project_id"`
	SectionID    string          `json:"section_id"`
	ParentID     string          `json:"parent_id"`
	URL          string          `json:"url"`
}

func (t Task) GetID() string {
	return t.ID
}

func (t Task) GetParentID() string {
	return t.ParentID
}

func (t Task) ToTask() *orgmode.Task {
	return &orgmode.Task{
		Title:             t.Content,
		Description:       t.Description,
		Date:              t.Due.toCommonDateTime(),
		Priority:          t.getPriority(),
		RecurringDuration: t.getRecurringPeriod(),
		HasParent:         t.hasParent(),
	}
}

func (t Task) getPriority() rune {
	return rune('A' - 1 + t.Priority)
}

func (t Task) getRecurringPeriod() time.Duration {
	return time.Duration(10)
}

func (t Task) hasParent() bool {
	if t.ParentID == "" {
		return false
	}
	return true
}
