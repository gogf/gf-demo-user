package model

// Chat Msg 消息结构体
type ChatMsg struct {
	Type string      `json:"type" v:"required#消息类型不能为空"`
	Data interface{} `json:"data" v:""`
	From string      `json:"name" v:""`
}

// 设置昵称请求
type ApiChatSetNameReq struct {
	Name string `json:"type" v:"required|max-length:21#取一个响当当的名字吧|用户昵称最长为21字节"`
}
