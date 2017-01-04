package utils

import "os/exec"
import (
  "net"
  "net/http"
  "io/ioutil"
  "fmt"
  "sort"
  "strconv"
  "strings"
  "os"
  "bufio"
)


func Fibonacci(i, q chan int) {
  x, y := 1, 1

  for{
    select{
      case <-i:
        x, y = y, x + y
        println(y)
      case <- q:
        return
    }
  }
}


func Exec(s, opts string) ([]byte, error) {
  cmd := exec.Command(s, opts)
  //cmd.Run()
  buf, err := cmd.Output()
  return buf, err
}

func Get(url string){
  res, err := http.Get(url)

  if err != nil {
    fmt.Printf("Error: %v", err.Error())
    return
  }

  b, err := ioutil.ReadAll(res.Body)
  defer res.Body.Close()

  if err == nil {
    fmt.Printf("Response: %v\n", string(b))
  }
}

func GetChildProcesses() {
  cmd := exec.Command("ps", "e", "-opid,ppid,comm")
  output, _ := cmd.Output()

  child := make(map[int][]int)

  for i, p := range strings.Split(string(output), "/n"){
    if i == 0 { continue }
    // if len(p) == 0 { continue }

    f := strings.Fields(p)
    //parent process
    fpp, _ := strconv.Atoi(f[1])
          fmt.Printf("%v", f)
    // child process
    fcp, _ := strconv.Atoi(f[0])
    child[fpp] = append(child[fpp], fcp)
  }
  schild := make([]int, len(child))
  i := 0
  for k, _ := range child { schild[i] = k; i++ }
  sort.Ints(schild)

  for _, ppid := range schild {
    fmt.Printf("Pid %d has %d child", ppid, len(child[ppid]))
    if len(child[ppid]) == 1 {
      fmt.Printf(": %v\n", child[ppid])
      continue
    }
    fmt.Printf("ren: %v", child[ppid])
  }
}

func WC() {
  var chars, words, lines int

  r := bufio.NewReader(os.Stdin)
  println("Reading...")

  for {
    switch s, ok := r.ReadString('\n'); true{
      case ok != nil:
        fmt.Printf("%d, %d, %d", chars, words, lines)
        return
      case include(strings.Fields(s), "quit"):
        fmt.Printf("%d, %d, %d\n, %s", chars, words, lines)
        return
      default:
        chars += len(s)
        words += len(strings.Fields(s))
        lines++
    }
  }
}

func Uniq() {
  r := bufio.NewReader(os.Stdin)
  s, _ := r.ReadString('\n')

  chars :=  strings.Fields(s)

  prev := chars[0]

  println("-------")
  println(prev)
  for _, el := range chars[1:] {
    if el != prev {
      println(el)
      prev = el
    }
  }
}

func Echo(c net.Conn){
  defer c.Close()

  line, err := bufio.NewReader(c).ReadString('\n')

  if err != nil {
    fmt.Printf("Failed to read %s\n", err.Error())
    return
  }

  _, err = c.Write([]byte(line))

  if err != nil {
    fmt.Println("Failed to write %s\n", err)
  }

}

func include(s []string, el string) bool{
  for _, v := range s {
    if v == el{
      return true
    }
  }
  return false
}
