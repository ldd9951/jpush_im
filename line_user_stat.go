package jpush_im

import (
	"encoding/json"
	"fmt"
	"bytes"
)

// LineStat 用户在线状态
type UserLineStat struct {
	Login    bool   `json:"login"`
	OnLine   bool   `json:"online"`
	PlatForm string `json:"platform,omitempty"` // SDK各平台：a-Android，i-iOS，j-JS，w-Windows
}

// UsersLineStat 批量用户在线状态
type UsersLineStat struct {
	UserName string         `json:"username"` // 用户名
	Devices  []UserLineStat `json:"devices"`  // 设备登陆状态数组，没有登陆过数组为空
}

// UserStat 用户在线状态查询
func (c *Client) UserStat(userName string) (stat *UserLineStat, err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s/userstat", userName)
	resp, err := c.request("GET", url, nil)
	stat = new(UserLineStat)
	// 反序列化
	err = json.Unmarshal(resp, stat)
	return
}

// UsersStat 批量用户在线状态查询
func (c *Client) UsersStat(usersName []string) (stat []*UsersLineStat, err error) {
	usersNameBody, err := json.Marshal(usersName)
	if err != nil {
		return
	}
	// 请求URL
	url := "/v1/users/userstat"
	resp, err := c.request("POST", url, bytes.NewReader(usersNameBody))
	stat =make([]*UsersLineStat,0)
	// 反序列化
	err = json.Unmarshal(resp, &stat)
	return
}
