// curd包提供了对当前已配置数据库的快速CURD操作。
//
// 业务保留字段，可通过query参数提交:
// x_schema 操作的数据库
// x_table  操作的数据表
// x_where  原始SQL条件语句(可直接提交主键数值)
// x_order  排序语句, 例如: id desc
// x_group  分组语句, 例如: type
// x_page   分页语句(记录影响限制语句), 例如: 0, 100

package curd

import (
	"fmt"
	"github.com/gogf/gf-demos/library/response"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

type Controller struct{}

const (
	PageSizeDefault = 10
	PageSizeMax     = 100
)

// 构造函数
func (c *Controller) Init(r *ghttp.Request) {
	s := getSchema(r)
	if s == "" {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			response.Json(r, 1, fmt.Sprintf(`schema "%s" does not exist`, s))
			r.ExitAll()
		}
	}()
	g.DB(s)
}

// 析构函数
func (c *Controller) Shut(r *ghttp.Request) {
	// 在这里可以对当前控制器的所有路由函数返回值做拦截处理
}

// 提供对数据表的直接CURD访问，查询单条数据记录。
func (c *Controller) One(r *ghttp.Request) {
	table := r.GetRouterString("table")
	where, err := getWhere(r)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	one, err := g.DB().Table(table).FindOne(where)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "", one)
}

// 提供对数据表的直接CURD访问，查询多条数据记录。
func (c *Controller) All(r *ghttp.Request) {
	table := r.GetRouterString("table")
	order, err := getOrder(r)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	group, err := getGroup(r)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	where, err := getWhere(r)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	one, err := g.DB().Table(table).Page(getPage(r)).Order(order).Group(group).FindAll(where)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "", one)
}

// 保存记录
func (c *Controller) Save(r *ghttp.Request) {
	table := r.GetRouterString("table")
	data := r.GetMap()
	if j, _ := r.GetJson(); j != nil {
		for k, v := range j.ToMap() {
			data[k] = v
		}
	}
	where, err := getWhere(r)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	one, err := g.DB().Table(table).Data(data).WherePri(where).Page(getPage(r)).Filter().Save()
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "", one)
}

// 更新记录
func (c *Controller) Update(r *ghttp.Request) {
	table := r.GetRouterString("table")
	data := r.GetMap()
	if j, _ := r.GetJson(); j != nil {
		for k, v := range j.ToMap() {
			data[k] = v
		}
	}
	where, err := getWhere(r)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	one, err := g.DB().Table(table).Data(data).WherePri(where).Page(getPage(r)).Filter().Update()
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "", one)
}

// 删除记录
func (c *Controller) Delete(r *ghttp.Request) {
	table := r.GetRouterString("table")
	where, err := getWhere(r)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	one, err := g.DB().Table(table).WherePri(where).Page(getPage(r)).Delete()
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "", one)
}

// 获得请求的数据库名称
func getSchema(r *ghttp.Request) string {
	return r.GetQueryString("x_schema")
}

// 获得数据库查询条件，内部做防注入处理，并保证一个请求只执行一条SQL语句。
func getWhere(r *ghttp.Request) (interface{}, error) {
	s := r.GetQueryString("x_where")
	if s != "" {
		return s, nil
	}
	m := r.GetQueryMap()
	var err error
	for k := range m {
		// 业务保留字段，需过滤
		if len(k) > 2 && k[0:2] == "x_" {
			delete(m, k)
			continue
		}
		if err = validateSqlStr(k); err != nil {
			return nil, err
		}
	}
	return m, nil
}

// 获得分页条件，记录影响执行限制条件
func getPage(r *ghttp.Request) (page, size int) {
	s := r.GetQueryString("x_page")
	if s == "" {
		return 0, PageSizeDefault
	}
	array := gstr.SplitAndTrim(s, ",")
	if len(array) == 0 {
		return 0, PageSizeDefault
	}
	page = gconv.Int(array[0])
	if page < 0 {
		page = 0
	}
	size = PageSizeDefault
	if len(array) > 1 {
		size = gconv.Int(array[1])
		if size > PageSizeMax {
			size = PageSizeMax
		}
	}
	return gconv.Int(array[0]), size
}

// 获得排序条件
func getOrder(r *ghttp.Request) (string, error) {
	s := r.GetQueryString("x_order")
	if s == "" {
		return "", nil
	}
	if err := validateSqlStr(s); err != nil {
		return "", err
	}
	return s, nil
}

// 获得分组条件
func getGroup(r *ghttp.Request) (string, error) {
	s := r.GetQueryString("x_group")
	if s == "" {
		return "", nil
	}
	if err := validateSqlStr(s); err != nil {
		return "", err
	}
	return s, nil
}

// 检查给定的字符串是否安全(可提交到数据库执行)
func validateSqlStr(s string) error {
	match, _ := gregex.MatchString(`([^,'"+-\*%\.\/<>=\(\)\w\s])+`, s)
	if len(match) > 1 {
		return gerror.Newf(`invalid string "%s" in query condition`, match[1])
	}
	if gstr.Contains(s, "--") {
		return gerror.Newf(`invalid character '--' in query condition`)
	}
	return nil
}
