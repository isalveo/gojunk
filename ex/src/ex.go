package main

import (
	"fmt"
	// "stack";
	"mapUtil"
	"maxUtil"
	"time"
  // "even"
	// "trees"
	// "crawler"
	// "utils"
	// "net"
	// "cruncher"
	// "flag"
	// "listUtil"
	// "cat"
	// "io"
	// "os"
	// "bufio"
	// "flag"
	// "bufio"
	// "os"
)

//import "math"
// import "unicode/utf8"
import "strconv"

func main() {
	// var counter int
	// for i :=1; i<=10; i++{
	//  counter += i
	//   fmt.Println("Sum:", counter)
	//  }
	/*  counter := 0
	Here:
	  if counter >= 10 {
	    return
	  }
	  counter++
	  fmt.Println("Sum:", counter)
	  goto Here


	var a [10]int
	for i:=0; i < 10; i++ {
	    a[i] = i
	  }
	 fmt.Println("Array:", a)
	}
	*/

	//FIttBuzz

	/*
	   for i:=1; i<=100; i++{
	     if (i % 3) == 0 && (i % 5) == 0 {
	       fmt.Println("FittBuzz")
	     } else {
	      if (i % 3)==0 {
	        fmt.Println("Fitt")
	      } else {
	        if (i % 5)==0 {
	         fmt.Println("Buzz")
	        } else {
	           println(i)
	        }
	       }
	      }
	    }

	     switch {
	      case (i % 3)==0 && (i % 5)==0:
	         println("FizzBuzz")
	       case (i % 3) == 0:
	         fmt.Println("Fizz")
	       case (i % 5) == 0:
	          println("Buzz")
	       default:
	          fmt.Println(i)
	     }
	   }

	   //AAAAA
	   a := "A"
	   for i:=0; i < 100; i++ {
	     a += "A"
	     fmt.Println(a)
	   }
	*/

	//STRING
	// s := "asSASA ddd dsjkdsjs dk"
	// ss := []byte(s)
	// rr := []rune(s)
	// copy(rr[4:4+3], []rune("XXX"))
	// fmt.Printf("Number of symbols %d", string(rr))

	//REVERSE STRING
	//  s, i := "foobar", 0
	//  rr := []rune(s)
	//  n := len(rr)
	//  rev := make([]rune, n)
	//  for i < n {
	//    rev[n-i-1] = rr[i]
	//    i++
	//  }
	// fmt.Println(string(rev))
	/*
	   s := "foobar"
	   ss := []rune(s)

	   for i,j := 0, len(ss) -1; i < j; i,j = i + 1, j-1{
	     ss[i], ss[j] = ss[j], ss[i]
	   }

	   fmt.Println(string(ss))

	   s := []float64{0.1, 0.2, 0.3}
	   sum := 0.0

	   for _, v := range s {
	     sum += v
	   }

	   fmt.Printf("Average=%f", sum/float64(len(s)))
	*/

	// FUNCTIONS
	/*
	   b := func(arg ...string){
	     for _, v := range arg{
	       fmt.Println("B ", v)
	     }
	   }

	   a := func(arg ...string){
	     b(arg...)
	   }

	   a("1", "bla")
	*/

	/////Map

	// defer func() {
	//   println("1END")
	// }()

	// defer func() {
	//   println("2END")
	// }()

	// m := map[int]func() int{
	//   1: func() int { return 10 },
	//   2: func() int { return 20 },
	// }

	// fmt.Println(m[1])

	//CALLBACK

	// b := func (x int, f func(int)) {
	//   f(x)
	// }
	// b(3333, printit)

	//CATCHING ERRORS
	// x := throwsError(panicit)

	// fmt.Println(x)
	//a("xxx")

	// x := []float64{0.1,0.2,0.3,.4}

	//fmt.Println(asc(1,2))

	//STACK
	// s := new(stack)

	// fmt.Printf("Value %v", s)
	// for _, v := range Map(strSquare, "a", "b", "c"){
	//   fmt.Printf("%s\n", v)
	// }

	// sss := []int{1,2,33,-4,1,55}
	// BubbleSort(sss)
	// fmt.Println(sss)

	// p := PlusTwo()
	// fmt.Printf("%v\n", p(3))
	// p := evennn.Odd()

	// fmt.Printf("Odd %v", p(4))

	/*
	   s := new(stack.Stack)

	   s.In(666)
	   s.In(555)
	   s.In(444)

	   fmt.Printf("%v", s)
	*/

	//CALCULATOR
	// var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	// var st = new(stack.Stack)

	// for {
	//   s, err := reader.ReadString('\n')
	//   var token string

	//   if err != nil {
	//    return
	//   }

	//   for _, v := range s {
	//     switch {
	//     case v >= '0' && v <= '9':
	//     token = token + string(v)
	//     case v == ' ':
	//     r, _ := strconv.Atoi(token)
	//     stack.Push(st, r)
	//     token = " "
	//     case v == '+':
	//     fmt.Printf("%d", stack.Pop(st) + stack.Pop(st))
	//     case v == '*':
	//     fmt.Printf("%d",stack.Pop(st)*stack.Pop(st))
	//     }
	//   }
	// }

	// var i int = 999
	// var p *int

	// p = &i

	// *p = 123
	// *p ++
	// fmt.Printf("%d", i)

	// type boBa struct {
	//   i int
	//   stack.Stack
	// }

	// type bOM boBa

	// // v := new(boBa)
	// v := new(boBa)
	// w := new(bOM)

	// v.In(666)
	// w.In(555)

	// fmt.Printf("===%v\n", v.Out())
	// fmt.Printf("----%v\n", w.Out())

	// s := []int{1,2,3}
	// fmt.Printf("%v", mapUtil.MapIt(intSquare, s))

	// xxx := []mapUtil.E{1,2,3}
	// xxx = []mapUtil.E{"1","2","3"}
	// fmt.Println(mapUtil.MapIt(DoubelIt, xxx))

	// l := listUtil.ListT

	// println(l == nil)

	// l.Push(1)
	// l.Push(2)
	// l.Push(4)
	// l.Push(55)

	// listUtil.Print(l)

	// for v, err := l.Pop(); err == nil; v, err = l.Pop(){
	//   fmt.Printf("%v\n", v)
	// }

	// flag.Parse()

	// if flag.NArg() == 0 {
	//   println("Nil args")
	//   cat.Cat(bufio.NewReader(os.Stdin))
	// }

	// for i := 0; i < flag.NArg(); i ++ {
	//   f, e := os.Open(flag.Arg(i))
	//   if e != nil {
	//     fmt.Fprintf(os. Stderr, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), e.Error())
	//     continue
	//   }
	//   cat.Cat(bufio.NewReader(f))
	// }

	/*
	   //mappings := map[string] int {
	   //  "s1": 1,f
	   //}
	*/

	// s := []maxUtil.E{1.0,2.0,3.1,0.4}
	// sf := ToFLoat(s)
	// // ss := []maxUtil.E{1,2,999,3,6,1}

	// fmt.Printf("Max value %v", maxUtil.Max2(sf))
	// c = make(chan int)

	// go Doit("Tea", 3)
	// go Doit("Cofee", 3)

	// println("Waiting")

	// var i int

	// for {
	//       select{
	//         case <-c:
	//           i++
	//           println(i)
	//           if i > 1{
	//             return
	//           }
	//       }
	//    }

	// Fibonacci

	// i := make(chan int)
	// q := make(chan int)

	// go func(){
	//   for j:= 2; j <=10; j++{
	//     i <- j
	//   }
	//   q <- -1
	// }()

	// utils.Fibonacci(i,q)

	//Shower
	// i := make(chan int)
	// q := make(chan int)

	// a := [10]int{9,8,7,6,5,4,3,2,1,0}

	// go Shower(i, q)

	// for _, v := range a {
	//   i <- v
	// }
	// q <- -1

	//Binary trees
	// fmt.Printf("Result: %v\n", trees.Same(trees.Tr1, trees.Tr2))

	//Crawler
	// crawler.Crawl("http://golang.org/", 4, crawler.FetcherT)

	//Commands
	// ret, _ := utils.Exec("/bin/ls", "-l")
	// fmt.Printf("Cmd: %v\n", ret)

	//Net/HTTP
	// utils.Get("http://google.com")

	//Children
	// utils.GetChildProcesses()

	//Words counts
	// utils.WC()
	// utils.Uniq()

	//Echo server

	// l, err := net.Listen("tcp", "127.0.0.1:8053")

	// if err != nil {
	//   fmt.Printf("Failed to listen: %s", err.Error())
	// }
	// println("...")

	// for {
	//   c, _ := l.Accept()
	//   go utils.Echo(c)
	// }

	// flag.Parse()
	// list := []int{1, 6, 7, 8, 75, cruncher.ADD, cruncher.SUB, cruncher.MUL, cruncher.DIV}

	// magic, ok := strconv.Atoi(flag.Arg(0))
	// if ok != nil {
	// 	return
	// }

	// f := make([]int, cruncher.MAXPOS)

	// cruncher.Solve(f, list, 0, magic)

	//Finger

     // flag.Parse()
     // ln, err := net.Listen("tcp", ":79")

     // if err != nil {
     //  fmt.Printf("Failed: %v", err.Error())
     // }

     // defer func() {
     //  err := recover()
     //  fmt.Printf("Error %v\n", err)
     // }()

     // for {
     //  l, err := ln.Accept()
     //  if err != nil {
     //    continue
     //  } else {
     //    go utils.HandleConnection(l)
     //  }
     // }
     // x, y := 1, 2
     // even.Swap(&x, &y)

     // z := 1.5
     // even.Square(&z)

     // fmt.Printf("%v %v %v", x, y, z)
     c := make(chan string)
     go g1(c)
     go g2(c)
     go gg(c)

     var input string
     fmt.Scanln(&input)
}

