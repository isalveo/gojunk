package mapUtil

import(
  "testing"
)

func TestMapIt(t *testing.T) {
  s := []E{1,2,3}
  var r []E

  r = MapIt(sampleFunc, s)

  ret := make([]int, 3)

  sample := [3]int{1,4,9}

  for i, v := range r {
    ret[i] = v.(int)
  }

  var retA [3]int
  copy(retA[0:3], ret[:])

  if retA != sample {
    t.Fail()
  }
}

func sampleFunc(x E) E {
  var ret E
  switch x.(type) {
  case int:
    i := x.(int)
    ret = i * i
  case string:
    st := x.(string)
    ret = st + st
  default:
    ret = "nothing has been done!"
  }
  return ret
}