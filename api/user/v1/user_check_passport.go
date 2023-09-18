package v1

import "github.com/gogf/gf/v2/frame/g"

type CheckPassportReq struct {
	g.Meta   `path:"/user/check-passport" method:"post" tags:"UserService" summary:"Check passport available"`
	Passport string `v:"required"`
}
type CheckPassportRes struct{}
