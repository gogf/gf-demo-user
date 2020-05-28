// curd包提供了对当前已配置数据库的快速CURD的API操作。
//
// 业务保留字段，可通过query参数提交:
// x_schema 操作的数据库
// x_where  原始SQL条件语句(可直接提交主键数值)
// x_order  排序语句, 例如: id desc
// x_group  分组语句, 例如: type
// x_page   分页语句(记录影响限制语句), 例如: 1,100

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
	PageSizeDefault = 10  // 查询数据时分页默认条数
	PageSizeMax     = 100 // 查询数据时分页最大条数
	AffectSizeMax   = 100 // 更新/删除操作时最大的受影响行数
)

// 请求构造函数
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
	// 在这里可以对当前控制器的所有路由函数返回值做拦截处理。
	// 交给你啦..
}

// @summary 查询单条数据记录
// @tags    快速CURD
// @produce json
// @param   table    path  string  true "操作的数据表"
// @param   x_schema query string false "操作的数据库"
// @param   x_where  query string false "原始SQL条件语句(可直接提交主键数值)"
// @param   x_order  query string false "排序语句, 例如: `id desc`"
// @param   x_group  query string false "分组语句, 例如: `type`"
// @param   x_page   query string false "分页语句(记录影响限制语句), 例如: `1,100`"
// @router  /curd/{table}/one [GET]
// @success 200 {object} response.JsonResponse "查询结果"
func (c *Controller) One(r *ghttp.Request) {
	table := r.GetRouterString("table")
	where, err := getWhere(r)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	one, err := g.DB().Schema(getSchema(r)).Table(table).FindOne(where)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "", one)
}

// @summary 查询多条数据记录
// @tags    快速CURD
// @produce json
// @param   table    path  string  true "操作的数据表"
// @param   x_schema query string false "操作的数据库"
// @param   x_where  query string false "原始SQL条件语句(可直接提交主键数值)"
// @param   x_order  query string false "排序语句, 例如: `id desc`"
// @param   x_group  query string false "分组语句, 例如: `type`"
// @param   x_page   query string false "分页语句(记录影响限制语句), 例如: `1,100`"
// @router  /curd/{table}/all [GET]
// @success 200 {object} response.JsonResponse "查询结果"
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
	one, err := g.DB().Schema(getSchema(r)).Table(table).Page(getPage(r)).Order(order).Group(group).FindAll(where)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "", one)
}

// @summary 保存数据记录
// @description 注意保存的数据通过表单提交，由于提交的数据字段不固定，因此这里没有写字段说明，并且无法通过`swagger`测试。
// @tags    快速CURD
// @produce json
// @param   table    path  string true  "操作的数据表"
// @param   x_schema query string false "操作的数据库"
// @router  /curd/{table}/save [POST]
// @success 200 {object} response.JsonResponse "执行结果"
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
	one, err := g.DB().Schema(getSchema(r)).Table(table).Data(data).WherePri(where).Limit(AffectSizeMax).Filter().Save()
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "", one)
}

// @summary 更新数据记录
// @description 注意修改的数据通过表单提交，由于提交的数据字段不固定，因此这里没有写字段说明，并且无法通过`swagger`测试。
// @tags    快速CURD
// @produce json
// @param   table    path  string true  "操作的数据表"
// @param   x_schema query string false "操作的数据库"
// @param   x_where  query string false "原始SQL条件语句(可直接提交主键数值)"
// @param   x_page   query string false "分页语句(记录影响限制语句), 例如: `1,100`"
// @router  /curd/{table}/update [POST]
// @success 200 {object} response.JsonResponse "执行结果"
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
	one, err := g.DB().Schema(getSchema(r)).Table(table).Data(data).WherePri(where).Limit(AffectSizeMax).Filter().Update()
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "", one)
}

// @summary 删除数据记录
// @tags    快速CURD
// @produce json
// @param   table    path  string true  "操作的数据表"
// @param   x_schema query string false "操作的数据库"
// @param   x_where  query string true  "原始SQL条件语句(可直接提交主键数值)"
// @param   x_page   query string false "分页语句(记录影响限制语句), 例如: `1,100`"
// @router  /curd/{table}/delete [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (c *Controller) Delete(r *ghttp.Request) {
	table := r.GetRouterString("table")
	where, err := getWhere(r)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if where == "" {
		response.JsonExit(r, 1, "执行删除时Where条件不能为空")
	}
	one, err := g.DB().Schema(getSchema(r)).Table(table).WherePri(where).Limit(AffectSizeMax).Delete()
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

// 获得分页条件
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
	// 是否包含不允许的字符
	match, _ := gregex.MatchString(`([^,'"+-\*%\.\/<>=\(\)\w\s])+`, s)
	if len(match) > 1 {
		return gerror.Newf(`invalid string "%s" in query condition`, match[1])
	}
	// 不能存在注释语句
	if gstr.Contains(s, "--") {
		return gerror.Newf(`invalid character '--' in query condition`)
	}
	return nil
}
