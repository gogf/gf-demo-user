package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ChatIndexReq struct {
	g.Meta `path:"/chat" method:"get"  tags:"Chat" summary:"Chat homepage"`
}
type ChatIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ChatNameReq struct {
	g.Meta `path:"/chat/name" method:"post"  tags:"Chat" summary:"Customize chat name page"`
	Name   string `v:"required|max-length:21#Why not an awesome name"`
}
type ChatNameRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ChatWebsocketReq struct {
	g.Meta `path:"/chat/websocket" method:"get"  tags:"Chat" summary:"Send message"`
}
type ChatWebsocketRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>" dc:"It redirects to homepage if success"`
}
