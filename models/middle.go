package models

import "gorm.io/gorm"

type TaskToTask struct {
	gorm.Model
	ParentTaskId uint
	ChildTaskId uint
}

func QueryChildTasks(parentTaskId uint) ([]Task, error) {
	taskToTasks := []TaskToTask{}
	err := db.Where("parent_task_id = ?", parentTaskId).Find(&taskToTasks).Error

	if err != nil {
		return nil, err
	}
	childTaskIds := []uint{}
	for _, taskToTask := range taskToTasks {
		childTaskIds = append(childTaskIds, taskToTask.ChildTaskId)
	}

	childTasks := []Task{}
	err = db.Where("id in ?", childTaskIds).Find(&childTasks).Error

	return childTasks, err
}

func AddTaskToTask(parentTaskIds []uint, childTaskId uint) error {
	taskToTasks := []TaskToTask{}
	for _, parentTaskId := range parentTaskIds {
		taskToTasks = append(taskToTasks, TaskToTask{
			ParentTaskId: parentTaskId,
			ChildTaskId: childTaskId,
		})
	}

	err := db.Create(taskToTasks).Error

	return err
}
