### Marcello's common Golang Library
![workflow](https://github.com/marcellowy/go-common/actions/workflows/go.yml/badge.svg?branch=main)
![Release](https://badgen.net/github/release/marcellowy/go-common)
![LAST_COMMIT](https://badgen.net/github/last-commit/marcellowy/go-common)
![MIT](https://badgen.net/github/license/micromatch/micromatch)

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

#### Copy
```shell
    package main
    
    import (
        "fmt"
        "github.com/marcellowy/go-common/tools"
    )
    
    func main() {
        _ = os.MkdirAll("./test", os.ModePerm)
        // copy parent "foo" and all subdirectory and "foo.txt" file to "./test" directory
        if err := tools.Copy("./test", []string{"../foo", "../foo.txt"});err!=nil{         
            // print log
        }
        
        // copy file
        // copy "b.txt" to "../../a.txt"
        if err := tools.CopyFile("../../a.txt", "./b.txt"); err !=nil {
            // print log
        }
    }     
```

### version
#### v0.0.14 2024/05/07
- added ftp download/upload/delete
- added download file from URL, support large file
- added func ReCreateDirectory removes the directory at the given path and creates a new empty directory at the same path.
- added func RemoveLastSeparator removes the last separator from the given path string.

#### v0.0.13 2024/04/29
- added func Copy copies the contents of the file at the given source path to the destination path. 
- added func DirHasPrefix checks if the directory path 's' has a prefix 'prefix'.
- added func CopyFile copies the contents of the file at the given source path to the destination path.
- added func Zip/Unzip zip or unzip file/directory
#### v0.0.12 2024/04/16
- added thread safe slice
- added thread safe map/slice unit test case
#### v0.0.10 2024/01/5
- bugfix
#### v0.0.9 2023/11/04
- remove old version
#### v0.0.9 2023/11/01
- fix parse time
#### v0.0.8 
- added gogf support
#### v0.0.5
- added auto register router for gin framework, see router/README.md
#### v0.0.4
- added GormLogger help write gorm log to zap log
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
