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
func (d todoistTaskDate) toOrgDate() string {
	//panic("not implemented")
	return ""
}

type todoistPriority int

func (p todoistPriority) toOrgPriority() rune {
	return rune('A' - 1 + p)
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
	Priority     todoistPriority `json:"priority"`
	ProjectID    string          `json:"project_id"`
	SectionID    string          `json:"section_id"`
	ParentID     string          `json:"parent_id"`
	URL          string          `json:"url"`
}

func (tt Task) GetID() string {
	return tt.ID
}

func (tt Task) GetParentID() string {
	return tt.ParentID
}

func (tt Task) ToTask() *orgmode.Task {
	return &orgmode.Task{
		Title:       tt.Content,
		Description: tt.Description,
		Date:        tt.Due.toOrgDate(),
		Priority:    tt.Priority.toOrgPriority(),
		IsRecurring: tt.Due.IsRecurring,
		HasParent:   tt.hasParent(),
	}
}

func (tt Task) hasParent() bool {
	if tt.ParentID == "" {
		return false
	}
	return true
}
