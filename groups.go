package jpush_im

import (
	"encoding/json"
	"bytes"
)

type FlagType int

const (
	PrivateGroup FlagType = 1 // 私有群（默认）
	PublicGroup  FlagType = 2 //  公开群
)

// CreateGroups 创建群
type CreateGroups struct {
	OwnerUserName   string   `json:"owner_username"`   //（必填）群主username
	Name            string   `json:"name"`             //（必填）群组名字
	MembersUserName []string `json:"members_username"` //（必填）成员 username
	Avatar          string   `json:"avatar,omitempty"` //（选填）群组头像，上传接口所获得的media_id
	Desc            string   `json:"desc,omitempty"`   //（选填）群描述
	Flag            FlagType `json:"flag,omitempty"`   //（选填） 类型  不指定flag，默认为1
}

// RequestGroup 创建群返回数据
type RequestGroup struct {
	Gid             int64    `json:"gid"`              // 群组ID
	OwnerUserName   string   `json:"owner_username"`   // 群主username
	Name            string   `json:"name"`             // 群组名字
	MembersUserName []string `json:"members_username"` // 成员 username
	Desc            string   `json:"desc"`             // 群描述
	MaxMemberCount  int      `json:"MaxMemberCount"`   // 最大成员数
}

// CreateGroup 创建群
func (c *Client) CreateGroup(group *CreateGroups) (reqMsg *RequestGroup, err error) {
	body, err := json.Marshal(group)
	if err != nil {
		return
	}
	// 请求URL
	url := "/v1/groups/"
	resp, err := c.request("POST", url, bytes.NewReader(body))

	reqMsg = new(RequestGroup)
	err = json.Unmarshal(resp, reqMsg)
	return
}

// GroupInfo 群信息
type GroupInfo struct {
	Gid            int64  `json:"gid"`            // 群组ID
	Name           string `json:"name"`           // 群组名字
	Desc           string `json:"desc"`           // 群描述
	AppKey         string `json:"appkey"`         // 用户所属于的应用的appkey
	MaxMemberCount int    `json:"MaxMemberCount"` // 最大成员数
	MTime          string `json:"mtime"`          // 用户最后修改时间
	CTime          string `json:"ctime"`          // 用户创建时间
}

