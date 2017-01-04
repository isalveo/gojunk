package trees

import(
  "golang.org/x/tour/tree"
  "fmt"
)

var (
  Tr1 *tree.Tree = tree.New(2)
  Tr2 *tree.Tree = tree.New(1)
)

func Walk(t *tree.Tree, c chan int){
  walk(t, c)
  close(c)
}

func walk(t *tree.Tree, c chan int){
  if t != nil {
    walk(t.Right, c)
    c <- t.Value
    walk(t.Left, c)
  }
}

func Same(t1, t2 *tree.Tree) bool{
  c1 := make(chan int)
  c2 := make(chan int)

  go Walk(t1, c1)
  go Walk(t2, c2)

  cc := 0

  for {
    select{
      case v1, ok := <- c1:
        if !ok {
          return true
        }
        cc++

        v2 := <- c2
        fmt.Println(v1)
        fmt.Println(v2)

        if v1 != v2 {
          return false
      }
    }
  }
  return true
}
