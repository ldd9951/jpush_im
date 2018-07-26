package jpush_im

import (
	"encoding/json"
	"bytes"
)

// RegisterUsers 注册管理员
func (c *Client) RegisterAdmin(reqAdmin []*RequestUsers) (err error) {
	adminBody, err := json.Marshal(reqAdmin)
	if err != nil {
		return
	}
	// 请求URL
	url := "/v1/users/"
	_, err = c.request("POST", url, bytes.NewReader(adminBody))
	return
}
