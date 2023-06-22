package v1

import (
	"github.com/gin-gonic/gin"
)

type Article struct{} //为了实现一系列结构体方法，定义了空结构体

func NewArticle() Article { //返回一个结构体对象，可以调用结构体方法
	return Article{}
}

//Article路由对应的处理方法（Handler），以结构体方法的形式实现

func (a Article) Get(c *gin.Context)    {}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
