package main

import (
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"reflect"
	"slices"
	"testing"
)

var nums [][]int

func fromSize(size uint) [][]int {
	nums := make([][]int, 0, 10)
	tmp1 := make([]int, 0, size)
	tmp2 := make([]int, 0, size)
	tmp3 := make([]int, 0, size/2)
	tmp4 := make([]int, 0, size/2)
	tmp5 := make([]int, 0, size)
	tmp6 := make([]int, 0, size)
	tmp7 := make([]int, 0, size)
	tmp8 := make([]int, 0, size)
	i := 0
	isize := int(size)
	for i = range isize {
		tmp1 = append(tmp1, i)
		tmp2 = append(tmp2, i-(isize-1))
		if i&1 == 0 {
			tmp3 = append(tmp3, i)
		} else {
			tmp4 = append(tmp4, i)
		}
		tmp5 = append(tmp5, isize/2)
		tmp6 = append(tmp6, 0)
		tmp7 = append(tmp7, isize)
		if i < isize/2 {
			tmp8 = append(tmp7, math.MaxInt32)
		} else {
			tmp8 = append(tmp7, math.MinInt32)
		}
	}
	nums = append(nums, tmp1)
	nums = append(nums, tmp2)
	nums = append(nums, append(tmp1, tmp2...))
	nums = append(nums, tmp3)
	nums = append(nums, tmp4)
	nums = append(nums, tmp5)
	nums = append(nums, tmp6)
	nums = append(nums, tmp7)
	nums = append(nums, tmp8)

	return nums
}

func createNums() [][]int {
	nums := make([][]int, 0)
	nums = append(nums, []int{})
	nums = append(nums, []int{-1})
	nums = append(nums, []int{0})
	nums = append(nums, []int{1})
	nums = append(nums, fromSize(2)...)
	nums = append(nums, fromSize(5)...)
	nums = append(nums, fromSize(10)...)
	nums = append(nums, fromSize(100)...)
	nums = append(nums, fromSize(1000)...)
	nums = append(nums, fromSize(10000)...)
	nums = append(nums, fromSize(999999)...)
	nums = append(nums, fromSize(1000000)...)
	return nums
}

func buildFailString(s string) (string, string) {
	tmp := s + " got %v != %v"
	return tmp, s + " j=%v," + tmp
}

func init() {
	nums = createNums()
}

func createNodes(nums ...int) (head, tail *Node, nodes []*Node) {
	nodes = make([]*Node, 0, len(nums))
	var node *Node
	for i := range nums {
		node = &Node{value: nums[i]}
		if i == 0 {
			head = node
			head.next = nil
			head.prev = nil
		} else {
			tail.next = node
			node.prev = tail
		}
		tail = node
		tail.next = nil
		nodes = append(nodes, node)
	}
	return head, tail, nodes
}

func findAndDelete(src []int, n int, all bool) []int {
	s := make([]int, len(src))
	copy(s, src)
	i := slices.IndexFunc(s, func(i int) bool {
		return i == n
	})
	if i == -1 {
		return s
	}
	if !all {
		return append(s[:i], s[i+1:]...)
	}
	s = slices.DeleteFunc(src, func(i int) bool {
		return i == n
	})
	return s
}

