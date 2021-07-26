package cronjob

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"goflow/models"
)

var Cron *cron.Cron

func InitCron()  {
	Cron = cron.New(cron.WithSeconds())

	users, err := models.QueryUsers()
	fmt.Println(users)
	if err != nil {
		panic("init cronjob false" + err.Error())
	}
	for _, user := range users {
		dags, err := models.QueryDAGs(user.ID)
		if err != nil {
			panic("init cronjob false" + err.Error())
		}
		for _, dag := range dags {
			fmt.Println("CronExpression", dag.CronExpression)
			_, err = Cron.AddJob(dag.CronExpression, &dag)
			if err != nil {
				panic("init cronjob false" + err.Error())
			}
		}
	}

	Cron.Start()
}