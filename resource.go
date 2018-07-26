package jpush_im

import (
	"fmt"
	"encoding/json"
)

// BaseFileInfo 文件基础信息
type BaseFileInfo struct {
	MediaId    string `json:"media_id"`       // 媒体文件上传到得到的KEY，用于生成下载URL
	MediaCrc32 int32  `json:"media_crc_32"`   // 文件的 CRC32 校验码
	Format     string `json:"format"`         // 语音类型
	Hash       string `json:"hash,omitempty"` // 图片hash值
	FSize      int    `json:"fsize"`          // 文件大小（字节数）
}

// FileImageInfo 图片
type FileImageInfo struct {
	BaseFileInfo
	Width  int `json:"width"`  // 原图片宽度
	Height int `json:"height"` // 原图片高度
}

// FileVoiceInfo 语音
type FileVoiceInfo struct {
	BaseFileInfo
}

// FileInfo 文件信息
type FileInfo struct {
	FName string `json:"fname"` // 发送与接收到的文件名
	BaseFileInfo
}

// Download 文件下载
func (c *Client) Download(mediaId string) (fileUrl string, err error) {
	// 请求URL
	url := fmt.Sprintf("/v1/resource?mediaId=%s", mediaId)
	resp, err := c.request("GET", url, nil)

	file := make(map[string]string, 0)
	// 反序列化
	err = json.Unmarshal(resp, &file)

	_, ok := file["url"]
	if ok {
		fileUrl = file["url"]
	}
	return
}

// UploadFile 上传文件
func (c *Client) UploadFile(filePath string) (fileInfo *FileInfo, err error) {
	file := make(map[string]string, 0)
	file["filename"] = filePath
	// 请求URL
	url := "/v1/resource?type=file"
	resp, err := c.request("POST", url, nil)
	fileInfo = new(FileInfo)
	// 反序列化
	err = json.Unmarshal(resp, fileInfo)
	return
}

// FileImageInfo 上传图片文件
func (c *Client) UploadImage(filePath string) (fileInfo *FileImageInfo, err error) {
	file := make(map[string]string, 0)
	file["filename"] = filePath
	// 请求URL
	url := "/v1/resource?type=image"
	resp, err := c.request("POST", url, nil)
	fileInfo = new(FileImageInfo)
	// 反序列化
	err = json.Unmarshal(resp, fileInfo)
	return
}

// UploadVoice 上传语音文件
func (c *Client) UploadVoice(filePath string) (fileInfo *FileVoiceInfo, err error) {
	file := make(map[string]string, 0)
	file["filename"] = filePath
	// 请求URL
	url := "/v1/resource?type=voice"
	resp, err := c.request("POST", url, nil)
	fileInfo = new(FileVoiceInfo)
	// 反序列化
	err = json.Unmarshal(resp, fileInfo)
	return
}
