package jpush_im
// 群组维护

import (
	"encoding/json"
	"bytes"
	"fmt"
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

// GroupInfo 群详情
type GroupInfo struct {
	Gid            int64  `json:"gid"`            // 群组ID
	Name           string `json:"name"`           // 群组名字
	Desc           string `json:"desc"`           // 群描述
	AppKey         string `json:"appkey"`         // 用户所属于的应用的appkey
	MaxMemberCount int    `json:"MaxMemberCount"` // 最大成员数
	MTime          string `json:"mtime"`          // 用户最后修改时间
	CTime          string `json:"ctime"`          // 用户创建时间
}

// GetGroupInfo 获取群信息   groupId： 群组ID
func (c *Client) GetGroupInfo(groupId int) (reqMsg *GroupInfo, err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/groups/%d", groupId)
	resp, err := c.request("GET", url, nil)

	reqMsg = new(GroupInfo)
	err = json.Unmarshal(resp, reqMsg)
	return
}

// UpdateGroup 更新群信息
type UpdateGroup struct {
	Name   string `json:"name"`             // 群组名字
	Desc   string `json:"desc"`             // 群描述
	Avatar string `json:"avatar,omitempty"` //（选填）群组头像，上传接口所获得的media_id
}

// UpdateGroup 更新群信息  groupId： 群组ID
func (c *Client) UpdateGroup(groupId int, group *UpdateGroup) (reqMsg *UpdateGroup, err error) {
	body, err := json.Marshal(group)
	if err != nil {
		return
	}
	// 请求URL
	url := fmt.Sprintf("/v1/groups/%d", groupId)
	resp, err := c.request("PUT", url, bytes.NewReader(body))

	reqMsg = new(UpdateGroup)
	err = json.Unmarshal(resp, reqMsg)
	return
}

// DeleteGroup 删除群信息   groupId： 群组ID
func (c *Client) DeleteGroup(groupId int) (err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/groups/%d", groupId)
	_, err = c.request("DELETE", url, nil)
	return
}

// GroupMembers 更新群组成员   groupId： 群组ID;  addMembers:添加组员；removeMembers：删除组员
func (c *Client) GroupMembers(groupId int, addMembers []string, removeMembers []string) (err error) {
	members := map[string]interface{}{
		"add":    addMembers,
		"remove": removeMembers,
	}
	body, err := json.Marshal(members)
	if err != nil {
		return
	}
	// 请求URL
	url := fmt.Sprintf("/v1/groups/%d/members", groupId)
	_, err = c.request("POST", url, bytes.NewReader(body))
	return
}

// GroupMembersList 组员列表
type GroupMembersList struct {
	BaseUsers
	UserName string `json:"username"`
	Flag     int    `json:"flag"` // 0 - 普通群成员;1 - 群主
}

// GetGroupMembersList 获取组员信息列表   groupId:群ID
func (c *Client) GetGroupMembersList(groupId int) (reqList []*GroupMembersList, err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/groups/%d/members/", groupId)
	resp, err := c.request("GET", url, nil)

	reqList = make([]*GroupMembersList, 0)
	// 反序列化
	err = json.Unmarshal(resp, &reqList)
	return
}

// GroupsList 获取当前应用的群组列表
type GroupsList struct {
	Total  int         `json:"total"`
	Start  int         `json:"start"`
	Count  int         `json:"count"`
	Groups []GroupInfo `json:"groups"`
}

// GetGroupsList 获取当前应用的群组列表
func (c *Client) GetGroupsList(start, count int) (info []*GroupsList, err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/groups/?start=%d&count=%d", start, count)
	resp, err := c.request("GET", url, nil)
	info = make([]*GroupsList, 0)
	err = json.Unmarshal(resp, &info)
	return
}

// GetUserGroups 获取某用户的群组列表  userName：当前用户名
func (c *Client) GetUserGroups(userName string) (info []*GroupInfo, err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s/groups/", userName)
	resp, err := c.request("GET", url, nil)

	info = make([]*GroupInfo, 0)
	err = json.Unmarshal(resp, &info)
	return
}

// GroupsShield 群消息屏蔽   userName：当前用户名;  add:添加屏蔽群(群ID)；remove：删除屏蔽群(群ID)
func (c *Client) GroupsShield(userName string, add []int, remove []int) (err error) {
	shield := map[string]interface{}{
		"add":    add,
		"remove": remove,
	}
	body, err := json.Marshal(shield)
	if err != nil {
		return
	}
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s/groupsShield", userName)
	_, err = c.request("POST", url, bytes.NewReader(body))
	return
}

// GroupsMessages 群成员禁言  groupId:群ID;  status：开启或关闭禁言 true表示开启 false表示关闭; username:用户名
func (c *Client) GroupsMessages(groupId int, status bool, username ... string) (err error) {
	body, err := json.Marshal(username)
	if err != nil {
		return
	}
	// 请求URL
	url := fmt.Sprintf("/groups/messages/%d/silence?status=%v", groupId, status)
	_, err = c.request("PUT", url, bytes.NewReader(body))
	return
}

// GroupsOwner 移交群主  groupId:群ID;  username:用户名
func (c *Client) GroupsOwner(groupId int,username string) (err error) {
	owner:=map[string]string{
		"username":username,
	}
	body, err := json.Marshal(owner)
	if err != nil {
		return
	}
	// 请求URL
	url := fmt.Sprintf("/groups/owner/%d", groupId)
	_, err = c.request("PUT", url, bytes.NewReader(body))
	return
}