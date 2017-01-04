package even

func Even(x int) bool {
return x % 2 == 0
}

func Odd() func(int) bool{
return func (x int) bool {
 return x % 2 != 0
}
}
