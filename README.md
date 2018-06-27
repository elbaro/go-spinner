A spinner in Golang.

```
package main

import (
    "time"
    "github.com/elbaro/go-spinner"
)

func main() {
    s := spinner.New("Downloading ..")
    time.Sleep(3 * time.Second)
    s.Done()
    // s.Fail()
}
```

[![asciicast](https://asciinema.org/a/MIvlqqIpgtwcZDsuWT1QQs79k.png)](https://asciinema.org/a/MIvlqqIpgtwcZDsuWT1QQs79k)

- leave

```go
// leave = true
s.Done()
s.Fail()

// leave = false
s.DoneClean()
s.FailClean()
```

- progress

```go
s := spinner.NewProgress("Fetching")
s.Add()
s.Add()
s.DoneClean()
s.Fail("second asset")
s.Stop()
```

- group

show multiple spinners at the same time.

```go
// not implemented.
// use spinner.NewProgress() for multi-threaded env.
```
