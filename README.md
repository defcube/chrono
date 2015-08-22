# chrono
Go library for passing time, such as waiting for a callback function to return true

You can see the [full documentation at godoc.org](https://godoc.org/github.com/defcube/chrono)

Example usage:

```
package main

import "github.com/defcube/chrono"

func main() {
    chrono.MustWaitFor(func() bool {
        // return True when the condition is satisfied, False if we should keep waiting
        // This function will be polled several times. The max timeout and sleep
        // between re-checking is configurable.
    })
}
```

