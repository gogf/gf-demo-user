package ctl_chat

import (
    "gitee.com/johng/gf/g/net/ghttp"
    "gitee.com/johng/gf/g/encoding/gjson"
    "gitee.com/johng/gf/g/util/gvalid"
    "gitee.com/johng/gf/g/os/glog"
    "gitee.com/johng/gf/g/frame/gmvc"
    "gitee.com/johng/gf/g/encoding/ghtml"
    "gitee.com/johng/gf/g/util/gconv"
    "gitee.com/johng/gf/g/container/garray"
    "gitee.com/johng/gf/g/container/gmap"
    "gitee.com/johng/gf/g/container/gset"
    "gitee.com/johng/gf/g/os/gcache"
    "fmt"
)

type Controller struct {
    gmvc.Controller
    ws *ghttp.WebSocket
}

// 消息结构体
type Msg struct {
    Type string      `json:"type" gvalid:"type@required#消息类型不能为空"`
    Data interface{} `json:"data" gvalid:""`
    From string      `json:"name" gvalid:"name@max-length:21#用户昵称最长为21字节"`
}

const (
    SEND_INTERVAL = 2000 // 允许客户端发送聊天消息的间隔时间(毫秒)
)

var (
    // 使用默认的并发安全Map
    users = gmap.NewMap()
    // 使用并发安全的Set，用以用户昵称唯一性校验
    names = gset.NewStringSet()
    // 使用特定的缓存对象，不使用全局缓存对象
    cache = gcache.New()
)

// 聊天室首页，只显示模板内容
func (c *Controller) Index() {
    if c.Session.Contains("name") {
        c.View.Assign("tplMain", "chat/include/chat.html")
    } else {
        c.View.Assign("tplMain", "chat/include/main.html")
    }
    c.View.Display("chat/index.html")
}

// 设置响当当大名
func (c *Controller) SetName() {

}

// WebSocket接口
func (c *Controller) WebSocket() {
    msg := &Msg{}

    // 初始化WebSocket请求
    if ws, err := c.Request.WebSocket(); err != nil {
        glog.Error(err)
        return
    } else {
        c.ws = ws
    }

    // 初始化时设置用户昵称为当前链接信息
    names.Add(c.Request.RemoteAddr)
    users.Set(c.ws, c.Request.RemoteAddr)
    // 初始化后向所有客户端发送上线消息
    c.writeUsers()

    for {
        // 阻塞读取WS数据
        _, msgByte, err := c.ws.ReadMessage()
        if err != nil {
            // 如果失败，那么表示断开，这里清除用户信息
            // 为简化演示，这里不实现失败重连机制
            names.Remove(gconv.String(users.Get(c.ws)))
            users.Remove(c.ws)
            // 通知所有客户端当前用户已下线
            c.writeUsers()
            break
        }
        // JSON参数解析
        if err := gjson.DecodeTo(msgByte, msg); err != nil {
            c.write(Msg{"error", "消息格式不正确: " + err.Error(), ""})
            continue
        }
        // 数据校验
        if e := gvalid.CheckStruct(msg, nil); e != nil {
            c.write(Msg{"error", e.String(), ""})
            continue
        }
        // 用户昵称存储
        name := ghtml.SpecialChars(msg.From)
        if name != "" {
            if v := users.Get(c.ws); v == nil || gconv.String(v) == c.Request.RemoteAddr {
                if names.Contains(name) {
                    c.write(Msg{"error", "用户昵称已存在", ""})
                    continue
                }
                users.Set(c.ws, name)
                names.Add(name)
                names.Remove(c.Request.RemoteAddr)
                c.writeUsers()
            }
        }
        msg.From = gconv.String(users.Get(c.ws))

        // 日志记录
        glog.Cat("chat").Println(msg)

        // WS操作类型
        switch msg.Type {
            // 发送消息
            case "send":
                // 发送间隔检查
                intervalKey := fmt.Sprintf("%p", c.ws)
                if !cache.SetIfNotExist(intervalKey, struct {}{}, SEND_INTERVAL) {
                    c.write(Msg{"error", "您的消息发送得过于频繁，请休息下再重试", ""})
                    continue
                }
                // 有消息时，群发消息
                if msg.Data != nil {
                    if err = c.writeGroup(
                        Msg {"send",
                        ghtml.SpecialChars(gconv.String(msg.Data)),
                        ghtml.SpecialChars(msg.From)}); err != nil {
                        glog.Error(err)
                    }
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
    users.RLockFunc(func(m map[interface{}]interface{}) {
        for user, _ := range m {
            user.(*ghttp.WebSocket).WriteMessage(ghttp.WS_MSG_TEXT, []byte(b))
        }
    })

    return nil
}

// 向客户端返回用户列表
func (c *Controller) writeUsers() error {
    array := garray.NewSortedStringArray(0)
    names.Iterator(func(v string) bool {
        array.Add(v)
        return true
    })
    if err := c.writeGroup(Msg{"list", array.Slice(), ""}); err != nil {
        return err
    }
    return nil
}

