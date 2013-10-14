package main

import (
	"container/list"
	"fmt"
	"os"
	"time"
)

const MAX = 81

var su [MAX]int
var count int = 0
var t0, t1 int64 = 0, 0

type point struct {
	x, y      int
	available []int
	value     int
}

func newPoint(x, y int) *point {
	p := new(point)
	p.x = x
	p.y = y
	p.available = make([]int, 0, 10)
	return p
}

type queue struct {
	*list.List
}

func newqueue() *queue {
	q := new(queue)
	q.List = list.New()
	return q
}

func (q *queue) push(value interface{}) {
	q.PushBack(value)
}

func (q *queue) pop() *list.Element {
	if nil != q {
		e := q.Front()
		if nil != e {
			q.Remove(e)
			return e
		}
	}
	return nil
}

func (q *queue) len() int {
	return q.Len()
}

func (q *queue) isnil() bool {
	e := q.Front()
	return nil == e
}

func mapCount(cs []int) map[int]int {
	m := make(map[int]int, 10)
	for _, c := range cs {
		_, ok := m[c]
		if ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}
	return m
}

func Nums(p *point) []int {
	block_x := p.x / 3
	block_y := p.y / 3
	start := block_y*27 + block_x*3

	nums := []int{
		su[p.y*9], su[p.y*9+1], su[p.y*9+2],
		su[p.y*9+3], su[p.y*9+4], su[p.y*9+5],
		su[p.y*9+6], su[p.y*9+7], su[p.y*9+8],

		su[p.x], su[p.x+9], su[p.x+18],
		su[p.x+27], su[p.x+36], su[p.x+45],
		su[p.x+54], su[p.x+63], su[p.x+72],

		su[start], su[start+1], su[start+2],
		su[start+9], su[start+10], su[start+11],
		su[start+18], su[start+19], su[start+20],
	}

	return nums
}

func NotInNums(num int, nums []int) bool {
	for _, n := range nums {
		if n == num {
			return false
		}
	}

	return true
}

func show() {
	fmt.Printf("========= count=%d =========\n", count)
	count++
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d\t", su[i*9+j])
		}
		fmt.Printf("\n")
	}
}

func initPoint() *queue {
	pq := newqueue()
	for i := 0; i < MAX; i++ {
		if 0 == su[i] {
			p := newPoint(i%9, i/9)
			nums := Nums(p)
			for j := 1; j < 10; j++ {
				if NotInNums(j, nums) {
					p.available = append(p.available, j)
				}
			}
			pq.push(p)
		}
	}

	return pq
}

// 单元唯一法：去除横、竖、块中已存在的数字，如果该格内
// 的数字是唯一的，填入该数字
func SolePosition() {
	again := 1
	for again > 0 {
		again = 0
		for i := 0; i < MAX; i++ {
			if 0 == su[i] {
				p := newPoint(i%9, i/9)
				nums := Nums(p)
				cn := 0
				tmp_j := 0
				for j := 1; j < 10; j++ {
					if NotInNums(j, nums) {
						cn++
						tmp_j = j
					}
				}
				if 1 == cn {
					su[i] = tmp_j
					again++
					fmt.Printf("su[%d][%d]=%d\n", p.x, p.y, tmp_j)
				}
			}
		}
		fmt.Printf("==== again=%d\n", again)
	}
}

//func MustHave() {
//	again := 1
//	for again > 0 {
//		again = 0
//		pq := initPoint()

//		for p:=pq.Front();nil!=p;p=p.Next() {
//			nums := Nums(p)
//			func sigle(p point,ns []int){
//				for _,n := range ns {
//					if 0==n {
//						available:=p.Value.(point).available

//					}
//			}
//		}

//		for i := 0; i < MAX; i++ {
//			if 0 == su[i] {
//				p := newPoint(i%9, i/9)
//				nums := Nums(p)
//				cn := 0
//				tmp_j := 0
//				for j := 1; j < 10; j++ {
//					if NotInNums(j, nums) {
//						cn++
//						tmp_j = j
//					}
//				}
//				if 1 == cn {
//					su[i] = tmp_j
//					again++
//					fmt.Printf("su[%d][%d]=%d\n", p.x, p.y, tmp_j)
//				}
//			}
//		}
//		fmt.Printf("==== again=%d\n", again)
//	}
//}

func check(p *point) bool {
	if 0 == p.value {
		fmt.Printf("not assign value to point p!!\n")
		return false
	}
	if NotInNums(p.value, Nums(p)) {
		return true
	}
	return false
}

func try(p *point, pq *queue) {
	availNum := p.available
	for _, value := range availNum {
		p.value = value
		if check(p) {
			su[p.y*9+p.x] = p.value
			pq := initPoint()
			if pq.len() <= 0 {
				show()
				t1 = time.Now().UnixNano()
				fmt.Printf("use Time:%d nano-second\n", (t1 - t0))
				os.Exit(0)
			} else {
				p2 := pq.pop().Value.(*point)
				try(p2, pq)
				su[p2.y*9+p2.x] = 0
				p2.value = 0
			}
		} else {
			p.value = 0
		}
	}
}

