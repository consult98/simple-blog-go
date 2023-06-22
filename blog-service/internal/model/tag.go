package model

type Tag struct { //文章标签model，与tag表字段对应
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
