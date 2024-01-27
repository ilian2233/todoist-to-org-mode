package orgmode

import "bytes"

type TaskCreator []PreOrgTask

type PreOrgTask interface {
	GetID() string
	GetParentID() string
	ToTask() (*Task, error)
}

func (s *TaskCreator) Add(item PreOrgTask) {
	*s = append(*s, item)
}

func (s *TaskCreator) CreateOrgmodeTasks() ([]Task, error) {
	tasksMap := map[string]*Task{}

	// Iterate over todoist tasksMap until all in the org-mode tasksMap map
	for len(tasksMap) != len(*s) {
		for _, object := range *s {
			// If ID exists in map, skip
			if _, ok := tasksMap[object.GetID()]; ok {
				continue
			}
			task, err := object.ToTask()
			if err != nil {
				return nil, err
			}
			// If no parent add to map
			if object.GetParentID() == "" {
				tasksMap[object.GetID()] = task
				continue
			}
			// If parent exists in map add to map and to parent task
			if _, ok := tasksMap[object.GetParentID()]; ok {
				currentTask := task
				tasksMap[object.GetID()] = currentTask
				parent := tasksMap[object.GetParentID()]
				parent.Subtasks = append(parent.Subtasks, currentTask)
			}
		}
	}

	var parentlessTasks []Task
	for _, v := range tasksMap {
		if !v.HasParent {
			parentlessTasks = append(parentlessTasks, *v)
		}
	}
	return parentlessTasks, nil
}

func TasksToBuffer(tasks []Task) (bytes.Buffer, error) {
	var buf bytes.Buffer
	for _, v := range tasks {
		if _, err := buf.WriteString(v.ToString()); err != nil {
			return bytes.Buffer{}, err
		}
	}
	return buf, nil
}
