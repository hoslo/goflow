package models

import (
	"errors"
	"github.com/eapache/queue"
	"gorm.io/gorm"
)

type DAG struct {
	gorm.Model
	Name string
}

func (dag *DAG) AddRootTask(rootTask *Task) error {
	dagToTask := DAGToTask{
		DAGId: dag.ID,
		TaskId: rootTask.ID,
		IsRoot: 1,
	}
	result := db.Create(&dagToTask)
	return result.Error
}

func (dag *DAG) AddTask(task *Task) error {
	dagToTask := DAGToTask{
		DAGId: dag.ID,
		TaskId: task.ID,
		IsRoot: 0,
	}
	result := db.Create(&dagToTask)
	return result.Error
}

func (dag DAG) AddEdge(from, to *Task) (error, error) {
	taskToTaskChildren := TaskToTask{
		MainTaskId: from.ID,
		TaskId: to.ID,
		IsParent: 0,
	}
	result1 := db.Create(&taskToTaskChildren)
	taskToTaskParent := TaskToTask{
		MainTaskId: to.ID,
		TaskId: from.ID,
		IsParent: 1,
	}
	result2 := db.Create(&taskToTaskParent)
	return result1.Error, result2.Error
}

func (dag *DAG) BFS(root *Task) error {
	q := queue.New()

	visitMap := make(map[string]bool)
	visitMap[root.Name] = true

	q.Add(root)

	for {
		if q.Length() == 0 {
			return nil
		}
		current := q.Remove().(*Task)

		var currentChildren []TaskToTask
		db.Select("taskId").Where("mainTaskId = ?", current.ID).Find(currentChildren)
		//fmt.Println("bfs Name", current.Name)

		for _, taskId := range currentChildren {
			var task Task
			db.Where("id = ?", taskId).First(task)
			if task.Name == root.Name {
				return errors.New("not a DAG")
			}
			if _, ok := visitMap[task.Name]; !ok {
				visitMap[task.Name] = true

				q.Add(task)
			}
		}
	}

}

func AddDAG(userId uint, dag *DAG) (*DAG, error) {
	result := db.Create(&dag)
	if result.Error != nil {
		return nil, result.Error
	}
	userToDAG := UserToDAG{
		UserId: userId,
		DAGId: dag.ID,
	}
	result = db.Create(&userToDAG)
	if result.Error != nil {
		return nil, result.Error
	}

	return dag, nil
}

func DeleteDAG(dag *DAG) (*DAG, error) {
	result := db.Delete(&dag)
	if result.Error != nil {
		return nil, result.Error
	}
	userToDAG := UserToDAG{
		DAGId: dag.ID,
	}
	result = db.Delete(&userToDAG)
	if result.Error != nil {
		return nil, result.Error
	}

	return dag, nil
}

func QueryDAGs(userId uint) ([]DAG, error) {
	var userTODAGs []UserToDAG
	result := db.Where("user_id = ?", userId).Find(&userTODAGs)
	if result.Error != nil {
		return nil, result.Error
	}
	var dagIds []uint
	for _, userTODAG := range userTODAGs {
		dagIds = append(dagIds, userTODAG.DAGId)
	}

	dags := []DAG{}
	if len(dagIds) != 0 {
		result = db.Where("id in ?", dagIds).Find(&dags)
	}

	return dags, result.Error
}

func QueryDAG(id uint) (DAG, error) {

	dags := DAG{}
	result := db.Where("id = ?", id).Find(&dags)

	return dags, result.Error
}
