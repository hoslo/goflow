package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name          string
	MaxTryCount   int
	WaitTime      int
	OperatorType  string
	OperatorId    uint
	StaticArgs    string
	HasParentArgs int
}

func (t *Task) NewTask(name string, operatorType string, operatorId uint, staticArgs string) *Task {
	return &Task{
		Name:         name,
		OperatorType: operatorType,
		OperatorId:   operatorId,
		StaticArgs:   staticArgs,
	}
}

type TaskInstance struct {
	gorm.Model
	TaskId     int
	TaskStatus int
	TryCount   int
	Message    string
}

func (ti *TaskInstance) NewTaskInstance(taskId int) *TaskInstance {
	return &TaskInstance{
		TaskId:     taskId,
		TaskStatus: 1,
	}
}

func AddTask(dagId uint, task *Task) (*Task, error) {
	result := db.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	dagToTask := DAGToTask{
		TaskId: task.ID,
		DAGId:  dagId,
	}
	result = db.Create(&dagToTask)
	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

func DeleteTask(task *Task) (*Task, error) {
	result := db.Delete(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	dagToTask := DAGToTask{
		TaskId: task.ID,
	}
	result = db.Delete(&dagToTask)
	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

func QueryTasks(dagId uint) ([]Task, error) {
	var dagToTasks []DAGToTask
	result := db.Where("dag_id = ?", dagId).Find(&dagToTasks)
	if result.Error != nil {
		return nil, result.Error
	}
	var taskIds []uint
	for _, dagToTask := range dagToTasks {
		taskIds = append(taskIds, dagToTask.TaskId)
	}

	tasks := []Task{}
	if len(taskIds) != 0 {
		result = db.Where("id in ?", taskIds).Find(&tasks)
	}

	return tasks, result.Error
}
