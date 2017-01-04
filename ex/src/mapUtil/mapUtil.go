package mapUtil

type E interface{}

func MapIt(f func(E) E, s []E) []E {
  ret := make([]E, len(s))
  for i, v := range s {
    ret[i] = f(v)
  }
  return ret
}
