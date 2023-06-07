### Marcello's Public Go Library

### 如何使用
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
	
	fmt.Println(tools.Md5("test")) // output: 098f6bcd4621d373cade4e832627b4f6
	
	var path = "path/file/test.txt"
	hash, err := tools.Md5File(path)
	if err!=nil{
		// if path not exists or not file
		return
    }
	fmt.Println(hash) // output: file md5 string
}
```

### 版本记录
#### v0.0.3(pre-release)
1. 新增 SliceTrimSame 移除slice中的相同元素
2. 新增 SliceRemove 移除slice中指定的元素
3. 标记 RemoveSameFromStringSlice 为废弃，使用 SliceTrimSame 代替
4. FormatTime 修改为泛型

#### v0.0.2
1. 新增Close方法，关闭 io.Closer，使ide不提示

#### v0.0.1
1. 正式从beta变成可用版本

### LICENSE
MIT