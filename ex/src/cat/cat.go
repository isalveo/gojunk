package cat

import(
  "io"
  "os"
  "fmt"
  "bufio"
  "flag"
)

var NumberFlag = flag.Bool("n", true, "number each line")

func Cat(r *bufio.Reader) {
  i := 1
  for {
    buf, e := r.ReadBytes('\n')

    if e == io.EOF {
      break
    }

    if *NumberFlag {
      fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
      i++
    } else {
      fmt.Fprintf(os.Stdout, "%s",buf)
    }
  }
  return
}