package even

import (
"testing"
"fmt"
)

func TestEven(t *testing.T) {
  if Even(2) != true {
    t.Log("Should be true")
    t.Fail()
  }
}

func TestOdd(t *testing.T) {
  p := Odd()
  if p(1) != true {
    t.Log("Boom")
    t.Fail()
  }
}

func ExampleEven(){
  if Even(2) {
    fmt.Println("Is even")
  }
  //Output:
  //Is even
}