package jpush_im

import (
	"fmt"
	"encoding/json"
)

// GetAdminsList 获取管理员列表
func (c *Client) GetAdminsList(start, count int) (list []*UserList, err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/admins?start=%d&count=%d", start, count)
	resp, err := c.request("GET", url, nil)
	err = json.Unmarshal(resp, list)
	return
}
