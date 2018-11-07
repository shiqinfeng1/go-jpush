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
