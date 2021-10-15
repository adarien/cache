# EXAMPLE

```go
package main

import (
	"fmt"
	"github.com/adarien/cache"
)

func main() {
	cache := cache.New()
	cache.Set("userID", 42)
	cache.Set("userName", "adarien")
	cache.Set("userStatus", true)
	fmt.Println(cache)

	userID := cache.Get("userID")
	fmt.Println(userID)

	cache.Delete("userID")
	userID = cache.Get("userID")
	fmt.Println(userID)
	fmt.Println(cache)
}
```
