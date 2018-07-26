package jpush_im

import (
	"encoding/json"
	"bytes"
)

// RequestUsers 注册用户
type RequestUsers struct {
	UserName string `json:"username,omitempty"` // （必填）用户名  开头：字母或者数字;字母、数字、下划线;英文点、减号、@
	BaseUsers
	Password string `json:"password,omitempty"` // （必填）用户密码。极光IM服务器会MD5加密保存。
	Extras   string `json:"extras,omitempty"`   //  (选填) 用户自定义json对象
}

// 用户基础信息
type BaseUsers struct {
	NickName  string `json:"nickname,omitempty"`  // （选填）用户昵称 不支持的字符：英文字符： \n \r\n
	Avatar    string `json:"avatar,omitempty"`    // （选填）头像  需要填上从文件上传接口获得的media_id
	Birthday  string `json:"birthday,omitempty"`  // （选填）生日 example: 1990-01-24
	Signature string `json:"signature,omitempty"` // （选填）签名 支持的字符：全部，包括 Emoji
	Gender    int    `json:"gender,omitempty"`    // （选填） 性别   0 - 未知， 1 - 男 ，2 - 女
	Region    string `json:"region,omitempty"`    // （选填）地区 支持的字符：全部，包括 Emoji
	Address   string `json:"address,omitempty"`   // （选填）地址 支持的字符：全部，包括 Emoji
}

// ResponseUser 响应
type ResponseUsers struct {
	UserName string `json:"username"` // 用户名
}

// RegisterUsers 注册用户
func (c *Client) RegisterUsers(reqUsers []*RequestUsers) (respUsers []*ResponseUsers, err error) {
	usersBody, err := json.Marshal(reqUsers)
	if err != nil {
		return
	}
	// 请求URL
	url := "/v1/users/"
	resp, err := c.request("POST", url, bytes.NewReader(usersBody))
	// 反序列化
	err = json.Unmarshal(resp, respUsers)
	return
}