func gg(c <-chan string) {
  timeout := time.After(10*time.Second)
  // for {
  select{
    case c1 := <- c:
      fmt.Println(c1)
    case c2 := <- c:
      fmt.Println(c2)
    case <- timeout:
      fmt.Println("Timeout")
    // default:
    //   fmt.Println("...waiting")
    //   time.Sleep(500*time.Millisecond)
    }
  // }
}

func g1(c chan<- string) {
  time.Sleep(5*time.Second)
  c <- "BOOM1"
}

func g2(c chan<- string) {
  time.Sleep(5*time.Second)
  c <- "BOOM2"
}

func Shower(i, q <-chan int) {
	for {
		select {
		case vv, _ := <-i:
			fmt.Println(vv)
		case <-q:
			return
		}
	}
}

var c chan int

func Doit(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Printf("%s is ready!\n", w)
	c <- -1
}

func ToFLoat(s []maxUtil.E) []float32 {
	a := make([]float32, len(s))
	for i, v := range s {
		a[i] = float32(v.(float64))
	}
	return a
}

func DoubelIt(x mapUtil.E) mapUtil.E {
	var ret mapUtil.E
	switch x.(type) {
	case int:
		v := x.(int)
		ret = v + v
	case string:
		v := x.(string)
		ret = v + v
	}
	return ret
}

