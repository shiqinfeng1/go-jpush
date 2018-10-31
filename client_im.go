package jpush

import (
	"fmt"

	"github.com/goinggo/mapstructure"
)

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
