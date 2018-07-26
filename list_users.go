package jpush_im

import (
	"time"
	"fmt"
	"encoding/json"
)

// UserList 用户列表或应用管理员列表
type UserList struct {
	Count int      `json:"count"` // 查询条数 最多支持500条
	Total int      `json:"total"` // 总条数
	Start int      `json:"start"` // 当前页
	Users []*Users `json:"users"` // 用户信息
}

// Users 用户信息
type Users struct {
	Mtime    time.Time `json:"mtime"`    // 用户最后修改时间
	UserName string    `json:"username"` // 用户登录名
	Ctime    time.Time `json:"ctime"`    // 用户创建时间
	NickName string    `json:"nickname"` // 用户昵称
}

// GetUsersList 获取用户列表
func (c *Client) GetUsersList(start, count int) (list []*UserList, err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/users?start=%d&count=%d", start, count)
	resp, err := c.request("GET", url, nil)
	err = json.Unmarshal(resp, list)
	return
}
