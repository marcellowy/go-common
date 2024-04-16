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
#### v0.0.6 2024/04/16
- add thread safe slice
- add thread safe map/slice unit test case
#### v0.0.5
- add auto register router for gin framework, see router/README.md
#### v0.0.4
- add GormLogger help write gorm log to zap log
- bugfix
#### v0.0.3
- add SliceTrimSame remove slice same element
- add SliceRemove remove slice specify element
- RemoveSameFromStringSlice Departed，use SliceTrimSame replace
- FormatTime use genericity 

#### v0.0.2
- add Close，close io.Closer

#### v0.0.1
- init version

### LICENSE
MIT
