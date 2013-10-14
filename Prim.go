package main

import (
	"fmt"
	"math"
)

const (
	MAXNUM = 6
)

type Node struct {
	ID        int
	Neighbors []int
	Weight    []int
}

type Graph struct {
	Vertex []*Node
}

type PrimTable struct {
	NodeID int
	Known  bool
	dv     int
	pv     int
}

func NewNode(id int) *Node {
	n := new(Node)
	n.ID = id
	n.Neighbors = make([]int, 0)
	n.Weight = make([]int, 0)
	return n
}

func InitGraph() *Graph {
	g := new(Graph)
	n1 := NewNode(1)
	n1.Neighbors = append(n1.Neighbors, 2, 3, 4)
	n1.Weight = append(n1.Weight, 6, 1, 5)
	g.Vertex = append(g.Vertex, n1)

	n2 := NewNode(2)
	n2.Neighbors = append(n2.Neighbors, 1, 3, 5)
	n2.Weight = append(n2.Weight, 6, 5, 3)
	g.Vertex = append(g.Vertex, n2)

	n3 := NewNode(3)
	n3.Neighbors = append(n3.Neighbors, 1, 2, 4, 5, 6)
	n3.Weight = append(n3.Weight, 1, 5, 5, 6, 4)
	g.Vertex = append(g.Vertex, n3)

	n4 := NewNode(4)
	n4.Neighbors = append(n4.Neighbors, 1, 3, 6)
	n4.Weight = append(n4.Weight, 5, 5, 2)
	g.Vertex = append(g.Vertex, n4)

	n5 := NewNode(5)
	n5.Neighbors = append(n5.Neighbors, 2, 3, 6)
	n5.Weight = append(n5.Weight, 3, 6, 6)
	g.Vertex = append(g.Vertex, n5)

	n6 := NewNode(6)
	n6.Neighbors = append(n6.Neighbors, 3, 4, 5)
	n6.Weight = append(n6.Weight, 4, 2, 6)
	g.Vertex = append(g.Vertex, n6)

	return g
}

func (g *Graph) InitPrimTable() []*PrimTable {
	fmt.Printf("%d\n", len(g.Vertex))
	pt := make([]*PrimTable, 0, len(g.Vertex))
	for _, node := range g.Vertex {
		p := new(PrimTable)
		p.NodeID = node.ID
		p.Known = false
		p.dv = math.MaxInt32
		p.pv = -1
		pt = append(pt, p)
	}
	return pt
}

func (g *Graph) Print() {
	fmt.Printf("\t1\t2\t3\t4\t5\t6\n")
	for i := range g.Vertex {
		fmt.Printf("%d\t", g.Vertex[i].ID)
		s := [MAXNUM]byte{48, 48, 48, 48, 48, 48}
		for j := range g.Vertex[i].Neighbors {
			s[g.Vertex[i].Neighbors[j]-1] = byte(48 + g.Vertex[i].Weight[j])
			//fmt.Printf("%d\t", g.Vertex[i].Neighbors[j])
		}
		for _, c := range s {
			fmt.Printf("%c\t", c)
		}
		fmt.Println()
	}
}

func (g *Graph) Prim() {
	pt := g.InitPrimTable()
	for i, p := range pt {
		if !p.Known {
			for _, n := range g.Vertex[p.NodeID-1].Neighbors {

			}
		}
	}
}

func main() {
	g := InitGraph()
	pt := g.InitPrimTable()
	for i, p := range pt {
		fmt.Printf("i=%d,p.NodeID=%d,p.dv=%d\n", i, p.NodeID, p.dv)
	}
	fmt.Println("\n=============")
	g.Print()
}
