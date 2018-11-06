package jpush

import "encoding/json"

type ErrorMsg struct {
	Error ErrorMsgBody `json:"error",omitempty`
}
type ErrorMsgBody struct {
	Message string `json:"message",omitempty`
	Code    int    `json:"code",omitempty`
}
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
	Login  bool `json:"login",omitempty`
	Online bool `json:"online",omitempty`
}
type MessageBodyText struct {
	Text   string           `json:"text"`             //消息内容 （必填）
	Extras *json.RawMessage `json:"extras",omitempty` //选填的json对象 开发者可以自定义extras里面的key value（选填）
}
type MessageBodyImage struct {
	MediaID    string `json:"media_id"`       //文件上传之后服务器端所返回的key，用于之后生成下载的url（必填）
	MediaCrc32 uint64 `json:"media_crc32"`    //文件的crc32校验码，用于下载大图的校验 （必填）
	Width      int    `json:"width"`          //图片原始宽度（必填）
	Height     int    `json:"height"`         //图片原始高度（必填）
	Format     string `json:"format"`         //图片格式（必填）
	Hash       string `json:"hash",omitempty` //图片hash值（可选）
	Fsize      int    `json:"fsize"`          //文件大小（字节数）（必填）
}
type MessageBodyVoice struct {
	MediaID    string `json:"media_id"`       //文件上传之后服务器端所返回的key，用于之后生成下载的url（必填）
	MediaCrc32 uint64 `json:"media_crc32"`    //文件的crc32校验码，用于下载大图的校验 （必填）
	Duration   int    `json:"duration"`       //音频时长（必填）
	Hash       string `json:"hash",omitempty` //音频hash值（可选）
	Fsize      int    `json:"fsize"`          //文件大小（字节数）（必填）
}
type Message struct {
	TargetType     string           `json:"target_type"`           //发送目标类型 single - 个人，group - 群组 chatroom - 聊天室（必填）
	MsgType        string           `json:"msg_type"`              //发消息类型 text - 文本，image - 图片, custom - 自定义消息（msg_body为json对象即可，服务端不做校验）voice - 语音 （必填）
	TargetName     string           `json:"target_name",omitempty` //接受者展示名（选填）
	TargetID       string           `json:"target_id"`             //目标id single填username group 填Group id chatroom 填chatroomid（必填）
	FromID         string           `json:"from_id"`               //发送者的username （必填
	FromName       string           `json:"from_name",omitempty`   //发送者展示名（选填）
	FromType       string           `json:"from_type"`             //发送消息者身份 当前只限admin用户，必须先注册admin用户 （必填）
	FromPlatform   string           `json:"from_platform",omitempty`
	TromAppkey     string           `json:"from_appkey",omitempty`
	TargetAppkey   string           `json:"target_appkey"` //跨应用目标appkey（选填）
	MsgBody        *json.RawMessage `json:"msg_body"`      //消息体,根据type不同结构不同
	CreateTime     uint64           `json:"create_time",omitempty`
	Version        int              `json:"version",omitempty` //版本号 目前是1 （必填）
	MsgID          int              `json:"msgid",omitempty`
	MsgLevel       int              `json:"msg_level",omitempty`       // 0代表应用内消息 1代表跨应用消息
	MsgCtime       uint64           `json:"msg_ctime",omitempty`       // 服务器接收到消息的时间，单位毫秒
	NoOffline      bool             `json:"no_offline",omitempty`      //消息是否离线存储 true或者false，默认为false，表示需要离线存储（选填）
	NoNotification bool             `json:"no_notification",omitempty` //消息是否在通知栏展示 true或者false，默认为false，表示在通知栏展示（选填）
}
type MessageHistoryResponse struct {
	Total    int       `json:"total",omitempty`
	Cursor   string    `json:"cursor",omitempty`
	Count    int       `json:"count",omitempty`
	Messages []Message `json:"messages",omitempty`
}

type RegisterUserInfo struct {
	USerName string `json:"username"`
	Password string `json:"password"`
}
