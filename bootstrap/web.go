package bootstrap

import (
	"github.com/aoaostar/mooc/pkg/config"
	"github.com/aoaostar/mooc/pkg/util"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
)

func InitWeb() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		file, err := os.Open("./view/index.html")
		if err != nil {
			logrus.Fatal(err.Error())
		}
		readAll, err := io.ReadAll(file)

		if err != nil {
			logrus.Fatal(err.Error())
		}
		_, err = io.WriteString(writer, string(readAll))

		if err != nil {
			logrus.Fatal(err.Error())
		}

	})
	http.HandleFunc("/ajax", func(writer http.ResponseWriter, request *http.Request) {

		text, err := util.ReadText("./logs/aoaostar.log", 0, 100)
		if err != nil {
			logrus.Error(err)

		}
		_, err = io.WriteString(writer, strings.Join(text, "\n"))
		if err != nil {
			logrus.Error(err)
		}

	})
	logrus.Infof("web端启动成功, 请访问 %s 查看服务状态", config.Conf.Global.Server)
	err := http.ListenAndServe(config.Conf.Global.Server, nil)
	if err != nil {
		logrus.Fatal(err.Error())
	}

}
