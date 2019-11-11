# github.com/gotoolkit/config Go項目參數配置


## Install

```console
go get -u github.com/gotoolkit/config
```

## Example


```console
mkdir -p config
echo '{"hello":"world"}' > config/config.json
```


```go
package main

import (
    "log"

    "github.com/gotoolkit/config"
)

func main() {
    err := config.Setup(config.WithWatchEnable(true))
	if err != nil {
		log.Fatalln(err)
    }
    
    log.Println(config.Get("hello"))
}
```