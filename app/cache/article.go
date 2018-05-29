package appCache

import (
	"fmt"
	"mugg/guapin/app/service"
	"mugg/guapin/model"

	"github.com/patrickmn/go-cache"
)

// Goods is
type articles struct{}

// GetAll is
func (b articles) GetAll() (interface{}, error) {

	var (
		data []model.Article
		err  error
	)

	articles, has := Caches.Get("articles")

	if has {
		fmt.Println("has", articles)
		return articles, err
	}

	data, err = service.articles{}.GetAll()
	fmt.Println("\b")
	fmt.Println(err)
	if err == nil {
		Caches.Set("articles", data, cache.DefaultExpiration)
	}

	return data, err
}
