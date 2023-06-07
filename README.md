# Convert a JSON format to struct
```
go get github.com/nlanatta/json_to_struct_converter
```

```go
package main

import (
	"fmt"
	"github.com/nlanatta/json_to_struct_converter"
)

// run 
func main() {
	converter := JsonToStruct()
	converter.Run(v.json)
	got := converter.Clipboard()

	fmt.Println(got)
}
```