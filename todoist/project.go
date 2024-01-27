package todoist

import (
	"todoist-to-org-mode/orgmode"
)

type Project struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	CommentCount   int    `json:"comment_count"`
	Color          string `json:"color"`
	IsShared       bool   `json:"is_shared"`
	Order          int    `json:"order"`
	IsFavorite     bool   `json:"is_favorite"`
	IsInboxProject bool   `json:"is_inbox_project"`
	IsTeamInbox    bool   `json:"is_team_inbox"`
	ViewStyle      string `json:"view_style"`
	URL            string `json:"url"`
	ParentID       any    `json:"parent_id"`
}

func (p Project) GetID() string {
	return p.ID
}

func (p Project) GetParentID() string {
	PID, ok := p.ParentID.(string)
	if !ok {
		return ""
	}
	return PID
}

func (p Project) ToTask() (*orgmode.Task, error) {
	return &orgmode.Task{
		Title:     p.Name,
		HasParent: p.hasParent(),
	}, nil
}

func (p Project) hasParent() bool {
	if p.GetParentID() == "" {
		return false
	}
	return true
}
