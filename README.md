# json

Untyped JSON API

## example

```go
package main

import "github.com/gocontrib/json"
import "fmt"

func main() {
  o, _ := json.ParseObject(`{"user": "bob"}`)
  fmt.Printf("%s\n", o.JSON())
}
```
