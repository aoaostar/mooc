package types

type ChaptersResponse struct {
	Code   int            `json:"_code"`
	Status bool           `json:"status"`
	Msg    string         `json:"msg"`
	Result ChaptersResult `json:"result"`
}
type ChaptersNodeList struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	VoteURL         string `json:"voteUrl"`
	TabVideo        bool   `json:"tabVideo"`
	TabFile         bool   `json:"tabFile"`
	TabVote         bool   `json:"tabVote"`
	TabWork         bool   `json:"tabWork"`
	TabExam         bool   `json:"tabExam"`
	VideoDuration   string `json:"videoDuration"`
	UnlockTime      string `json:"unlockTime"`
	NodeLock        int    `json:"nodeLock"`
	UnlockTimeStamp int    `json:"unlockTimeStamp"`
	VideoState      int    `json:"videoState"`
	Duration        string `json:"duration"`
	Index           string `json:"index"`
	Idx             int    `json:"idx"`
}
type ChaptersList struct {
	ID       int                `json:"id"`
	Name     string             `json:"name"`
	NodeList []ChaptersNodeList `json:"nodeList"`
	Idx      int                `json:"idx"`
}
type ChaptersResult struct {
	List []ChaptersList `json:"list"`
}

type StudyNodeResponse struct {
	Code     int             `json:"_code"`
	NeedCode bool            `json:"need_code"`
	Status   bool            `json:"status"`
	Msg      string          `json:"msg"`
	Result   StudyNodeResult `json:"result"`
}
type StudyNodeData struct {
	StudyID int `json:"studyId"`
}
type StudyNodeResult struct {
	Data StudyNodeData `json:"data"`
}
