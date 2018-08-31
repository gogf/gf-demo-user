package ctl_chat

import (
    "gitee.com/johng/gf/g/net/ghttp"
    "gitee.com/johng/gf/g/encoding/gjson"
    "gitee.com/johng/gf/g/util/gvalid"
    "gitee.com/johng/gf/g/os/glog"
    "gitee.com/johng/gf/g/util/gconv"
    "gitee.com/johng/gf/g/frame/gmvc"
)

type Controller struct {
    gmvc.Controller
    ws *ghttp.WebSocket
}

type Msg struct {
    Type string `json:"type" gvalid:"type@required#消息类型不能为空"`
    Data string `json:"data" gvalid:"data@max-length:80#消息内容不能超过80个字符"`
    From string `json:"name" gvalid:"name@required|length:1,7#用户昵称不能为空|用户昵称长度为1-7个字符"`
}

var (
    users = make(map[*ghttp.WebSocket]string)
)

// 聊天室首页
func (c *Controller) Index() {
    c.Response.Template("chat/index.html", nil)
}

// WebSocket接口
func (c *Controller) WebSocket() {
    msg := &Msg{}
    if ws, err := c.Request.WebSocket(); err != nil {
        glog.Error(err)
        return
    } else {
        c.ws = ws
    }
    for {
        _, msgByte, err := c.ws.ReadMessage()
        if err != nil {
            delete(users, c.ws)
            glog.Error(err)
            break
        }
        if err := gjson.DecodeTo(msgByte, msg); err != nil {
            if err = c.write(Msg{"error", "消息格式不正确: " + err.Error(), ""}); err != nil {
                break
            }
            continue
        }
        // 数据校验
        if e := gvalid.CheckStruct(msg, nil); e != nil {
            if err = c.write(Msg{"error", e.String(), ""}); err != nil {
                break
            }
            continue
        }
        // 判断是否重复连接
        if _, ok := users[c.ws]; !ok {
            if msg.From != "" {
                users[c.ws] = msg.From
            } else {
                users[c.ws] = "匿名"
            }
        }
        switch msg.Type {
            // 发送消息
            case "send":
                // 有消息时，群发消息
                if len(msg.Data) > 0 {
                    if err = c.writeGroup(Msg{"send", msg.Data, ""}); err != nil {
                        glog.Error(err)
                    }
                }
            // 获得当前在线人数
            case "stat":
                if err = c.write(Msg{"stat", gconv.String(len(users)), ""}); err != nil {
                    glog.Error(err)
                }
        }
    }
}

// 向客户端写入消息
func (c *Controller) write(msg Msg) error {
    b, err := gjson.Encode(msg)
    if err != nil {
        return err
    }
    return c.ws.WriteMessage(ghttp.WS_MSG_TEXT, []byte(b))
}

// 群发消息
func (c *Controller) writeGroup(msg Msg) error {
    b, err := gjson.Encode(msg)
    if err != nil {
        return err
    }
    for user, _ := range users {
        user.WriteMessage(ghttp.WS_MSG_TEXT, []byte(b))
    }
    return nil
}

