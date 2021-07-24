package models

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"net/http"
)

type Operator interface {
	Run() error
}

type HttpOperator struct {
	gorm.Model
	Name       string
	Url        string
	Method     string
	ParentArgs string
	Success    int
	Response   string
	ReturnArgs string
	Error      string
}

var _ Operator = (*HttpOperator)(nil)

func (ho *HttpOperator) Run() error {
	client := &http.Client{}
	bytesData, err := jsoniter.Marshal(ho.ParentArgs)
	if err != nil {
		ho.Error = err.Error()
		return err
	}
	Request, err := http.NewRequest(ho.Method, ho.Url, bytes.NewReader(bytesData))
	if err != nil {
		ho.Error = err.Error()
		return err
	}
	Response, err := client.Do(Request)
	if err != nil {
		ho.Error = err.Error()
		return err
	}
	ho.Response = utils.ToString(Response)
	return nil
}