func TestNodes_create(t *testing.T) {
	type want struct {
		nodes []*Node
		ll    LinkedList2
	}
	type test struct {
		name string
		args []int
		want want
	}

	tests := make([]test, 0, len(nums))
	var (
		name             string
		ll               LinkedList2
		node, head, tail *Node
		nodes            []*Node
	)
	sFailBuild := "fail build data %s: %v != %v"
	for i := range nums {
		name = fmt.Sprintf("%d_len_%d", i, len(nums[i]))
		ll = LinkedList2{}
		nodes = make([]*Node, 0, len(nums[i]))
		for j := range nums[i] {
			ll.AddInTail(Node{value: nums[i][j]})
			nodes = append(nodes, ll.tail)
		}
		j := 0
		node = ll.head
		for node != nil {
			if node.value != nums[i][j] {
				t.Errorf(sFailBuild, name, node.value, nums[i][j])
			}
			if !reflect.DeepEqual(*nodes[j], *node) {
				t.Errorf(sFailBuild, name, *nodes[j], *node)
			}
			node = node.next
			j++
		}
		if j != len(nodes) {
			t.Errorf(sFailBuild, name, j, len(nodes))
		}
		tests = append(tests, test{
			name: name,
			args: nums[i],
			want: want{
				nodes: nodes,
				ll:    ll,
			},
		})
	}

	sFailRun, sFailRunWithIndex := buildFailString("createNodes()")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head, tail, nodes = createNodes(tt.args...)
			if head == nil && len(nodes) == 0 && tt.want.ll.head == nil && len(tt.want.nodes) == 0 {
				return
			}
			if head.value != tt.want.ll.head.value {
				t.Errorf(sFailRun, *head, *tt.want.ll.head)
			}
			if tail.value != tt.want.ll.tail.value {
				t.Errorf(sFailRun, *tail, *tt.want.ll.tail)
			}
			if len(nodes) != len(tt.want.nodes) {
				t.Errorf(sFailRun, len(nodes), len(tt.want.nodes))
			}
			node = head
			for j := range nodes {
				if nodes[j].value != tt.want.nodes[j].value {
					t.Errorf(sFailRunWithIndex, j, *nodes[j], *tt.want.nodes[j])
				}
				if nodes[j].value != node.value {
					t.Errorf(sFailRunWithIndex, j, *nodes[j], *node)
				}
				node = node.next
			}
			node = tail
			for j := len(tt.want.nodes) - 1; j >= 0; j-- {
				if nodes[j].value != tt.want.nodes[j].value {
					t.Errorf(sFailRunWithIndex, j, *nodes[j], *tt.want.nodes[j])
				}
				if nodes[j].value != node.value {
					t.Errorf(sFailRunWithIndex, j, *nodes[j], *node)
				}
				node = node.prev
			}
		})
	}
}

func TestLinkedList2_AddInTail(t *testing.T) {
	type want struct {
		nodes []*Node
		ll    LinkedList2
	}
	type test struct {
		name string
		args []*Node
		want want
	}

	tests := make([]test, 0, len(nums))
	var (
		name             string
		ll               LinkedList2
		node, head, tail *Node
		nodes            []*Node
	)
	for i := range nums {
		name = fmt.Sprintf("%d_len_%d", i, len(nums[i]))
		head, tail, nodes = createNodes(nums[i]...)
		ll = LinkedList2{
			head: head,
			tail: tail,
		}
		tests = append(tests, test{
			name: name,
			args: nodes,
			want: want{
				nodes: nodes,
				ll:    ll,
			},
		})
	}

	sFailRun, sFailRunWithIndex := buildFailString("AddInTail()")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll = LinkedList2{}
			for i := range tt.args {
				ll.AddInTail(*tt.args[i])
			}
			head, tail = ll.head, ll.tail
			if head == nil && tt.want.ll.head == nil && len(tt.want.nodes) == 0 {
				return
			}
			if head.value != tt.want.ll.head.value {
				t.Errorf(sFailRun, *head, *tt.want.ll.head)
			}
			if tail.value != tt.want.ll.tail.value {
				t.Errorf(sFailRun, *tail, *tt.want.ll.tail)
			}
			node = head
			tmp := tt.want.ll.head
			for j := range tt.want.nodes {
				if node.value != tt.want.nodes[j].value {
					t.Errorf(sFailRunWithIndex, j, *node, *tt.want.nodes[j])
				}
				if node.value != tmp.value {
					t.Errorf(sFailRunWithIndex, j, *node, *tmp)
				}
				node = node.next
				tmp = tmp.next
			}
			node = tail
			tmp = tt.want.ll.tail
			for j := len(tt.want.nodes) - 1; j >= 0; j-- {
				if node.value != tt.want.nodes[j].value {
					t.Errorf(sFailRunWithIndex, j, *node, *tt.want.nodes[j])
				}
				if node.value != tmp.value {
					t.Errorf(sFailRunWithIndex, j, *node, *tmp)
				}
				node = node.prev
				tmp = tmp.prev
			}
		})
	}
}

func TestLinkedList2_Clean(t *testing.T) {
	type want struct {
		node *Node
	}
	type test struct {
		name string
		args LinkedList2
		want want
	}

	tests := make([]test, 0, len(nums))
	var (
		name       string
		ll         LinkedList2
		head, tail *Node
	)
	for i := range nums {
		name = fmt.Sprintf("%d_len_%d", i, len(nums[i]))
		head, tail, _ = createNodes(nums[i]...)
		ll = LinkedList2{
			head: head,
			tail: tail,
		}
		tests = append(tests, test{
			name: name,
			args: ll,
			want: want{},
		})
	}

	sFailRun, _ := buildFailString("Clean()")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.Clean()
			head = tt.args.head
			if head != tt.want.node {
				t.Errorf(sFailRun, head, tt.want.node)
			}
			tail = tt.args.tail
			if tail != tt.want.node {
				t.Errorf(sFailRun, head, tt.want.node)
			}
		})
	}
}

