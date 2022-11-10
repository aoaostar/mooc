package bootstrap

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"yinghua/pkg/config"
)

func InitConfig() error {

	file, err := os.Open("./config.json")
	if err != nil {
		return errors.New("读取配置文件失败: " + err.Error())
	}
	readAll, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(readAll, &config.Conf)

	if err != nil {
		return err
	}
	return nil

}
