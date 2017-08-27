# Off

A Go implementation of Off, which is a lightweight format for small configuration 
files. See http://bindh3x.io/off.

# Installation

> $ go get github.com/umahmood/go-off 

# Usage

Assuming the following Off configuration file 'test.off':
```
;; HTTP Server configuration file
;; Author: bindh3x

;; server
host 0.0.0.0
port 8080
banner "HTTPServer 1.0"

;; privileges
run_as_root true
user root
group wheel

;; array with int and string elements
my_array {1|2|3|Hello, World!|https://www.kernel.org/|200|<p>Hello</p>}

;; admin contact information
admin_email admin@example.com
admin_phone +12345678
```
Load the config. file:
```
package main

import (
    "fmt"
    "os"

    off "github.com/umahmood/go-off"
)

func main() {
    file, err := os.Open("test.off")
    if err != nil {
        // handle error
    }
    defer file.Close()

    config, err := off.LoadConfig(file)
    if err != nil {
        // handle error
    }

    if val, err := config.String("banner"); err == nil {
        fmt.Println("banner:", val)
    }

    if val, err := config.Bool("run_as_root"); err == nil {
        fmt.Println("run_as_root:", val)
    }

    if val, err := config.Int("port"); err == nil {
        fmt.Println("port:", val)
    }

    a, err := config.Array("my_array")
    if err != nil {
        // handle error
    }
    x := a[0].(int)
    fmt.Println("element at index 0:", x)
    y := a[3].(string)
    fmt.Println("element at index 3:", y)
    z := a[6].(string)
    fmt.Println("element at index 6:", z)
}
```
Output:
```
banner: "HTTPServer 1.0"
run_as_root: true
port 8080
element at index 0: 1
element at index 3: Hello, World!
element at index 6: <p>Hello</p>
```

# Documentation

> http://godoc.org/github.com/umahmood/go-off

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
