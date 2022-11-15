package bootstrap

import (
	"fmt"
	"github.com/aoaostar/mooc/pkg/config"
	"github.com/aoaostar/mooc/pkg/task"
	"github.com/aoaostar/mooc/pkg/util"
	"github.com/aoaostar/mooc/pkg/yinghua"
	"github.com/sirupsen/logrus"
)

func Run() {

	InitLog()

	util.Copyright()

	err := InitConfig()

	if err != nil {
		logrus.Fatal(err)
	}

	go InitWeb()

	for _, user := range config.Conf.Users {
		send(user)
	}
	task.Start()

}
func send(user config.User) {

	instance := yinghua.New(user)

	err := instance.Login()
	if err != nil {
		logrus.Fatal(err)
	}
	instance.Output(fmt.Sprintf("登录成功"))

	err = instance.GetCourses()
	if err != nil {
		logrus.Fatal(err)
	}

	instance.Output(fmt.Sprintf("获取全部在学课程成功, 共计 %d 门\n", len(instance.Courses)))

	for _, course := range instance.Courses {
		task.Tasks = append(task.Tasks, task.Task{
			User:   user,
			Course: course,
			Status: false,
		})
	}
}
