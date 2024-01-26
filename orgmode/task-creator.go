package orgmode

type TaskCreator []PreOrgTask

type PreOrgTask interface {
	GetID() string
	GetParentID() string
	ToTask() *Task
}

func (s *TaskCreator) Add(item PreOrgTask) {
	*s = append(*s, item)
}

func (s *TaskCreator) CreateOrgmodeTasks() []Task {
	tasksMap := map[string]*Task{}

	// Iterate over todoist tasksMap until all in the org-mode tasksMap map
	for len(tasksMap) != len(*s) {
		for _, object := range *s {
			// If ID exists in map, skip
			if _, ok := tasksMap[object.GetID()]; ok {
				continue
			}
			// If no parent add to map
			if object.GetParentID() == "" {
				tasksMap[object.GetID()] = object.ToTask()
				continue
			}
			// If parent exists in map add to map and to parent task
			if _, ok := tasksMap[object.GetParentID()]; ok {
				currentTask := object.ToTask()
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
	return parentlessTasks
}
