package jpush_im

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
)

// Client 客户端
type Client struct {
	AppKey       string
	MasterSecret string
	IMUrl        string
	BasicAuth    string
}


// NewClient 初始化IM
func NewClient(key, secret string) *Client {
	//  base64_auth_string 的生成算法为：base64(appKey:masterSecret)
	auth := fmt.Sprintf("%s:%s", key, secret)
	// base64 转换
	base := base64.StdEncoding.EncodeToString([]byte(auth))

	cline := new(Client)
	cline.AppKey = key
	cline.MasterSecret = secret
	cline.BasicAuth = base
	cline.IMUrl = "https://api.im.jpush.cn"
	return cline
}

/*
request http接口请求

method：请求方式;GET,POST,PUT,DELETE；
URL：请求地址；
body：内容；
*/
func (c *Client) request(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, c.IMUrl+url, body)
	if err != nil {
		return nil, err
	}
	// header 添加认证信息
	req.Header.Set("Authorization", c.getAuthorization())
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	//
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// 读取body的数据
	buf, err := ioutil.ReadAll(resp.Body)
	// 判断返回状态
	if resp.StatusCode==200 ||  resp.StatusCode==201 {
		return buf, nil
	}else {
		erro := new(Errors)
		// 反序列化
		err = json.Unmarshal(buf, erro)
		return nil, erro
	}
	// 关闭body
	defer resp.Body.Close()

	return nil,errors.New("")
}

// getAuthorization 获取认证信息
func (c *Client) getAuthorization() string {
	str := c.AppKey + ":" + c.MasterSecret
	buf := []byte(str)
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(buf))
}
