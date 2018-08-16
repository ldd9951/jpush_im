# jpush_im
===================

## 概述
-----------------------------------
   这是IM REST API 的 go 版本封装开发包
   IM REST API 文档：https://docs.jiguang.cn/jmessage/server/rest_api_im/#_24

## 发送消息例子
-----------------------------------
```
package main

import (
	"jpush-im"
	"log"
	"time"
	"fmt"
)

var AppKey = ""
var MasterSecret = ""

func main() {
	client := jpush_im.NewClient(AppKey, MasterSecret)
	msg := new(jpush_im.MessagesSend)
	msg.Version = 1
	msg.TargetType = jpush_im.Single
	msg.TargetId = "javen"
	msg.FromType = jpush_im.Admin
	msg.MsgType = jpush_im.Text
	msg.FromId = "admin"
	text := new(jpush_im.MsgText)
	text.Text = fmt.Sprintf("当前时间%s", time.Now().Format("2006-01-02 15:04:05"))
	msg.MsgBody = text

	meg, err := client.SendMsg(msg)
	if err != nil {
		log.Println("发送错误", err)
		return
	}

	log.Println("发送成功", meg)
}

```

## 使用
-----------------------------------
|   函数   |   说明   |
| ---- | ---- |
| 用户维护 |  |
|   GetUserInfo(userName string) (user *UserInfo, err error)   | 获取用户信息     |
| UpdateUser(userName string, reqUsers *UpdateUserInfo) (err error) | 更新用户信息 |
| UpdateUserPassword(userName string, password string) (err error) | 修改密码 |
| DeleteUser(userName string) (err error) | 删除用户 |
| DeleteUsers(usersName []string) (err error) | 批量删除用户 |
| AddBlacklistUser(userName string, usersName []string) (err error) | 添加黑名单 |
| DeleteBlacklistUser(userName string, usersName []string) (err error) | 移除黑名单 |
| GetBlacklistUsers(userName string, usersName []string) (user []*UserInfo, err error) | 黑名单列表 |
| Nodisturb(userName string, single Disturb, group Disturb, global int) (user []*UserInfo, err error) | 免打扰设置 |
| Forbidden(userName string, disable bool) | 禁用用户 |
| UserStat(userName string) (stat *UserLineStat, err error) | 用户在线状态查询 |
| UsersStat(usersName []string) (stat []*UsersLineStat, err error) | 批量用户在线状态查询 |
| 用户注册 |  |
| RegisterUsers(reqUsers []*RequestUsers) (respUsers []*ResponseUsers, err error) | Admin 注册 |
| GetAdminsList(start, count int) (list []*UserList, err error) | 获取应用管理员列表 |
| 消息相关 |  |
| SendMsg(msg *MessagesSend) (reqMsg *RequestMsg, err error) | 发送消息 |
| 媒体文件下载与上传 |  |
| Download(mediaId string) (fileUrl string, err error) | 文件下载 |
| UploadFile(filePath string) (fileInfo *FileInfo, err error) | 文件上传 |
| 群组维护 |  |
| CreateGroup(group *CreateGroups) (reqMsg *RequestGroup, err error) | 创建群组 |
| GetGroupInfo(groupId int) (reqMsg *GroupInfo, err error) | 获取群组详情 |
| UpdateGroup(groupId int, group *UpdateGroup) (reqMsg *UpdateGroup, err error) | 更新群组信息 |
| DeleteGroup(groupId int) (err error) | 删除群组 |
| GroupMembers(groupId int, addMembers []string, removeMembers []string) | 更新群组成员 |
| GetGroupMembersList(groupId int) (reqList []*GroupMembersList, err error) | 获取群组成员列表 |
| GetUserGroups(userName string) (info []*GroupInfo, err error) | 获取某用户的群组列表 |
| GetGroupsList(start, count int) (info []*GroupsList, err error) | 获取当前应用的群组列表 |
| GroupsShield(userName string, add []int, remove []int) | 群消息屏蔽 |
| GroupsMessages(groupId int, status bool, username ... string) | 群成员禁言 |
| GroupsOwner(groupId int,username string) | 移交群主 |
| 好友| |
| AddFriends(user string, username ... string) (err error)  | 添加好友 |
| DelFriends(user string, username ... string) (err error) | 删除好友 |
| UpdateFriends(user string, friends []*Friends) (err error) | 更新好友备注 |
| GetFriendsList(user string)(friends []*FriendsList,err error) | 获取好友列表 |



