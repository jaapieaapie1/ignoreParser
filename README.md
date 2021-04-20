## Ignore parser
______
A go module used for parsing ignore files like .gitignore (and using them)

<br><br>
Install this package:
`go get github.com/jaapieaapie1/ignoreParser`
<br><br>


### Example

```go
package main

import (
	"os"
	"github.com/jaapieaapie1/ignoreParser"
)

func main() {
	file, _ := os.Open(".gitignore")
	
	ignore, _ := ignoreParser.ParseIgnoreFile(file)

	if ignore.ShouldIgnore("hello.txt") {
            println("Don't use hello.txt")
	}
}
```
