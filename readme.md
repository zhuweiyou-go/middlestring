# middlestring

```bash
go get -u github.com/zhuweiyou-go/middlestring
```

```go
package main

import (
	"github.com/zhuweiyou-go/middlestring"
)

func main() {
	// normal
	middlestring.Get("(hello)(world)", "(", ")") // "hello"
	
	// more
	middlestring.Get("(hello)(world)", "(", ")", true) // "hello)(world"
	
	// not found
	middlestring.Get("(hello)(world)", "[", "]") // ""
}
```
