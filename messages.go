package jpush_im

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// TargetType 接收者类型
type TargetType string

// FromType 发送方来源
type FromType string

// MsgType 消息类型
type MsgType string

const (
	// 接收者类型
	Single TargetType = "single"
	Group  TargetType = "group"

	// 发送方来源
	User  FromType = "user"
	Robot FromType = "robot"
	Admin FromType = "admin"

	// 消息类型
	Text   MsgType = "text"
	Voice  MsgType = "voice"
	Image  MsgType = "image"
	Custom MsgType = "custom"
)

// MessagesSend 发送消息
type MessagesSend struct {
	Version        int         `json:"version"`                   // (必填) 协议版本号。第一版本：1，以此类推
	TargetType     TargetType  `json:"target_type"`               // (必填) 接收者类型 选项：single, group
	TargetId       string      `json:"target_id"`                 // (必填) 接收者ID。 可能值：${username}, ${gid}
	TargetAppKey   string      `json:"target_appkey,omitempty"`   // (可选) 跨应用目标appkey（选填）
	TargetName     string      `json:"target_name,omitempty"`     // (可选) 接收者的展示名。
	FromType       FromType    `json:"from_type,omitempty"`       // (必填) 发送方来源 当前只限admin用户，必须先注册admin用户
	FromId         string      `json:"from_id"`                   // (必填) 发送者 username
	FromName       string      `json:"from_name,omitempty"`       // (可选) 发送方展示名
	NoOffline      bool        `json:"no_offline,omitempty"`      // (可选) 消息是否离线存储 true或者false，默认为false，表示需要离线存储
	NoNotification bool        `json:"no_notification,omitempty"` // (可选) 消息是否在通知栏展示 true或者false，默认为false，表示在通知栏展示
	Title          string      `json:"title,omitempty"`           // (可选) 通知的标题
	Alert          string      `json:"alert,omitempty"`           // (可选) 通知的内容
	MsgType        MsgType     `json:"msg_type"`                  // (必填) 消息类型
	MsgBody        interface{} `json:"msg_body"`                  // (必填) 消息实体
}

type RequestMsg struct {
	MsgId    int64 `json:"msg_id"`
	MsgCtime int64  `json:"msg_ctime"`
}

// MsgText  msg_type = text
type MsgText struct {
	Extras string `json:"extras,omitempty"` // (可选) JsonObject 用于附加参数。所有的消息类型都可以带此字段。
	Text   string `json:"text"`             // (必填) 文本类型消息内容
}

// VoiceText   msg_type = voice
type VoiceText struct {
	MediaId    string `json:"media_id"`       // (必填) 媒体文件上传到得到的KEY，用于生成下载URL
	MediaCrc32 int32  `json:"media_crc_32"`   // (必填) 文件的 CRC32 校验码
	Duration   int    `json:"duration"`       // (必填) 语音时长（单位：秒）
	Format     string `json:"format"`         // (必填) 语音类型
	Hash       string `json:"hash,omitempty"` // (可选) 图片hash值
	FSize      int    `json:"fsize"`          // (必填) 文件大小（字节数）
}

// ImageText  msg_type = image
type ImageText struct {
	MediaId    string `json:"media_id"`           // (必填) 媒体文件上传到得到的KEY，用于生成下载URL
	MediaCrc32 int32  `json:"media_crc_32"`       // (必填) 文件的 CRC32 校验码
	Width      int    `json:"width"`              // (必填) 原图片宽度
	Height     int    `json:"height"`             // (必填) 原图片高度
	Format     string `json:"format"`             // (必填) 图片格式
	Hash       string `json:"hash,omitempty"`     // (可选) 图片hash值
	FSize      int    `json:"fsize"`              // (必填) 文件大小（字节数）
	ImgLink    string `json:"img_link,omitempty"` // (可选) 图片链接
}

// SendMsg 发送消息
func (c *Client) SendMsg(msg *MessagesSend) (reqMsg *RequestMsg, err error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return
	}
	// 请求URL
	url := "/v1/messages"
	resp, err := c.request("POST", url, bytes.NewReader(body))

	reqMsg = new(RequestMsg)
	err = json.Unmarshal(resp, reqMsg)
	return
}

// RetractMessages 消息撤回   username：发送此msg的用户名；msgid：消息msgid
func (c *Client) RetractMessages(userName string, msgId int64) (err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/messages/%s/%d/retract", userName, msgId)
	_, err = c.request("POST", url, nil)
	return
}
