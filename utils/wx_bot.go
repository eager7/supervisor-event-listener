package utils

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
)

const (
	BotKeyPct = "ffaef543-6649-4c8e-8d6e-50cbdb6dc3f7" //通知给管理员
)

type MarkDownParam struct {
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
	MsgType string `json:"msgtype"`
}

type ReqSendCropWxRobotTextMsgResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func WxRobotInfo(textContent, botKey string) error {
	return SendCropWxRobotMsg(fmt.Sprintf(`<font color="info">%s</font>`, textContent), botKey)
}

func WxRobotWarn(textContent, botKey string) error {
	return SendCropWxRobotMsg(fmt.Sprintf(`<font color="warning">%s</font>`, textContent), botKey)
}

func WxRobotDebug(textContent, botKey string) error {
	return SendCropWxRobotMsg(fmt.Sprintf(`<font color="comment">%s</font>`, textContent), botKey)
}

func SendCropWxRobotMsg(textContent, botKey string) error {
	const botUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="
	apiParam := MarkDownParam{
		Markdown: struct {
			Content string `json:"content"`
		}{
			Content: textContent,
		},
		MsgType: "markdown",
	}
	url := botUrl + botKey
	ret := ReqSendCropWxRobotTextMsgResp{}
	req := gorequest.New().Post(url).Retry(3, time.Second*2)
	resp, _, errs := req.SendStruct(&apiParam).EndStruct(&ret)
	if errs != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("SendCropWxRobotTextMsg: %v", errs)
	}
	return nil
}