func TestLinkedList2_Count(t *testing.T) {
	type want struct {
		count int
	}
	type test struct {
		name string
		args LinkedList2
		want want
	}

	tests := make([]test, 0, len(nums))
	var (
		name       string
		ll         LinkedList2
		head, tail *Node
		nodes      []*Node
	)
	for i := range nums {
		name = fmt.Sprintf("%d_len_%d", i, len(nums[i]))
		head, tail, nodes = createNodes(nums[i]...)
		ll = LinkedList2{
			head: head,
			tail: tail,
		}
		tests = append(tests, test{
			name: name,
			args: ll,
			want: want{count: len(nodes)},
		})
	}

	sFailRun, _ := buildFailString("Count()")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.Count(); got != tt.want.count {
				t.Errorf(sFailRun, got, tt.want.count)
			}
		})
	}
}

func TestLinkedList2_Delete(t *testing.T) {
	type args struct {
		n     int
		all   bool
		ll    LinkedList2
		nodes []*Node
	}
	type want struct {
		nodes []*Node
		ll    LinkedList2
	}
	type test struct {
		name string
		args args
		want want
	}

	tests := make([]test, 0, len(nums))
	var (
		name             string
		ll               LinkedList2
		head, tail, node *Node
		nodes            []*Node
		n                int
		all              bool
	)
	for i := range nums {
		name = fmt.Sprintf("%d_len_%d", i, len(nums[i]))
		head, tail, _ = createNodes(nums[i]...)
		ll = LinkedList2{
			head: head,
			tail: tail,
		}
		cc := make([]int, len(nums[i]))
		copy(cc, nums[i])
		sz := len(cc)
		for j := range nums[i] {
			if i%5 == 0 {
				n, all = 12345667890, false
				break
			}
			n, all = nums[i][j], nums[i][j]%2 == 0
			tmp := findAndDelete(cc, n, all)
			if sz != len(tmp) {
				cc = tmp
				break
			}
		}
		head, tail, nodes = createNodes(cc...)
		tests = append(tests, test{
			name: name,
			args: args{
				all: all,
				n:   n,
				ll:  ll,
			},
			want: want{
				nodes: nodes,
				ll: LinkedList2{
					head: head,
					tail: tail,
				},
			},
		})
	}

	sFailRun, sFailRunWithIndex := buildFailString("Delete()")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.ll.Delete(tt.args.n, tt.args.all)
			head, tail = tt.args.ll.head, tt.args.ll.tail
			if head == nil && tt.want.ll.head == nil && len(tt.want.nodes) == 0 {
				return
			}
			if (head == nil) != (tt.want.ll.head == nil) {
				t.Errorf(sFailRun, head, tt.want.ll.head)
			}
			if (tail == nil) != (tt.want.ll.tail == nil) {
				t.Errorf(sFailRun, tail, tt.want.ll.tail)
			}
			if head.value != tt.want.ll.head.value {
				t.Errorf(sFailRun, *head, *tt.want.ll.head)
			}
			if tail.value != tt.want.ll.tail.value {
				t.Errorf(sFailRun, *tail, *tt.want.ll.tail)
			}
			node = head
			tmp := tt.want.ll.head
			for j := range tt.want.nodes {
				if node.value != tt.want.nodes[j].value {
					t.Errorf(sFailRunWithIndex, j, *node, *tt.want.nodes[j])
				}
				if node.value != tmp.value {
					t.Errorf(sFailRunWithIndex, j, *node, *tmp)
				}
				node = node.next
				tmp = tmp.next
			}
			node = tail
			tmp = tt.want.ll.tail
			for j := len(tt.want.nodes) - 1; j >= 0; j-- {
				if node.value != tt.want.nodes[j].value {
					t.Errorf(sFailRunWithIndex, j, *node, *tt.want.nodes[j])
				}
				if node.value != tmp.value {
					t.Errorf(sFailRunWithIndex, j, *node, *tmp)
				}
				node = node.prev
				tmp = tmp.prev
			}
		})
	}
}

