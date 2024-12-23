package main

import (
	"strings"
)

// todo actual tree

type knot struct {
	name string
	knot []knot
}

type depth struct {
	d, i int
}

/*func main() {
	tree := parse(text)
	for _, k := range tree {
		fmt.Println(k.name)
		for _, n := range k.knot {
			fmt.Println("	", n.name)
		}
	}
}*/

func mkDepths(text string) []depth {
	depthSlice := []depth{}
	d := 0
	for i, c := range text {
		switch c {
		case '[':
			d++
			depthSlice = append(depthSlice, depth{d: d, i: i})
		case ']':
			d--
			depthSlice = append(depthSlice, depth{d: d, i: i})
		}
	}
	return depthSlice
}

func s2t(str string, depth []depth, start, finish int) []knot {
	tree := []knot{}
	ss := strings.Fields(strings.Trim(str[start:finish], "[]")) // todo bufio.ScanWords?
	if len(ss) > 0 {
		tree = append(tree, knot{name: ss[0], knot: nil})
		for _, v := range ss[1:] {
			tree[0].knot = append(tree[0].knot, knot{name: v, knot: nil})
		}
	}
	return tree
}

func parse(text string) []knot {
	tree := []knot{}
	depths := mkDepths(text)
	pr := 0
	for _, v := range depths {
		tree = append(tree, s2t(text, depths, pr, v.i)...)
		pr = v.i
	}
	return tree
}

// todo
/*func walk(tree []knot, fn func(knot)) {
	for _, k := range tree {
		fn(k)
	}
}*/
