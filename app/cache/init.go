package appCache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Caches is
var Caches *cache.Cache

func init() {
	Caches = cache.New(5*time.Minute, 10*time.Minute)
}