func TestLinkedList2_Find(t *testing.T) {
	type args struct {
		n  int
		ll LinkedList2
	}
	type want struct {
		n      int
		hasErr bool
	}
	type test struct {
		name string
		args args
		want want
	}

	tests := make([]test, 0, len(nums))
	var (
		name       string
		head, tail *Node
		nodes      []*Node
		n, wantN   int
		hasErr     bool
	)
	for i := range nums {
		name = fmt.Sprintf("%d_len_%d", i, len(nums[i]))
		head, tail, nodes = createNodes(nums[i]...)
		n, wantN, hasErr = -1, -1, true
		if len(nums[i]) > 0 {
			num := nums[i][len(nums[i])/2] / 2
			tmp := slices.IndexFunc(nodes, func(find *Node) bool {
				return find.value == num
			})
			if tmp == -1 {
				n, wantN, hasErr = num, -1, true
			} else {
				n, wantN, hasErr = num, num, tmp == -1
			}
		}
		tests = append(tests, test{
			name: name,
			args: args{
				n: n,
				ll: LinkedList2{
					head: head,
					tail: tail,
				},
			},
			want: want{
				n:      wantN,
				hasErr: hasErr,
			},
		})
	}

	sFailRun, _ := buildFailString("Find()")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, err := tt.args.ll.Find(tt.args.n)
			if (err != nil) != tt.want.hasErr {
				t.Errorf(sFailRun, err != nil, tt.want.hasErr)
			}
			if node.value != tt.want.n {
				t.Errorf(sFailRun, node.value, tt.want.n)
			}
		})
	}
}

func TestLinkedList2_FindAll(t *testing.T) {
	type args struct {
		n  int
		ll LinkedList2
	}
	type want struct {
		nodes []Node
	}
	type test struct {
		name string
		args args
		want want
	}

	tests := make([]test, 0, len(nums))
	var (
		name       string
		head, tail *Node
		nodes      []*Node
		n          int
	)
	for i := range nums {
		name = fmt.Sprintf("%d_len_%d", i, len(nums[i]))
		head, tail, nodes = createNodes(nums[i]...)
		tmp := make([]Node, 0, 10)
		if len(nums[i]) > 0 {
			num := nums[i][len(nums[i])/2] / 2
			for j := range nodes {
				if nodes[j].value == num {
					tmp = append(tmp, *nodes[j])
				}
			}
			n = num
		}
		tests = append(tests, test{
			name: name,
			args: args{
				n: n,
				ll: LinkedList2{
					head: head,
					tail: tail,
				},
			},
			want: want{nodes: tmp},
		})
	}

	sFailRun, sFailRunWithIndex := buildFailString("FindAll()")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ll.FindAll(tt.args.n)
			if len(got) != len(tt.want.nodes) {
				log.Println("------>", len(got), len(tt.want.nodes))
				if tt.name == "13_len_5" {
					log.Println("------->", tt.args.n)
					node := tt.args.ll.head
					for node != nil {
						log.Println("---->", *node)
						node = node.next
					}
				}
				t.Errorf(sFailRun, len(got), len(tt.want.nodes))
			}
			for i := range tt.want.nodes {
				if got[i].value != tt.want.nodes[i].value {
					t.Errorf(sFailRunWithIndex, i, got[i], tt.want.nodes[i])
				}
			}
		})
	}
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func TestLinkedList2_Insert(t *testing.T) {
	type args struct {
		after *Node
		node  Node
		ll    LinkedList2
	}
	type want struct {
		nodes []*Node
		ll    LinkedList2
	}
	type test struct {
		name string
		args args
		want want
	}

	tests := make([]test, 0, len(nums))
	var (
		name             string
		head, tail, node *Node
		tmp              Node
		nodes            []*Node
		ll               LinkedList2
	)
	value := 10101010
	for i := range nums {
		name = fmt.Sprintf("%d_len_%d", i, len(nums[i]))
		head, tail, nodes = createNodes(nums[i]...)
		ll = LinkedList2{
			head: head,
			tail: tail,
		}
		cc := make([]int, len(nums[i]))
		copy(cc, nums[i])
		node = nil
		idx := 0
		if len(nums[i]) > 0 {
			num := nums[i][len(nums[i])/2] / 2
			switch {
			case num%15 == 0:
			case num%6 == 0:
				idx = (len(cc) - 1) / 2
			case num%2 == 0:
				idx = len(cc) - 1
			default:
				idx = randRange(0, len(cc)-1)
			}
			if len(cc)-1 == idx {
				cc = append(cc, value)
			} else {
				cc = append(cc[:idx+2], cc[idx+1:]...)
				cc[idx+1] = value
			}
			node = nodes[idx]
		}
		tmp = Node{value: value}
		head, tail, nodes = createNodes(cc...)
		tests = append(tests, test{
			name: name,
			args: args{
				after: node,
				node:  tmp,
				ll:    ll,
			},
			want: want{
				nodes: nodes,
				ll: LinkedList2{
					head: head,
					tail: tail,
				},
			},
		})
	}

	sFailRun, sFailRunWithIndex := buildFailString("Insert()")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.ll.Insert(tt.args.after, tt.args.node)
			head, tail = tt.args.ll.head, tt.args.ll.tail
			if head == nil && tt.want.ll.head == nil && len(tt.want.nodes) == 0 {
				return
			}
			if (head == nil) != (tt.want.ll.head == nil) {
				t.Errorf(sFailRun, head, tt.want.ll.head)
			}
			if (tail == nil) != (tt.want.ll.tail == nil) {
				t.Errorf(sFailRun, tail, tt.want.ll.tail)
			}
			if head.value != tt.want.ll.head.value {
				t.Errorf(sFailRun, *head, *tt.want.ll.head)
			}
			if tail.value != tt.want.ll.tail.value {
				t.Errorf(sFailRun, *tail, *tt.want.ll.tail)
			}
			node = head
			tmp := tt.want.ll.head
			for j := range tt.want.nodes {
				if node.value != tt.want.nodes[j].value {
					t.Errorf(sFailRunWithIndex, j, *node, *tt.want.nodes[j])
				}
				if node.value != tmp.value {
					t.Errorf(sFailRunWithIndex, j, *node, *tmp)
				}
				node = node.next
				tmp = tmp.next
			}
			node = tail
			tmp = tt.want.ll.tail
			for j := len(tt.want.nodes) - 1; j >= 0; j-- {
				if node.value != tt.want.nodes[j].value {
					t.Errorf(sFailRunWithIndex, j, *node, *tt.want.nodes[j])
				}
				if node.value != tmp.value {
					t.Errorf(sFailRunWithIndex, j, *node, *tmp)
				}
				node = node.prev
				tmp = tmp.prev
			}
		})
	}
}

