package models

import "gorm.io/gorm"

type UserToDAG struct {
	gorm.Model
	UserId uint
	DAGId  uint
}

type DAGToTask struct {
	gorm.Model
	DAGId  uint
	TaskId uint
	IsRoot int8
}

type TaskToTask struct {
	gorm.Model
	MainTaskId uint
	TaskId     uint
	IsParent   int8
}

type TaskToOperator struct {
	gorm.Model
	TaskId       uint
	OperatorId   uint
	OperatorType string
}

type TaskToTaskInstance struct {
	gorm.Model
	TaskId       uint
	TaskInstance uint
}
