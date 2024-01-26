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

func (tp Project) GetID() string {
	return tp.ID
}

func (tp Project) GetParentID() string {
	PID, ok := tp.ParentID.(string)
	if !ok {
		return ""
	}
	return PID
}

func (tp Project) ToTask() *orgmode.Task {
	return &orgmode.Task{
		Title:     tp.Name,
		HasParent: tp.hasParent(),
	}
}

func (tp Project) hasParent() bool {
	if tp.GetParentID() == "" {
		return false
	}
	return true
}
