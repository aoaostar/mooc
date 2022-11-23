package yinghua

import (
	"bytes"
	"errors"
	"fmt"
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/aoaostar/mooc/pkg/config"
	"github.com/aoaostar/mooc/pkg/util"
	"github.com/aoaostar/mooc/pkg/yinghua/types"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

type YingHua struct {
	User    config.User
	Courses []types.CoursesList
	client  *resty.Client
	sync.Mutex
}

func New(user config.User) *YingHua {

	var client = resty.New()
	client.SetBaseURL(user.BaseURL)
	client.SetRetryCount(3)
	client.SetHeader("user-agent", browser.Mobile())
	return &YingHua{
		User:   user,
		client: client,
	}

}

func (i *YingHua) Login() error {

	resp := new(types.LoginResponse)
	resp2, err := i.client.R().SetFormData(map[string]string{
		"platform":  "Android",
		"username":  i.User.Username,
		"password":  i.User.Password,
		"pushId":    "140fe1da9e67b9c14a7",
		"school_id": strconv.Itoa(i.User.SchoolID),
		"imgSign":   "533560501d19cc30271a850810b09e3e",
		"imgCode":   "cryd",
	}).
		SetResult(resp).
		Post("/api/login.json")

	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return errors.New(resp.Msg)
	}

	i.client.SetCookies(resp2.Cookies())

	i.client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		// 这里要加锁，避免并发操作map,引发panic
		i.Lock()
		defer i.Unlock()
		c.FormData.Set("token", resp.Result.Data.Token)
		return nil
	})

	return nil

}

func (i *YingHua) GetCourses() error {

	resp := new(types.CoursesResponse)
	_, err := i.client.R().
		SetResult(resp).
		Post("/api/course.json")

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return errors.New(resp.Msg)
	}
	i.Courses = resp.Result.List
	return nil
}

func (i *YingHua) GetChapters(course types.CoursesList) ([]types.ChaptersList, error) {

	resp := new(types.ChaptersResponse)
	_, err := i.client.R().
		SetResult(resp).
		SetFormData(map[string]string{
			"courseId": strconv.Itoa(course.ID),
		}).
		Post("/api/course/chapter.json")

	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, errors.New(resp.Msg)
	}
	return resp.Result.List, nil
}

func (i *YingHua) StudyCourse(course types.CoursesList) error {
	chapters, err := i.GetChapters(course)
	if err != nil {
		return err
	}
	for _, chapter := range chapters {
		i.StudyChapter(chapter)
	}

	return nil
}

func (i YingHua) StudyChapter(chapter types.ChaptersList) {

	i.Output(fmt.Sprintf("当前第 %d 章, [%s][chapterId=%d]", chapter.Idx, chapter.Name, chapter.ID))
	for _, node := range chapter.NodeList {
		// 试题跳过
		if node.TabVideo {
			i.StudyNode(node)
		}
	}

}

func (i YingHua) StudyNode(node types.ChaptersNodeList) {
startStudy:
	i.Output(fmt.Sprintf("当前第 %d 课, [%s][nodeId=%d]", node.Idx, node.Name, node.ID))
	var studyTime = 1
	var studyId = 0
	var nodeProgress = types.NodeVideoData{
		StudyTotal: types.NodeVideoStudyTotal{
			Progress: "0.00",
		},
	}
	var flag = true
	go func() {
		for flag {
			var err error
			nodeProgress, err = i.GetNodeProgress(node)
			if err != nil {
				i.OutputWith(fmt.Sprintf("%s[nodeId=%d], %s[studyId=%d]", node.Name, node.ID, err.Error(), studyId), logrus.Errorf)
				flag = false
				break
			}
			if nodeProgress.StudyTotal.State == "2" {
				node.VideoState = 2
				break
			}
			time.Sleep(time.Second * 10)
		}
	}()

	for node.VideoState != 2 {
		if !flag {
			goto startStudy
		}

		var formData = map[string]string{
			"nodeId":    strconv.Itoa(node.ID),
			"studyTime": strconv.Itoa(studyTime),
			"studyId":   strconv.Itoa(studyId),
		}
	captcha:
		var resp = new(types.StudyNodeResponse)
		_, err := i.client.R().
			SetFormData(formData).
			SetResult(resp).
			Post("/api/node/study.json")
		if err != nil {
			i.OutputWith(fmt.Sprintf("%s[nodeId=%d], %s[studyId=%d][studyTime=%d]", node.Name, node.ID, err.Error(), studyId, studyTime), logrus.Errorf)
			continue
		}
		if resp.Code != 0 {
			i.OutputWith(fmt.Sprintf("%s[nodeId=%d], %s[studyId=%d][studyTime=%d]", node.Name, node.ID, resp.Msg, studyId, studyTime), logrus.Errorf)
			if resp.NeedCode {
				formData["code"] = i.FuckCaptcha() + "_"
				goto captcha
			}
			flag = false
			break
		}
		studyId = resp.Result.Data.StudyID
		if nodeProgress.StudyTotal.Progress == "" {
			nodeProgress.StudyTotal.Progress = "0.00"

		}
		parseFloat, err := strconv.ParseFloat(nodeProgress.StudyTotal.Progress, 64)

		if err != nil {
			i.OutputWith(fmt.Sprintf("%s[nodeId=%d], %s[studyId=%d]", node.Name, node.ID, err.Error(), studyId), logrus.Errorf)
			continue
		}
		i.Output(fmt.Sprintf("%s[nodeId=%d], %s[studyId=%d], 当前进度: %.f%%", node.Name, node.ID, resp.Msg, studyId, parseFloat*100))
		studyTime += 10
		time.Sleep(time.Second * 10)
	}
}

func (i YingHua) GetNodeProgress(node types.ChaptersNodeList) (types.NodeVideoData, error) {

	var resp = new(types.NodeVideoResponse)
	_, err := i.client.R().
		SetFormData(map[string]string{
			"nodeId": strconv.Itoa(node.ID),
		}).
		SetResult(resp).
		Post("/api/node/video.json")
	if err != nil {
		i.OutputWith(fmt.Sprintf("%s[nodeId=%d], %s", node.Name, node.ID, err.Error()), logrus.Errorf)
		return resp.Result.Data, nil
	}
	if resp.Code != 0 {
		return resp.Result.Data, errors.New(resp.Msg)
	}
	return resp.Result.Data, nil
}

func (i YingHua) FuckCaptcha() string {

	i.Output("正在识别验证码")
	response, err := i.client.R().
		Get(fmt.Sprintf("/service/code/aa?t=%d", time.Now().UnixNano()))

	if err != nil {
		i.OutputWith(err.Error(), logrus.Errorf)
	}
	var resp = new(types.Captcha)
	client := resty.New()
	_, err = client.R().
		SetFileReader("file", "image.png", bytes.NewReader(response.Body())).
		SetResult(resp).
		Post("https://api.opop.vip/captcha/recognize")

	if err != nil {
		i.OutputWith(err.Error(), logrus.Errorf)
	}
	if resp.Status != "ok" {
		i.OutputWith(resp.Message, logrus.Errorf)
	}
	s := resp.Data.(string)
	i.Output(fmt.Sprintf("验证码识别成功: %s", s))
	return s
}
func (i YingHua) Output(message string) {
	i.OutputWith(message, logrus.Infof)
}

func (i YingHua) OutputWith(message string, writer func(format string, args ...interface{})) {
	writer("[协程ID=%d][%s] %s", util.GetGid(), i.User.Username, message)
}
