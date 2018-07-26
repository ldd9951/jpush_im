package jpush_im

import (
	"fmt"
)

// Forbidden 禁用用户  userName：用户名， disable：true代表禁用用户，false代表激活用户
func (c *Client) Forbidden(userName string, disable bool) (err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/users/%s/forbidden?disable=%v", userName, disable)
	_, err = c.request("GET", url, nil)
	return
}
