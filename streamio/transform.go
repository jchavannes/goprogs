package main

import "os"
import "bufio"
import "strings"
import "encoding/csv"

func main() {
  reader := bufio.NewReader(os.Stdin)
  writer := csv.NewWriter(os.Stdout)

  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      break
    }
    columns := strings.Split(line, "\t")
    for i, val := range columns {
      columns[i] = strings.Trim(val, "\n")
    }
    writer.Write(columns)
    writer.Flush()
  }

}
