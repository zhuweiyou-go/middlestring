# getmiddlestring

```bash
go get -u github.com/zhuweiyou-go/getmiddlestring
```

```go
package main

import (
	"github.com/zhuweiyou-go/getmiddlestring"
)

func main() {
	// normal
	getmiddlestring.Get("(hello)(world)", "(", ")") // "hello"
	
	// more
	getmiddlestring.Get("(hello)(world)", "(", ")", true) // "hello)(world"
	
	// not found
	getmiddlestring.Get("(hello)(world)", "[", "]") // ""
}
```
