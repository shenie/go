package main

import (
    "log"
    "time"
    "os/exec"
    "os"
    "fmt"
)

func main() {
  if len(os.Args) < 2 {
    log.Fatal("No command")
  }
  for true {
    time.Sleep(5 * time.Second)
    out, err := exec.Command(os.Args[1], os.Args[2:]...).Output()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("> %s\n", time.Now())
    fmt.Printf("%s", out)
  }
}