func PlusTwo() func(int) int {
	return func(x int) int { return x + 2 }
}

func BubbleSort(x []int) {
	for i := 0; i < len(x)-1; i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i] < x[j] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
}

func Max(n []int) int {
	var max int
	max = n[0]

	for i, v := range n {
		if i == 0 {
			continue
		}

		if max <= v {
			max = v
		}
	}
	return max
}

func Min(a []int) int {
	var min int
	for i, v := range a {
		switch i {
		case 1:
			min = v
		default:
			if min >= v {
				min = v
			}
		}
	}
	return min
}

func Map(f func(string) string, s ...string) []string {
	ss := make([]string, len(s))
	for i, v := range s {
		ss[i] = f(v)
	}
	return ss
}

func strSquare(s string) string {
	return s + strconv.Itoa(2)
}

func intSquare(n int) int {
	return n * n
}

func funcMap(ff func(int) int, n ...int) []int {
	s := make([]int, len(n))
	for i, v := range n {
		s[i] = ff(v)
	}
	return s
}

func fib(n int) (s []int) {

	switch n {
	case 1:
		s = append(s, 1)
		return s
	case 2:
		s = append(s, 1)
		s = append(s, 1)
		return s
	default:
		s = append(s, 1)
		s = append(s, 1)

		for i := 2; i < n; i++ {
			s = append(s, s[i-1]+s[i-2])
		}
		return s
	}
	return
}

func printInt(x ...int) {
	for _, v := range x {
		fmt.Printf(typeOf(v))
	}
}

func typeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// type stack struct {
//   i int
//   data [10]int
// }

//STACK
// func (s *stack) push(x int) (*stack){
//   if s.i > 8 {
//     return s
//   }
//   s.data[s.i] = x
//   s.i++
//   return s
// }

// func (s *stack) String() string{
//   var str string
//   for j := 0; j <= s.i; j++ {
//     str += "[" + strconv.Itoa(j) + ":" + strconv.Itoa(s.data[j]) + "]"
//   }
//   return str
// }

//ASC
func asc(x, y int) (int, int) {
	if x <= y {
		return x, y
	} else {
		return y, x
	}
}

//AVERAGE
func average(x []float64) (ave float64) {
	var sum float64
	for _, v := range x {
		sum += v
	}
	ave = sum / float64(len(x))
	return
}

func throwsError(f func(int)) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()

	f(3)
	return
}

func panicit(x int) {
	panic(x)
}

// func printit(x int) {
// }
