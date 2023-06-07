### Marcello's Public Go Library

### How to use
```shell
go get github.com/marcellowy/go-common
```

### Example
#### md5 and md5 file

```go
package main

import (
	"fmt"
	"github.com/marcellowy/go-common/tools"
)

func main() {
	
	// md5
	fmt.Println(tools.Md5("test")) // output: 098f6bcd4621d373cade4e832627b4f6
	
	// file md5
	var path = "path/file/test.txt"
	hash, err := tools.Md5File(path)
	if err!=nil{
		// if path not exists or not file
		return
    }
	fmt.Println(hash) // output: hash string
}
```

### version
#### v0.0.3(pre-release)
1. add SliceTrimSame remove slice same element
2. add SliceRemove remove slice specify element
3. RemoveSameFromStringSlice Departed，use SliceTrimSame replace
4. FormatTime use genericity 

#### v0.0.2
1. add Close，close io.Closer

#### v0.0.1
1. init version

### LICENSE
MIT