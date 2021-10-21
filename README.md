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

	userID, err := cache.Get("userID")
	if err != nil {
	    fmt.Println(err)
	} else {
	    fmt.Println(userID)
	}

	cache.Delete("userID")
	
	userID, err = cache.Get("userID")
	if err != nil {
	    fmt.Println(err)
	} else {
	    fmt.Println(userID)
	}
	
	fmt.Println(cache)
}
```
