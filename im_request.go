package jpush

import "encoding/json"

type User struct {
	UserName string `json:"username",omitempty`
	NickName string `json:"nickname",omitempty`
	MTime    string `json:"mtime",omitempty`
	CTime    string `json:"ctime",omitempty`
}
type UserListResponse struct {
	Total int    `json:"total",omitempty`
	Start int    `json:"start",omitempty`
	Count int    `json:"count",omitempty`
	Users []User `json:"users",omitempty`
}

type UserStatusResponse struct {
	login  bool `json:"login",omitempty`
	online bool `json:"online",omitempty`
}
type MessageBody struct {
	text   string           `json:"text",omitempty`
	extras *json.RawMessage `json:"extras",omitempty`
}
type Message struct {
	TargetType   string      `json:"target_type",omitempty`
	MsgType      string      `json:"msg_type",omitempty`
	TargetName   string      `json:"target_name",omitempty`
	TargetID     string      `json:"target_id",omitempty`
	FromID       string      `json:"from_id",omitempty`
	FromName     string      `json:"from_name",omitempty`
	FromType     string      `json:"from_type",omitempty`
	FromPlatform string      `json:"from_platform",omitempty`
	TromAppkey   string      `json:"from_appkey",omitempty`
	TargetAppkey string      `json:"target_appkey",omitempty`
	MsgBody      MessageBody `json:"msg_body",omitempty`
	CreateTime   uint64      `json:"create_time",omitempty`
	Version      int         `json:"version",omitempty`
	MsgID        int         `json:"msgid",omitempty`
	MsgLevel     int         `json:"msg_level",omitempty` // 0代表应用内消息 1代表跨应用消息
	MsgCtime     uint64      `json:"msg_ctime",omitempty` // 服务器接收到消息的时间，单位毫秒
}
type MessageHistoryResponse struct {
	Total    int       `json:"total",omitempty`
	Cursor   string    `json:"cursor",omitempty`
	Count    int       `json:"count",omitempty`
	Messages []Message `json:"messages",omitempty`
}
