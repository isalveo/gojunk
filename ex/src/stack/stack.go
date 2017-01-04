package stack

import "strconv"

type Stack struct{
  i int
  Data [10]int
}

func (s *Stack) In(v int) {
  if s.i > 9 {
    return
  }
  s.Data[s.i] = v
  s.i++
}


func (s *Stack) Out() (ret int) {
 s.i--
 if s.i < 0 {
  s.i = 0
  return
 }
 ret = s.Data[s.i]
 return
}

var Push = (*Stack).In
var Pop = (*Stack).Out

func (s Stack) String() string{
  var str string
  for j := 0; j < s.i; j++ {
    str += "[" + strconv.Itoa(j) + ":" + strconv.Itoa(s.Data[j]) + "]"
  }
  return str
}

