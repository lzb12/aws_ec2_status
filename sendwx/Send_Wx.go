package wx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type WeChatMessage struct {
	ToUser  string `json:"touser"`
	ToParty string `json:"toparty"`
	ToTag   string `json:"totag"`
	MsgType string `json:"msgtype"`
	AgentID int    `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

func SendWeChatMessage(agentid int, accessToken, toUser, content string) error {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", accessToken)

	message := WeChatMessage{
		ToUser:  toUser,
		MsgType: "text",
		AgentID: agentid, // 替换为你的企业微信应用AgentID
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
	}

	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查HTTP响应
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("发送企业微信消息失败，状态码：%d", resp.StatusCode)
	}

	// TODO: 解析响应，处理错误情况

	return nil
}

func GetAccessToken(corpid, corpsecret string) (string, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpid, corpsecret)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var accessTokenResp AccessTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&accessTokenResp)
	if err != nil {
		return "", err
	}

	if accessTokenResp.ErrCode != 0 {
		return "", fmt.Errorf("获取企业微信AccessToken失败，错误码：%d，错误信息：%s", accessTokenResp.ErrCode, accessTokenResp.ErrMsg)
	}

	return accessTokenResp.AccessToken, nil
}

func SendWxMessage(corpid, corpsecret, toUser, message string, agentid int) {
	accessToken, err := GetAccessToken(corpid, corpsecret)
	if err != nil {
		fmt.Println("获取企业微信AccessToken失败：", err)
	}
	err = SendWeChatMessage(agentid, accessToken, toUser, message)
	if err != nil {
		fmt.Println("发送企业微信消息失败：", err)
	} else {
		fmt.Println("发送企业微信消息成功！")
	}
}
