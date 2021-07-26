package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

type DAG struct {
	name string
}

func (d *DAG) Run()  {
	fmt.Println(22222222222)
}

func main() {
	d := &DAG{}
	C := cron.New(cron.WithSeconds())
	_, err := C.AddJob("*/5 * * * * *", d)
	if err != nil {
		fmt.Println(err)
	}
	C.Start()
	select{}
}
