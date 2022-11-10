package types

// NodeVideoResponse GetNodeProgress
type NodeVideoResponse struct {
	Code   int    `json:"_code"`
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Result Result `json:"result"`
}
type NodeVideoStudyTotal struct {
	Duration string `json:"duration"`
	Progress string `json:"progress"`
	State    string `json:"state"`
}
type NodeVideoCheat struct {
	State    int `json:"state"`
	Duration int `json:"duration"`
}
type NodeVideoData struct {
	VideoID       string              `json:"videoId"`
	VideoMime     string              `json:"videoMime"`
	VideoDuration int                 `json:"videoDuration"`
	VideoToken    string              `json:"videoToken"`
	StudyTotal    NodeVideoStudyTotal `json:"study_total"`
	Cheat         NodeVideoCheat      `json:"cheat"`
}
type Result struct {
	Data NodeVideoData `json:"data"`
}
