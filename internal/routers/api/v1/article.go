package v1

import (
	"blog-server/global"
	"blog-server/internal/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	model.BlogArticle
}

func NewArticle() Article {
	return Article{}
}

func (t Article) Get(c *gin.Context)  {}
func (t Article) List(c *gin.Context) {}
func (t Article) Create(c *gin.Context) {
	db := global.DBEngine
	a := NewArticle()
	db.Limit(1).Find(&a)
	fmt.Println(a)
	c.JSON(http.StatusOK,
		map[string]string{"res": " ok"})
}
func (t Article) Update(c *gin.Context) {}
func (t Article) Delete(c *gin.Context) {}
