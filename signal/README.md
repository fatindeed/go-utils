# Signal Dispatcher

This library provides:

- Graceful shutdown
- Simple integration like [pcntl](https://www.php.net/manual/book.pcntl.php)
- Customize signal

## Installation

Use go get to install the latest version of the library.

```sh
go get github.com/fatindeed/go-utils
```

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/fatindeed/go-utils/signal"
)

func main() {
	defer signal.Stop()
	for i := 0; i < 10; i++ {
		err := signal.Dispatch()
		if err != nil {
			fmt.Printf("%s at %s", err, time.Now().String())
			break
		}
		fmt.Println(time.Now().String())
		time.Sleep(time.Second * 2)
	}
}
```

Output will like this

```sh
$ go run main.go
2022-08-31 12:54:36.55567 +0800 CST m=+0.001286001
2022-08-31 12:54:38.562498 +0800 CST m=+2.008137001
2022-08-31 12:54:40.564897 +0800 CST m=+4.010528001
^C
signal: interrupt at 2022-08-31 12:54:42.567635 +0800 CST m=+6.013267001%
```

This library only catch `SIGINT` and `SIGTERM` by default.

If you want to customize signals, add `signal.Listen(...os.Signal)` in your code.

```go
signal.Listen(syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)
defer signal.Stop()
```