func TestLinkedList2_InsertFirst(t *testing.T) {
	type args struct {
		node Node
		ll   LinkedList2
	}
	type want struct {
		nodes []*Node
		ll    LinkedList2
	}
	type test struct {
		name string
		args args
		want want
	}

	tests := make([]test, 0, len(nums))
	var (
		name             string
		head, tail, node *Node
		tmp              Node
		nodes            []*Node
		ll               LinkedList2
	)
	value := 101010101
	for i := range nums {
		name = fmt.Sprintf("%d_len_%d", i, len(nums[i]))
		head, tail, nodes = createNodes(nums[i]...)
		ll = LinkedList2{
			head: head,
			tail: tail,
		}
		cc := append([]int{value}, nums[i]...)
		tmp = Node{value: value}
		head, tail, nodes = createNodes(cc...)
		tests = append(tests, test{
			name: name,
			args: args{
				node: tmp,
				ll:   ll,
			},
			want: want{
				nodes: nodes,
				ll: LinkedList2{
					head: head,
					tail: tail,
				},
			},
		})
	}

	sFailRun, sFailRunWithIndex := buildFailString("InsertFirst()")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.ll.InsertFirst(tt.args.node)
			head, tail = tt.args.ll.head, tt.args.ll.tail
			if head == nil && tt.want.ll.head == nil && len(tt.want.nodes) == 0 {
				return
			}
			if (head == nil) != (tt.want.ll.head == nil) {
				t.Errorf(sFailRun, head, tt.want.ll.head)
			}
			if (tail == nil) != (tt.want.ll.tail == nil) {
				t.Errorf(sFailRun, tail, tt.want.ll.tail)
			}
			if head.value != tt.want.ll.head.value {
				t.Errorf(sFailRun, *head, *tt.want.ll.head)
			}
			if tail.value != tt.want.ll.tail.value {
				t.Errorf(sFailRun, *tail, *tt.want.ll.tail)
			}
			node = head
			tmp := tt.want.ll.head
			for j := range tt.want.nodes {
				if node.value != tt.want.nodes[j].value {
					t.Errorf(sFailRunWithIndex, j, *node, *tt.want.nodes[j])
				}
				if node.value != tmp.value {
					t.Errorf(sFailRunWithIndex, j, *node, *tmp)
				}
				node = node.next
				tmp = tmp.next
			}
			node = tail
			tmp = tt.want.ll.tail
			for j := len(tt.want.nodes) - 1; j >= 0; j-- {
				if node.value != tt.want.nodes[j].value {
					t.Errorf(sFailRunWithIndex, j, *node, *tt.want.nodes[j])
				}
				if node.value != tmp.value {
					t.Errorf(sFailRunWithIndex, j, *node, *tmp)
				}
				node = node.prev
				tmp = tmp.prev
			}
		})
	}
}
