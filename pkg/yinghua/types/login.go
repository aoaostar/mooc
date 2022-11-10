package types

type LoginResponse struct {
	Code   int         `json:"_code"`
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Result LoginResult `json:"result"`
}
type LoginReset struct {
	State int    `json:"state"`
	Pwd   string `json:"pwd"`
}
type LoginData struct {
	ID           int        `json:"id"`
	Token        string     `json:"token"`
	Avatar       string     `json:"avatar"`
	Number       string     `json:"number"`
	Name         string     `json:"name"`
	IDCard       string     `json:"idCard"`
	Gender       string     `json:"gender"`
	EntryYear    string     `json:"entryYear"`
	Mobile       string     `json:"mobile"`
	WeChat       string     `json:"weChat"`
	Email        string     `json:"email"`
	Intro        string     `json:"intro"`
	ClassID      int        `json:"classId"`
	ClassName    string     `json:"className"`
	CollegeID    int        `json:"collegeId"`
	CollegeName  string     `json:"collegeName"`
	Province     int        `json:"province"`
	City         int        `json:"city"`
	Region       int        `json:"region"`
	ProvinceName string     `json:"provinceName"`
	CityName     string     `json:"cityName"`
	RegionName   string     `json:"regionName"`
	Address      string     `json:"address"`
	Point        int        `json:"point"`
	Rank         int        `json:"rank"`
	TipPass      int        `json:"tipPass"`
	Force        int        `json:"force"`
	Reset        LoginReset `json:"reset"`
}
type LoginResult struct {
	Data LoginData `json:"data"`
}
