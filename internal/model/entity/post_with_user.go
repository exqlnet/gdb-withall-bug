package entity

import "github.com/gogf/gf/v2/frame/g"

type PostWithUser struct {
	g.Meta `orm:"table:post"`

	Post

	User User `json:"user" orm:"with:id=userId"`
}
