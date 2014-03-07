package main

import "os"
import "bufio"
import "fmt"
import "time"

func main() {
  reader := bufio.NewReader(os.Stdin)

  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      break
    }
    time.Sleep(1 * time.Second)
    fmt.Printf(line)
  }
}
