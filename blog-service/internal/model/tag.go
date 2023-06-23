package model

import "github.com/consult98/simple-blog-go/pkg/app"

type Tag struct { //文章标签model，与tag表字段对应
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct { //用于Swagger展示响应结果
	List  []*Tag
	Pager *app.Pager
}
