package cache_test

import (
	"context"
	"fmt"
	"time"

	pcache "github.com/patrickmn/go-cache"

	cache "github.com/otternq/patrickmn-go-cache"
)

func Example() {
	var (
		c            = pcache.New(time.Minute, time.Minute)
		cacheWrapper = cache.Wrap(c, "default")

		found bool
		val   interface{}
	)

	cache.RegisterAllViews()

	cacheWrapper.Set(context.TODO(), "key", "value", time.Second)

	val, found = cacheWrapper.Get(context.TODO(), "key")

	fmt.Println("Found:", found)
	fmt.Println("Val:", val)

	// Output:
	// Found: true
	// Val: value
}
