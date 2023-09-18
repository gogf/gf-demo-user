package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
)

type ProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"UserService" summary:"Get the profile of current user"`
}
type ProfileRes struct {
	*entity.User
}
