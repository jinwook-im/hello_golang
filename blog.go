// hello_golang/blog/blog.go
package hello_golang

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Blog struct {
	Articles []Article
}

func New() *Blog {
	return &Blog{}
}

func (b *Blog) SaveArticle(article Article) {
	b.Articles = append(b.Articles, article)
}

func (b *Blog) FetchAll() []Article {
	return b.Articles
}
