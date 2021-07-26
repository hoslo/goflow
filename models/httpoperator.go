package models

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type Operator interface {
	Run() error
}

type HttpOperator struct {
	gorm.Model
	UserId     uint `gorm:"index"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Method     string `json:"method"`
}

type HttpOperatorInstance struct {
	gorm.Model
	Name       string `json:"name"`
	Url        string `json:"url"`
	Method     string `json:"method"`
	HttpOperatorId uint `json:"http_operator_id"`
	Success bool `json:"success"`
	ResponseBody string `json:"response_body"`
	Code int `json:"code"`
	ReturnArgs string `json:"return_args"`
	ParentArgs string `json:"parent_args"`
	Error string `json:"error"`
}

var _ Operator = (*HttpOperatorInstance)(nil)

func NewHttpOperatorInstance(httpOperator HttpOperator, httpOperatorId uint, parentArgs string) *HttpOperatorInstance {
	return &HttpOperatorInstance{
		HttpOperatorId: httpOperatorId,
		Name: httpOperator.Name,
		Url: httpOperator.Url,
		Method: httpOperator.Method,
		ParentArgs: parentArgs,
	}
}

func (hoi *HttpOperatorInstance) Run() error {
	client := &http.Client{}
	bytesData, err := jsoniter.Marshal(hoi.ParentArgs)
	if err != nil {
		hoi.Error = err.Error()
		return err
	}
	Request, err := http.NewRequest(hoi.Method, hoi.Url, bytes.NewReader(bytesData))
	if err != nil {
		hoi.Error = err.Error()
		return err
	}
	Response, err := client.Do(Request)
	if err != nil {
		hoi.Error = err.Error()
		return err
	}
	ResponseBody, err := ioutil.ReadAll(Response.Body)
	hoi.ResponseBody = string(ResponseBody)
	if err != nil {
		hoi.Error = err.Error()
		return err
	}
	hoi.ReturnArgs = jsoniter.Get(ResponseBody, "data").ToString()
	hoi.Code = jsoniter.Get(ResponseBody, "ret").ToInt()

	return nil
}

func AddHttpOperatorInstance(HttpOperatorInstance *HttpOperatorInstance) (*HttpOperatorInstance, error) {
	result := db.Create(&HttpOperatorInstance)

	return HttpOperatorInstance, result.Error
}

func UpdateHttpOperatorInstance(HttpOperatorInstance *HttpOperatorInstance) (*HttpOperatorInstance, error) {
	result := db.Updates(&HttpOperatorInstance)


	return HttpOperatorInstance, result.Error
}


func AddHttpOperator(HttpOperator *HttpOperator) (*HttpOperator, error) {
	result := db.Create(&HttpOperator)

	return HttpOperator, result.Error
}

func DeleteHttpOperator(HttpOperator *HttpOperator) (*HttpOperator, error) {
	result := db.Delete(&HttpOperator)


	return HttpOperator, result.Error
}

func QueryHttpOperators(userId uint) ([]HttpOperator, error) {
	HttpOperators := []HttpOperator{}

	result := db.Where("user_id = ?", userId).Find(&HttpOperators)


	return HttpOperators, result.Error
}

func QueryHttpOperator(id uint) (HttpOperator, error) {

	HttpOperators := HttpOperator{}
	result := db.Where("id = ?", id).Find(&HttpOperators)

	return HttpOperators, result.Error
}
