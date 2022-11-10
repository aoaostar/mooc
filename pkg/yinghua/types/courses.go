package types

type CoursesResponse struct {
	Code   int           `json:"_code"`
	Status bool          `json:"status"`
	Msg    string        `json:"msg"`
	Result CoursesResult `json:"result"`
}
type CoursesList struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	Mode          int         `json:"mode"`
	CollegeID     int         `json:"collegeId"`
	CategoryID    int         `json:"categoryId"`
	Lecturers     string      `json:"lecturers"`
	StartDate     string      `json:"startDate"`
	EndDate       string      `json:"endDate"`
	Cover         string      `json:"cover"`
	Credit        string      `json:"credit"`
	Intro         string      `json:"intro"`
	Code          string      `json:"code"`
	StuCount      string      `json:"stuCount"`
	Proclamation  interface{} `json:"proclamation"`
	ClusterID     string      `json:"clusterId"`
	PeriodName    string      `json:"periodName"`
	AddTime       string      `json:"addTime"`
	CreateID      string      `json:"createId"`
	SchoolID      string      `json:"schoolId"`
	CateBid       string      `json:"cateBid"`
	CateMid       string      `json:"cateMid"`
	SignStartTime string      `json:"signStartTime"`
	SignEndTime   string      `json:"signEndTime"`
	SignScope     string      `json:"signScope"`
	SignClass     string      `json:"signClass"`
	LecturerName  string      `json:"lecturerName"`
	Offline       string      `json:"offline"`
	Mission       string      `json:"mission"`
	SignLimit     int         `json:"signLimit"`
	LineLock      string      `json:"lineLock"`
	AddDate       string      `json:"addDate"`
	TplID         string      `json:"tplId"`
	VideoCount    int         `json:"videoCount"`
	VideoLearned  int         `json:"videoLearned"`
	Sign          int         `json:"sign"`
	CollegeName   string      `json:"collegeName"`
	CategoryName  string      `json:"categoryName"`
	Teachers      string      `json:"teachers"`
	ScoreRuleURL  string      `json:"scoreRuleUrl"`
	SignState     int         `json:"signState"`
	Progress      float32     `json:"progress"`
	Progress1     string      `json:"progress1"`
	ResultScore   float32     `json:"resultScore"`
	ResultRank    int         `json:"resultRank"`
	Learned       int         `json:"learned"`
	Learning      int         `json:"learning"`
	StudentCount  int         `json:"studentCount"`
	ClassTeacher  string      `json:"classTeacher"`
	SignInID      int         `json:"signInId"`
	SignedIn      int         `json:"signedIn"`
	State         int         `json:"state"`
	LastNodeID    int         `json:"lastNodeId"`
	TabVideo      bool        `json:"tabVideo"`
	TabFile       bool        `json:"tabFile"`
	TabVote       bool        `json:"tabVote"`
	TabWork       bool        `json:"tabWork"`
	TabExam       bool        `json:"tabExam"`
}
type CoursesResult struct {
	List []CoursesList `json:"list"`
}
