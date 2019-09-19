package chat

import (
	"fmt"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/encoding/ghtml"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

// Controller 控制器结构体
type Controller struct {
	gmvc.Controller
	ws *ghttp.WebSocket
}

// Msg 消息结构体
type Msg struct {
	Type string      `json:"type" gvalid:"type@required#消息类型不能为空"`
	Data interface{} `json:"data" gvalid:""`
	From string      `json:"name" gvalid:""`
}

const (
	// SendInterval 允许客户端发送聊天消息的间隔时间(毫秒)
	SendInterval  = 1000
	nameCheckRule = "required|max-length:21"
	nameCheckMsg  = "取一个响当当的名字吧|用户昵称最长为21字节"
)

var (
	// 使用默认的并发安全Map
	users = gmap.New()
	// 使用并发安全的Set，用以用户昵称唯一性校验
	names = gset.NewStrSet()
	// 使用特定的缓存对象，不使用全局缓存对象
	cache = gcache.New()
)

// Index 聊天室首页，只显示模板内容
func (c *Controller) Index() {
	if c.Session.Contains("chat_name") {
		c.View.Assign("tplMain", "chat/include/chat.html")
	} else {
		c.View.Assign("tplMain", "chat/include/main.html")
	}
	c.View.Display("chat/index.html")
}

// SetName 设置响当当的名字
func (c *Controller) SetName() {
	name := c.Request.GetString("name")
	name = ghtml.Entities(name)
	c.Session.Set("chat_name_temp", name)
	if err := gvalid.Check(name, nameCheckRule, nameCheckMsg); err != nil {
		c.Session.Set("chat_name_error", err.String())
		c.Response.RedirectBack()
	} else if names.Contains(name) {
		c.Session.Set("chat_name_error", "用户昵称已被占用")
		c.Response.RedirectBack()
	} else {
		c.Session.Set("chat_name", name)
		c.Session.Remove("chat_name_temp")
		c.Session.Remove("chat_name_error")
		c.Response.RedirectTo("/chat")
	}
}

// WebSocket 接口
func (c *Controller) WebSocket() {
	msg := &Msg{}

	// 初始化WebSocket请求
	if ws, err := c.Request.WebSocket(); err == nil {
		c.ws = ws
	} else {
		glog.Error(err)
		return
	}

	name := c.Session.GetString("chat_name")
	if name == "" {
		name = c.Request.RemoteAddr
	}

	// 初始化时设置用户昵称为当前链接信息
	names.Add(name)
	users.Set(c.ws, name)

	// 初始化后向所有客户端发送上线消息
	c.writeUsers()

	for {
		// 阻塞读取WS数据
		_, msgByte, err := c.ws.ReadMessage()
		if err != nil {
			// 如果失败，那么表示断开，这里清除用户信息
			// 为简化演示，这里不实现失败重连机制
			names.Remove(name)
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
		msg.From = name

		// 日志记录
		glog.Cat("chat").Println(msg)

		// WS操作类型
		switch msg.Type {
		// 发送消息
		case "send":
			// 发送间隔检查
			intervalKey := fmt.Sprintf("%p", c.ws)
			if !cache.SetIfNotExist(intervalKey, struct{}{}, SendInterval) {
				c.write(Msg{"error", "您的消息发送得过于频繁，请休息下再重试", ""})
				continue
			}
			// 有消息时，群发消息
			if msg.Data != nil {
				if err = c.writeGroup(
					Msg{"send",
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
		for user := range m {
			user.(*ghttp.WebSocket).WriteMessage(ghttp.WS_MSG_TEXT, []byte(b))
		}
	})

	return nil
}

// 向客户端返回用户列表
func (c *Controller) writeUsers() error {
	array := garray.NewSortedStrArray()
	names.Iterator(func(v string) bool {
		array.Add(v)
		return true
	})
	if err := c.writeGroup(Msg{"list", array.Slice(), ""}); err != nil {
		return err
	}
	return nil
}
