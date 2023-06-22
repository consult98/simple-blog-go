package v1

import (
	"github.com/gin-gonic/gin"
)

type Tag struct{} //为了实现一系列结构体方法，定义了空结构体

func NewTag() Tag { //返回一个结构体对象，可以调用结构体方法
	return Tag{}
}

//Tag路由对应的处理方法 (Handler)，以结构体方法的形式实现

func (t Tag) Get(c *gin.Context)    {}
func (t Tag) List(c *gin.Context)   {}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
