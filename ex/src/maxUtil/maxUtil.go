package maxUtil

type E interface{}

func Max(s []E) E {
  var max E = 0
  var maxF float64

  for _, v := range s{
    switch v.(type){
    case int:
      vv := v.(int)
      if vv > max.(int){
        max = v
      }
    case float64:
      vv := v.(float64)
      if vv > maxF{
        max = v
        maxF = vv
      }
    }
  }

  return max
}

func Max2(s []E) E {
  var max E = s[0]
  for i := 1; i < len(s); i++{
    if less(max, s[i]){
      max = s[i]
    }
  }
  return max
}

func less(x, y interface{}) bool {
  switch x.(type){
  case int:
    if _, ok := y.(int); ok {
      return x.(int) < y.(int)
    }
  case float32:
    if _, ok := y.(float32); ok {
      return x.(float32) < y.(float32)
    }
  }
  return false
}
