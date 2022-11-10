package util

import (
	"bufio"
	"bytes"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
	"yinghua/pkg/config"
)

func SaveJson(filename, data string) {

	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0744)

	_, _ = file.WriteString(data)
}

func GetGid() (gid uint64) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func ReadText(filename string, line int, limit int) ([]string, error) {
	var data []string
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}

	var count = 0
	s := bufio.NewScanner(file)
	for s.Scan() {
		count++
	}
	//超出行数
	if line > count {
		return []string{}, nil
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	r := bufio.NewScanner(file)
	if line == 0 {
		if limit == 0 {
			line = count
		} else {
			line = limit
		}
	}
	index := 0
	for r.Scan() {
		if limit != 0 && len(data) >= limit {
			break
		}
		if index >= count-line {
			data = append(data, strings.TrimSpace(r.Text()))
		}
		index++
	}
	return data, nil
}

func Copyright() {
	logrus.Infof(`
+---------------------------------------------------------------------------------------+
        ___       ___       ___       ___       ___       ___       ___       ___   
       /\  \     /\__\     /\  \     /\  \     /\  \     /\  \     /\__\     /\  \  
      /::\  \   /:/  /    _\:\  \   /::\  \   /::\  \   _\:\  \   /:/ _/_   /::\  \ 
     /::\:\__\ /:/__/    /\/::\__\ /:/\:\__\ /::\:\__\ /\/::\__\ |::L/\__\ /::\:\__\
     \/\::/  / \:\  \    \::/\/__/ \:\/:/  / \;:::/  / \::/\/__/ |::::/  / \:\:\/  /
       /:/  /   \:\__\    \:\__\    \::/  /   |:\/__/   \:\__\    L;;/__/   \:\/  / 
       \/__/     \/__/     \/__/     \/__/     \|__|     \/__/               \/__/  

               Version: %s Runtime: %s/%s Go Version: %s
+---------------------------------------------------------------------------------------+
`, config.VERSION, runtime.GOOS, runtime.GOARCH, runtime.Version())
}
