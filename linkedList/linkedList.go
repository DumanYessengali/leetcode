package linkedList

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL
	Tail *NodeL
}
type NodeI struct {
	Data int
	Next *NodeI
}

func ListPushBack(l *List, data interface{}) {
	node := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	l.Tail.Next = node
	l.Tail = l.Tail.Next
}
func ListPushFront(l *List, data interface{}) {
	node := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	node.Next = l.Head
	l.Head = node
}
func ListSize(l *List) int {
	c := 0
	for l.Head != nil {
		c++
		l.Head = l.Head.Next
	}
	return c
}
func ListLast(l *List) interface{} {
	if l.Tail == nil {
		return nil
	}
	return l.Tail.Data
}
func PrintList(l *List) {
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
}
func ListClear(l *List) { l.Head = nil }
func ListAt(l *NodeL, pos int) *NodeL {
	c := l
	n := 0
	for c != nil {
		if n == pos {
			return c
		}
		n++
		c = c.Next
	}
	return nil
}
func ListReverse(l *List) {
	new_l := &List{}
	t := *l
	for t.Head != nil {
		ListPushFront(new_l, t.Head.Data)
		t.Head = t.Head.Next
	}
	l.Head = new_l.Head
	l.Tail = new_l.Tail
}
func ListForEach(l *List, f func(*NodeL)) {
	new_l := &List{}
	for l.Head != nil {
		f(l.Head)
		ListPushBack(new_l, l.Head.Data)
		l.Head = l.Head.Next
	}
	l.Head = new_l.Head
	l.Tail = new_l.Tail
}
func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}
func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
func IsPositive_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	case string, rune:
		return false
	}
	return false
}
func IsNegative_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) < 0
	case string, rune:
		return false
	}
	return false
}
func IsNotNumeric_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	case string, rune:
		return true
	}
	return true
}
func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	new_l := &List{}
	for l.Head != nil {
		if cond(l.Head) {
			f(l.Head)
		}
		ListPushBack(new_l, l.Head.Data)
		l.Head = l.Head.Next
	}
	l.Head = new_l.Head
	l.Tail = new_l.Tail
}
func PrintElem(node *NodeL)         { fmt.Println(node.Data) }
func StringToInt(node *NodeL)       { node.Data = 2 }
func CompStr(a, b interface{}) bool { return a == b }
func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	for l.Head != nil {
		if comp(l.Head.Data, ref) {
			return &l.Head.Data
		}
		l.Head = l.Head.Next
	}
	return nil
}
func ListRemoveIf(l *List, data_ref interface{}) {
	new_l := &List{}
	for l.Head != nil {
		if l.Head.Data == data_ref {
			l.Head = l.Head.Next
			continue
		}
		ListPushBack(new_l, l.Head.Data)
		l.Head = l.Head.Next
	}
	l.Head = new_l.Head
	l.Tail = new_l.Tail
}
func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}
	l1.Tail.Next = l2.Head
}
func PrintList2(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}
func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}
	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
func ListSort(l *NodeI) *NodeI {
	var arr []int
	for l != nil {
		arr = append(arr, l.Data)
		l = l.Next
	}
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j-1] > arr[j] {
				intermediate := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = intermediate
			}
		}
	}
	for _, v1 := range arr {
		l = listPushBack(l, v1)
	}
	return l
}
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	var arr []int
	for l != nil {
		arr = append(arr, l.Data)
		l = l.Next
	}
	arr = append(arr, data_ref)
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j-1] > arr[j] {
				intermediate := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = intermediate
			}
		}
	}
	for _, v1 := range arr {
		l = listPushBack(l, v1)
	}
	return l
}
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	var arr []int
	for n1 != nil {
		arr = append(arr, n1.Data)
		n1 = n1.Next
	}
	for n2 != nil {
		arr = append(arr, n2.Data)
		n2 = n2.Next
	}
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j-1] > arr[j] {
				intermediate := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = intermediate
			}
		}
	}
	for _, v1 := range arr {
		n1 = listPushBack(n1, v1)
	}
	return n1
}
func main() {
	var link *NodeI
	var link2 *NodeI
	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)
	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)
	PrintList2(SortedListMerge(link2, link))
}
