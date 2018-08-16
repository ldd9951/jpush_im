package jpush_im

import (
	"encoding/json"
	"fmt"
	"bytes"
)

// UserInfo 查询用户信息
type UserInfo struct {
	BaseUsers
	UserName string `json:"username"`
	Mtime    string `json:"mtime"` // 用户最后修改时间
	Ctime    string `json:"ctime"` // 用户创建时间
}

// UpdateUserInfo 更新用户信息
type UpdateUserInfo struct {
	Extras string `json:"extras,omitempty"` //  (选填) 用户自定义json对象
	BaseUsers
}

// GetUserInfo 获取用户信息
func (c *Client) GetUserInfo(userName string) (user *UserInfo, err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s", userName)
	resp, err := c.request("GET", url, nil)

	user = new(UserInfo)
	// 反序列化
	err = json.Unmarshal(resp, user)
	return
}

// UpdateUser 更新用户信息
func (c *Client) UpdateUser(userName string, reqUsers *UpdateUserInfo) (err error) {
	usersBody, err := json.Marshal(reqUsers)
	if err != nil {
		return
	}
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s", userName)
	_, err = c.request("PUT", url, bytes.NewReader(usersBody))
	return
}

// UpdateUserPassword 修改用户密码
func (c *Client) UpdateUserPassword(userName string, password string) (err error) {
	pwd := map[string]interface{}{
		"new_password": password,
	}
	usersBody, err := json.Marshal(pwd)
	if err != nil {
		return
	}
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s/password", userName)
	_, err = c.request("PUT", url, bytes.NewReader(usersBody))
	return
}

// DeleteUser 删除用户
func (c *Client) DeleteUser(userName string) (err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s", userName)
	_, err = c.request("DELETE", url, nil)
	return
}

// DeleteUsers 批量删除用户
func (c *Client) DeleteUsers(usersName []string) (err error) {
	usersNameBody, err := json.Marshal(usersName)
	if err != nil {
		return
	}
	// 请求URL
	url := "/v1/users/"
	_, err = c.request("DELETE", url, bytes.NewReader(usersNameBody))
	return
}

// BlacklistUser 添加黑名单  userName：当前用户名；usersName：加入黑名单的用户名
func (c *Client) AddBlacklistUser(userName string, usersName []string) (err error) {
	usersNameBody, err := json.Marshal(usersName)
	if err != nil {
		return
	}
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s/blacklist", userName)
	_, err = c.request("PUT", url, bytes.NewReader(usersNameBody))
	return
}

// DeleteBlacklistUser 移除添加黑名单  userName：当前用户名；usersName：加入黑名单的用户名
func (c *Client) DeleteBlacklistUser(userName string, usersName []string) (err error) {
	usersNameBody, err := json.Marshal(usersName)
	if err != nil {
		return
	}
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s/blacklist", userName)
	_, err = c.request("DELETE", url, bytes.NewReader(usersNameBody))
	return
}

// GetBlacklistUsers 获取黑名表列表  userName：当前用户名；usersName：加入黑名单的用户名
func (c *Client) GetBlacklistUsers(userName string, usersName []string) (user []*UserInfo, err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s/blacklist", userName)
	resp, err := c.request("GET", url, nil)

	user = make([]*UserInfo, 0)
	// 反序列化
	err = json.Unmarshal(resp, &user)
	return
}

// Nodisturb 免打扰设置
type Nodisturb struct {
	Single Disturb `json:"single,omitempty"` //  单聊免打扰，支持add remove数组 （可选）
	Group  Disturb `json:"group,omitempty"`  // 群聊免打扰，支持add remove数组（可选
	Global int     `json:"global"`           // 全局免打扰，0或1 0表示关闭，1表示打开 （可选）
}

type Disturb struct {
	Add    []string `json:"add,omitempty"`    // 添加
	Remove []string `json:"remove,omitempty"` // 移除
}

// Nodisturb  免打扰设置  userName：当前用户名；
// single: 单聊免打扰，支持add remove数组 （可选）;
// group 群聊免打扰，支持add remove数组（可选）;
// global 全局免打扰，0或1 0表示关闭，1表示打开 （可选）
func (c *Client) Nodisturb(userName string, single Disturb, group Disturb, global int) (user []*UserInfo, err error) {
	nodisturb := new(Nodisturb)
	nodisturb.Group = group
	nodisturb.Single = single
	nodisturb.Global = global
	body, err := json.Marshal(nodisturb)
	if err != nil {
		return
	}
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s/nodisturb", userName)
	_, err = c.request("POST", url, bytes.NewBuffer(body))
	return
}