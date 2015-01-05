# json [![Build Status](https://drone.io/github.com/gocontrib/json/status.png)](https://drone.io/github.com/gocontrib/json/latest)

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
