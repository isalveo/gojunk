package stack

import "testing"
import "fmt"

func TestIn(t *testing.T) {
  s := new(Stack)

  s.In(111)

  if s.Data[0] != 111 {
    t.Fail()
  }
}

func TestOut(t *testing.T) {
  s := new(Stack)

  s.In(555)

  if s.Out() != 555 {
    t.Fail()
  }
}

func ExamplePrint() {
  s := new(Stack)

  s.In(555); s.In(444); s.In(333)
  fmt.Printf("%v", s)
  //Output:
  //[0:555][1:444][2:333]
}