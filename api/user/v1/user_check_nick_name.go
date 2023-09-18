package v1

import "github.com/gogf/gf/v2/frame/g"

type CheckNickNameReq struct {
	g.Meta   `path:"/user/check-nick-name" method:"post" tags:"UserService" summary:"Check nickname available"`
	Nickname string `v:"required"`
}
type CheckNickNameRes struct{}
