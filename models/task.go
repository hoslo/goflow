package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	DagId		 uint `gorm:"index" json:"dag_id"`
	IsRoot       bool  `json:"is_root"`
	Name          string `json:"name"`
	MaxTryCount   int `json:"max_try_count"`
	WaitTime      int `json:"wait_time"`
	OperatorType  string `json:"operator_type"`
	OperatorId    uint `json:"operator_id"`
	StaticArgs    string `json:"static_args"`
	HasParentArgs int `json:"has_parent_args"`
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
	DagId      uint `json:"dag_id"`
	TaskId     uint `json:"task_id"`
	DagInstanceId      string `json:"dag_instance_id"`
	OperatorInstanceId uint `json:"operator_instance_id"`
	OperatorInstanceType  string `json:"operator_instance_type"`
	TaskStatus int `json:"task_status"`
	TryCount   int `json:"try_count"`
	Message    string `json:"message"`
}

func NewTaskInstance(dagInstanceId string, dagId uint, taskId uint, operatorInstanceId uint, operatorInstanceType string) *TaskInstance {
	return &TaskInstance{
		TaskId:     taskId,
		DagId:     dagId,
		DagInstanceId:     dagInstanceId,
		TaskStatus: 1,
		OperatorInstanceId: operatorInstanceId,
		OperatorInstanceType: operatorInstanceType,
	}
}

func AddTaskInstance(TaskInstance *TaskInstance) (*TaskInstance, error) {
	result := db.Create(&TaskInstance)

	return TaskInstance, result.Error
}

func UpdateTaskInstance(TaskInstance *TaskInstance) (*TaskInstance, error) {
	result := db.Updates(&TaskInstance)

	return TaskInstance, result.Error
}

func QueryTaskInstances(dagId uint) ([]TaskInstance, error) {
	taskInstances := []TaskInstance{}
	result := db.Where("dag_id = ?", dagId).Find(&taskInstances)

	return taskInstances, result.Error
}

func ExistTaskByName(task *Task) error {
	result := db.Debug().Where("name = ?", task.Name).First(task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func AddTask(task *Task) (*Task, error) {
	result := db.Create(&task)
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

	return task, nil
}

func QueryTasks(dagId uint) ([]Task, error) {

	tasks := []Task{}
	result := db.Where("dag_id = ?", dagId).Find(&tasks)

	return tasks, result.Error
}
