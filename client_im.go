package jpush

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/goinggo/mapstructure"
)

//RegisterUsers 批量注册用户
func (c *Client) RegisterUsers(infos []*RegisterUserInfo) error {

	link := c.imUrl + "/v1/users/"

	buf, err := json.Marshal(infos)
	if err != nil {
		return err
	}
	_, err = c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return err
	}
	return nil
}

//RegisterAdmin 注册管理员
func (c *Client) RegisterAdmin(userName, password string) error {

	link := c.imUrl + "/v1/admins/"
	admin := RegisterUserInfo{
		UserName: userName,
		Password: password,
	}
	buf, err := json.Marshal(admin)
	if err != nil {
		return err
	}
	_, err = c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return err
	}
	return nil
}

//SendMessage 发送消息
// msg_type:text
// {
//     "version": 1,
//     "target_type": "single",
//     "target_id": "javen",
//     "from_type": "admin",
//     "from_id": "fang",
//     "msg_type": "text",
//     "msg_body": {
//         "extras": {},
//         "text": "Hello, JMessage!"
//     }
// }

// msg_type:image
// {
//     "version": 1,
//     "target_type": "single",
//     "target_id": "javen",
//     "from_type": "admin",
//     "from_id": "fang",
//     "msg_type": "image",
//     "msg_body": {
//     "media_id": "qiniu/image/CE0ACD035CBF71F8",
//     "media_crc32":2778919613,
//     "width":3840,
//     "height":2160,
//     "fsize":3328738,
//     "format":"jpg"
//     }
// }

// msg_type:voice
// {
//     "version": 1,
//     "target_type": "single",
//     "target_id": "ppppp",
//     "from_type": "admin",
//      "from_id": "admin_caiyh",
//     "msg_type": "voice",
//     "msg_body": {
//     "media_id": "qiniu/voice/j/A96B61EB3AF0E5CDE66D377DEA4F76B8",
//     "media_crc32":1882116055,
//     "hash":"FoYn15bAGRUM9gZCAkvf9dolVH7h",
//     "fsize" :12344;
//      "duration": 6
//     }
// }
// msg_type:custom
// {
//     "version": 1,
//     "target_type": "single",
//     "target_id": "ppppp",
//     "from_type": "admin",
//     "from_id": "admin_caiyh",
//     "msg_type": "custom",
//     "msg_body": {
//         json define yourself
//     }
// }
func (c *Client) SendMessage(fromID, toID, fromName, toName, msgType string, content interface{}) error {

	link := c.imUrl + "/v1/messages"
	msg := MessageMinimum{
		Version:        1,
		TargetType:     "single",
		FromType:       "admin",
		MsgType:        msgType,
		TargetName:     toName,
		FromName:       fromName,
		TargetID:       toID,
		FromID:         fromID,
		NoOffline:      true,
		NoNotification: true,
		MsgBody:        content,
	}
	buf, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return err
	}
	return nil
}

//UsersListAll 用户列表
func (c *Client) UsersListAll(start, count int) (*UserListResponse, error) {
	args := fmt.Sprintf("/v1/users/?start=%d&count=%d", start, count)
	link := c.imUrl + args
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	mapped, err := resp.Map() //{"error": {"message": "basic authentication failed", "code": 899008}}
	if err != nil {
		return nil, err
	}
	if _, exsit := mapped["error"]; exsit == true {
		var e ErrorMsg
		if err = mapstructure.Decode(mapped, &e); err != nil {
			return nil, err
		} else {
			return nil, fmt.Errorf("JPush Returned Code:%d Msg:%v", e.Error.Code, e.Error.Message)
		}

	}
	var s UserListResponse
	if err = mapstructure.Decode(mapped, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

//UserStatus 用户状态
func (c *Client) UserStatus(userName string) (*UserStatusResponse, error) {
	args := fmt.Sprintf("/v1/users/%s/userstat", userName)
	link := c.imUrl + args
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	mapped, err := resp.Map()
	if err != nil {
		return nil, err
	}
	if _, exsit := mapped["error"]; exsit == true {
		var e ErrorMsg
		if err = mapstructure.Decode(mapped, &e); err != nil {
			return nil, err
		} else {
			return nil, fmt.Errorf("JPush Returned Code:%d Msg:%v", e.Error.Code, e.Error.Message)
		}

	}
	var s UserStatusResponse
	if err = mapstructure.Decode(mapped, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

//MessageHistory 消息历史
func (c *Client) MessageHistory(userName string, count uint32, cursor, start, end string) (*MessageHistoryResponse, error) {
	var args string
	if cursor == "" {
		args = fmt.Sprintf("/users/%s/messages?count=%d&begin_time=%s&end_time=%s",
			userName,
			count,
			start,
			end)
	} else {
		args = fmt.Sprintf("/users/%s/messages??cursor=%s",
			userName,
			cursor)
	}
	link := c.imReportUrl + args
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	if len(resp.data) == 0 {
		return &MessageHistoryResponse{}, nil
	}
	mapped, err := resp.Map()
	if err != nil {
		return nil, err
	}
	if _, exsit := mapped["error"]; exsit == true {
		var e ErrorMsg
		if err = mapstructure.Decode(mapped, &e); err != nil {
			return nil, err
		} else {
			return nil, fmt.Errorf("JPush Returned Code:%d Msg:%v", e.Error.Code, e.Error.Message)
		}

	}
	var s MessageHistoryResponse
	if err = mapstructure.Decode(mapped, &s); err != nil {
		return nil, err
	}
	return &s, nil
}
