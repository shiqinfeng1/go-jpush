package jpush

import (
	"fmt"

	"github.com/goinggo/mapstructure"
)

func (c *Client) UsersListAll(start, count uint32) (*UserListResponse, error) {
	args := fmt.Sprintf("/v1/users/?start=%d&count=%d", start, count)
	link := c.imUrl + args
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	mapped, err := resp.Map()
	if err != nil {
		return nil, err
	}
	var s UserListResponse
	if err = mapstructure.Decode(mapped, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

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
	var s UserStatusResponse
	if err = mapstructure.Decode(mapped, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

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
	mapped, err := resp.Map()
	if err != nil {
		return nil, err
	}
	var s MessageHistoryResponse
	if err = mapstructure.Decode(mapped, &s); err != nil {
		return nil, err
	}
	return &s, nil
}
