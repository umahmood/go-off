/*
Package off a Go implementation of off. A lightweight format for small
configuration files. See http://bindh3x.io/off.

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
            fmt.Println("port", val)
        }

        a, err := config.Array("my_array")
        if err != nil {
            // handle error
        }
        x := a[0].(int)
        fmt.Println("element at index 0", x)
        y := a[3].(string)
        fmt.Println("element at index 3", y)
        z := a[6].(string)
        fmt.Println("element at index 6", z)
    }
*/
package off
