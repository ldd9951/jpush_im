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
