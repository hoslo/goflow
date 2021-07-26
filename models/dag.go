package models

import (
	"errors"
	"fmt"
	"github.com/eapache/queue"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type DAG struct {
	gorm.Model
	UserId uint `gorm:"index" json:"user_id"`
	Name string `json:"name"`
	CronExpression string `json:"cron_expression"`
}

func BFSDag(dagInstanceId string, dagId uint, tasks []Task)  {
	for _, rootTask := range tasks {
		operatorId := rootTask.OperatorId
		OperatorType := rootTask.OperatorType


		if OperatorType == "httpoperator" {
			httpOperator, err := QueryHttpOperator(operatorId)
			if err != nil {
				fmt.Println(err)
				return
			}
			httpOperatorInstance := NewHttpOperatorInstance(httpOperator, httpOperator.ID, "{}")
			httpOperatorInstance, err = AddHttpOperatorInstance(httpOperatorInstance)
			if err != nil {
				fmt.Println(err)
				return
			}
			rootTaskInstance := NewTaskInstance(dagInstanceId, dagId, rootTask.ID, httpOperatorInstance.ID, "httpoperator")
			rootTaskInstance, err = AddTaskInstance(rootTaskInstance)
			if err != nil {
				fmt.Println(err)
				return
			}
			rootTaskInstance.TaskStatus = 1
			for {
				err = httpOperatorInstance.Run()
				if err != nil {
					rootTaskInstance.TryCount += 1
					rootTaskInstance.TaskStatus = 2
					if rootTask.MaxTryCount == rootTaskInstance.TryCount || rootTask.MaxTryCount == 0 {
						break
					}
				} else if httpOperatorInstance.Code == 201 {
					time.Sleep(time.Duration(rootTask.WaitTime))
				} else {
					httpOperatorInstance.Success = true
					rootTaskInstance.TaskStatus = 3
					break
				}
			}

			httpOperatorInstance, err = UpdateHttpOperatorInstance(httpOperatorInstance)
			if err != nil {
				fmt.Println(err)
				return
			}

			rootTaskInstance, err = UpdateTaskInstance(rootTaskInstance)
			if err != nil {
				fmt.Println(err)
				return
			}
			if rootTaskInstance.TaskStatus == 2 {
				return
			}



			childTasks, err := QueryChildTasks(rootTask.ID)
			if err != nil {
				fmt.Println(err)
				return
			}
			BFSDag(dagInstanceId, dagId, childTasks)
		}
	}
}

func (dag *DAG) Run()  {
	tasks, err := QueryTasks(dag.ID)
	dagInstanceId := uuid.NewV4().String()
	if err != nil {
		fmt.Println(err)
		return
	}
	rootTasks := []Task{}
	leafTasks := []Task{}
	for _, task := range tasks {
		if task.IsRoot == true {
			rootTasks = append(rootTasks, task)
		}else {
			leafTasks = append(leafTasks, task)
		}
	}

	if len(rootTasks) == 0 {
		return
	}


	BFSDag(dagInstanceId, dag.ID, rootTasks)

}


func (dag DAG) AddEdge(from, to *Task) error {
	taskToTask := &TaskToTask{
		ParentTaskId: from.ID,
		ChildTaskId: to.ID,
	}
	result := db.Create(&taskToTask)
	return result.Error
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

		var currentChildren []Task
		db.Where("PTaskId = ?", current.ID).Find(currentChildren)
		//fmt.Println("bfs Name", current.Name)

		for _, task := range currentChildren {
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

func AddDAG(dag *DAG) (*DAG, error) {
	result := db.Create(&dag)
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

	return dag, nil
}

func QueryDAGs(userId uint) ([]DAG, error) {

	dags := []DAG{}

	result := db.Where("user_id = ?", userId).Find(&dags)

	return dags, result.Error
}

func QueryDAG(id uint) (DAG, error) {

	dags := DAG{}
	result := db.Where("id = ?", id).Find(&dags)

	return dags, result.Error
}
