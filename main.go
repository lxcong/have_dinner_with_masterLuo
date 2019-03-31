package main

import (
	"flag"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

var config *Config

func hello(ctx *context.Context) {
	//配置微信参数
	wxConf := &wechat.Config{
		AppID:          config.Wx.AppID,
		AppSecret:      config.Wx.AppSecret,
		Token:          config.Wx.Token,
	}
	wc := wechat.NewWechat(wxConf)

	// 传入request和responseWriter
	server := wc.GetServer(ctx.Request, ctx.ResponseWriter)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		text := message.NewText(tlAI(msg.Content))
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

func main() {
	var configFile string
	var err error

	flag.StringVar(&configFile, "c", "config.json", "specify config file")
	config, err = ParseConfig(configFile)

	if err != nil {
		panic("a vailid json config file must exist")
	}

	beego.Any("/", hello)
	beego.Run(":8081")
}
