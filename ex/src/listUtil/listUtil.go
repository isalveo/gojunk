package listUtil

import(
  "fmt"
  "errors"
  // "container/list"
)
var ListT *List = new(List)

type Value int

type Node struct{
  Value
  prev, next *Node
}

type List struct{
    PrevT, NextT *Node
  head, tail *Node
}

func (l *List) Front() *Node{
  return l.head
}

func (n *Node) Next() *Node{
  return n.next
}

func (l *List) Push(v Value) *List{
  n := &Node{Value: v}

  if l.head == nil {
    l.head = n
  } else {
    l.tail.next = n
    n.prev = l.tail
  }
  l.tail = n

  return l
}

var errEmpty = errors.New("List is empty")

func (l *List) Pop() (v Value, err error) {
  if l.head == nil {
    err = errEmpty
  } else {
    v = l.head.Value
    l.head = l.head.next
  }
  if l.head == nil {
    l.tail = nil
  }
  return
}

func Print(l *List) {
  for e := l.Front(); e != nil; e = e.Next(){
    fmt.Printf("%v\n", e.Value)
  }
}

// var List *list.List = list.New()

// func Print(l *list.List) {
//   for e := l.Front(); e != nil; e = e.Next(){
//     fmt.Printf("%v\n", e.Value)
//   }
// }