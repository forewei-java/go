package v1

import (
	"github.com/forewei-java/go/blog-service/pkg/app"
	"github.com/forewei-java/go/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

// article.go
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
