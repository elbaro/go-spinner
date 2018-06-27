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
    s.Stop()
}
```

[![asciicast](https://asciinema.org/a/MIvlqqIpgtwcZDsuWT1QQs79k.png)](https://asciinema.org/a/MIvlqqIpgtwcZDsuWT1QQs79k)


No fancy config. simple interface.

### TODO
- [ ] multiple spinners
