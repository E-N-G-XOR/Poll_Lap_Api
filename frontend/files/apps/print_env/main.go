package main

import (
  "os"
  "fmt"
)

func main() {
  for _, pair := range os.Environ() {
    fmt.Println(pair)
  }
}