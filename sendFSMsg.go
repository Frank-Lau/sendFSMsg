package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendFSMsg(msg string) {
	sendMsg := struct {
		MsgType string `json:"msg_type"`
		Content struct {
			Text string `json:"text"`
		} `json:"content"`
	}{}
	sendMsg.Content.Text = msg
	sendMsg.MsgType = "text"
	jsonStr, err := json.Marshal(sendMsg)
	if err != nil {
		fmt.Println("SendFSMsg json 序列化失败")
	}

	//发送请求
	c := http.DefaultClient
	//jsonStr := []byte(`{ "msg_type": "text", "content": {"text": "hello world hello java hello go"}}`)
	//req, err := http.NewRequest(http.MethodPost, configs.GetFSConf(), bytes.NewBuffer(jsonStr))
	req, err := http.NewRequest(http.MethodPost, "https://open.feishu.cn/open-apis/bot/v2/hook/caaa84be-4c9c-4153-9565-09f5fc845ba9", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("SendFSMsg发送HTTP请求时发生错误:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = c.Do(req)
	if err != nil {
		fmt.Println("SendFSMsg发送HTTP请求时发生错误:", err)
		return
	}
}