func main() {
	t0 = time.Now().UnixNano()

	//su = [MAX]int{
	//	8, 0, 0, 0, 0, 0, 0, 0, 0,
	//	0, 0, 3, 6, 0, 0, 0, 0, 0,
	//	0, 7, 0, 0, 9, 0, 2, 0, 0,
	//	0, 5, 0, 0, 0, 7, 0, 0, 0,
	//	0, 0, 0, 0, 4, 5, 7, 0, 0,
	//	0, 0, 0, 1, 0, 0, 0, 3, 0,
	//	0, 0, 1, 0, 0, 0, 0, 6, 8,
	//	0, 0, 8, 5, 0, 0, 0, 1, 0,
	//	0, 9, 0, 0, 0, 0, 4, 0, 0,
	//}

	//su = [MAX]int{
	//	0, 0, 0, 0, 0, 0, 0, 0, 0,
	//	0, 0, 0, 0, 0, 0, 0, 0, 0,
	//	0, 0, 0, 0, 0, 0, 0, 0, 0,
	//	0, 0, 0, 0, 0, 0, 0, 0, 0,
	//	0, 0, 0, 0, 0, 0, 0, 0, 0,
	//	0, 0, 0, 0, 0, 0, 0, 0, 0,
	//	0, 0, 0, 0, 0, 0, 0, 0, 0,
	//	0, 0, 0, 0, 0, 0, 0, 0, 0,
	//	0, 0, 0, 0, 0, 0, 0, 0, 0,
	//}

	//su = [MAX]int{
	//	0, 9, 0, 0, 0, 0, 0, 2, 0,
	//	0, 0, 0, 0, 0, 0, 0, 0, 8,
	//	0, 0, 7, 6, 0, 5, 0, 1, 0,
	//	9, 0, 0, 0, 0, 0, 0, 0, 6,
	//	0, 4, 3, 0, 8, 0, 0, 9, 0,
	//	0, 0, 0, 2, 3, 0, 8, 0, 0,
	//	0, 6, 0, 9, 4, 0, 0, 0, 0,
	//	0, 0, 0, 1, 0, 8, 3, 6, 0,
	//	0, 5, 0, 0, 6, 2, 1, 0, 0,
	//}

	//su = [MAX]int{
	//	0, 0, 2, 3, 5, 4, 0, 0, 0,
	//	4, 0, 3, 0, 0, 2, 0, 5, 0,
	//	5, 0, 0, 0, 0, 0, 2, 3, 4,
	//	0, 5, 0, 2, 3, 0, 0, 0, 0,
	//	2, 0, 0, 0, 4, 5, 0, 0, 3,
	//	0, 3, 0, 0, 0, 0, 4, 2, 5,
	//	3, 2, 0, 0, 0, 0, 5, 0, 0,
	//	0, 0, 0, 0, 2, 0, 0, 4, 0,
	//	0, 4, 0, 0, 0, 0, 0, 0, 0,
	//}

	//su = [MAX]int{
	//	0, 0, 0, 4, 9, 0, 0, 0, 0,
	//	0, 0, 6, 1, 0, 0, 7, 0, 0,
	//	0, 0, 3, 0, 5, 0, 0, 0, 0,
	//	0, 8, 7, 0, 0, 3, 0, 0, 0,
	//	0, 0, 0, 0, 0, 0, 0, 5, 0,
	//	0, 0, 0, 0, 0, 0, 0, 0, 9,
	//	9, 1, 0, 0, 0, 0, 0, 4, 0,
	//	0, 0, 0, 0, 0, 0, 0, 0, 0,
	//	0, 0, 0, 0, 0, 7, 3, 0, 0,
	//}

	//su = [MAX]int{
	//	0, 3, 0, 0, 0, 0, 0, 5, 0,
	//	0, 0, 0, 0, 0, 0, 0, 0, 0,
	//	0, 0, 0, 8, 0, 0, 7, 0, 0,
	//	8, 0, 7, 0, 1, 0, 0, 0, 0,
	//	0, 0, 0, 0, 6, 0, 0, 4, 0,
	//	2, 0, 0, 0, 0, 0, 0, 0, 0,
	//	0, 9, 0, 0, 0, 0, 2, 0, 0,
	//	0, 4, 0, 0, 5, 3, 0, 0, 0,
	//	0, 0, 1, 0, 0, 0, 0, 0, 8,
	//}

	//su = [MAX]int{
	//	1, 2, 3, 4, 5, 6, 7, 8, 0,
	//	0, 0, 0, 6, 0, 1, 0, 4, 3,
	//	0, 0, 0, 0, 0, 0, 0, 0, 0,
	//	0, 1, 0, 5, 0, 0, 0, 0, 0,
	//	0, 0, 0, 1, 0, 6, 0, 0, 0,
	//	3, 0, 0, 0, 0, 0, 0, 0, 5,
	//	5, 3, 0, 0, 0, 0, 0, 6, 1,
	//	0, 0, 0, 0, 0, 0, 0, 0, 4,
	//	0, 0, 0, 0, 0, 0, 0, 0, 0,
	//}

	su = [MAX]int{
		0, 0, 9, 0, 0, 6, 0, 8, 0,
		2, 0, 8, 3, 1, 0, 5, 0, 6,
		0, 6, 4, 0, 7, 0, 2, 0, 3,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 1,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		1, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 0, 0,
	}

	show()
	//SolePosition()
	pq := initPoint()

	fmt.Printf("len(pq)=%d\n", pq.len())
	for p := pq.pop(); nil != p; p = pq.pop() {
		fmt.Printf("(%d,%d)=>%d\n", p.Value.(*point).x, p.Value.(*point).y, p.Value.(*point).available)
	}

	//fmt.Printf("----\n")
	//show()
	//p := pq.pop().Value.(*point)
	//try(p, pq)

	pp := newPoint(0, 0)
	ns := Nums(pp)
	fmt.Printf("---------\n")
	m := mapCount(ns[0:9])
	for i, v := range m {
		fmt.Printf("m[%d]=%d\n", i, v)
	}
	fmt.Printf("---------\n")
	m = mapCount(ns[9:18])
	for i, v := range m {
		fmt.Printf("m[%d]=%d\n", i, v)
	}
	fmt.Printf("---------\n")
	m = mapCount(ns[18:27])
	for i, v := range m {
		fmt.Printf("m[%d]=%d\n", i, v)
	}

}
