package jpush_im

// 好友管理

import (
	"encoding/json"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
)

// AddFriends 添加好友  user:当前用户；username:用户名
func (c *Client) AddFriends(user string, username ... string) (err error) {
	body, err := json.Marshal(username)
	if err != nil {
		return
	}
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s/friends", user)
	_, err = c.request("POST", url, bytes.NewReader(body))
	return
}

// DelFriends 删除好友  user:当前用户;username:用户名
func (c *Client) DelFriends(user string, username ... string) (err error) {
	body, err := json.Marshal(username)
	if err != nil {
		return
	}
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s/friends", user)
	_, err = c.request("DELETE", url, bytes.NewReader(body))
	return
}

// Friends 更新好友备注
type Friends struct {
	NoteName string `json:"note_name"` // 表示要添加的好友列表
	Others   string `json:"others"`    // 其他备注信息
	UserName string `json:"username"`
}

// UpdateFriends 更新好友备注  user:当前用户;
func (c *Client) UpdateFriends(user string, friends []*Friends) (err error) {
	body, err := json.Marshal(friends)
	if err != nil {
		return
	}
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s/friends", user)
	_, err = c.request("PUT", url, bytes.NewReader(body))

	if v, ok := err.(*Errors); ok {
		if v.Err.Code == 899002 {
			err = errors.New("用户不存在")
			return
		}
		if v.Err.Code == 899003 {
			err = errors.New("Request Body json格式不符合要求，json参数不符合要求")
			return
		}
	}
	return
}

// FriendsList 获取好友列表
type FriendsList struct {
	BaseUsers
	UserName string `json:"username"`
	Mtime    string `json:"mtime"` // 用户最后修改时间
	Ctime    string `json:"ctime"` // 用户创建时间
	NoteName string `json:"note_name"`
	Others   string `json:"others"`
	AppKey   string `json:"appkey"`
}

// GetFriendsList 获取好友列表
func (c *Client) GetFriendsList(user string)(friends []*FriendsList,err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s/friends", user)
	resp, err := c.request("GET", url, nil)
	friends = make([]*FriendsList,0)
	// 反序列化
	err = json.Unmarshal(resp, friends)
	return
}
