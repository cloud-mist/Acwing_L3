// 最小步
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = int(1e5)

type p struct {
	x, y string
}

var (
	dist       map[string]int
	pre        map[string]p
	g          [2][4]byte
	q          [N]string
	start, end string
)

func main() {
	defer ot.Flush()
	dist = make(map[string]int)
	pre = make(map[string]p)

	start = "12345678"
	for i := 0; i < 8; i++ {
		end += rnS()
	}

	bfs()

	fmt.Println(dist[end])

	res := ""
	for end != start {
		res += pre[end].x
		end = pre[end].y
	}

	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(string(res[i]))
	}
}

func bfs() {
	if start == end {
		return
	}

	hh, tt := 0, -1
	tt++
	q[tt] = start

	for hh <= tt {
		t := q[hh]
		hh++

		var op [3]string
		op[0] = move0(t)
		op[1] = move1(t)
		op[2] = move2(t)

		for i := 0; i < 3; i++ {
			m := op[i]
			// 如果没有被访问过
			if dist[m] == 0 {

				dist[m] = dist[t] + 1
				pre[m] = p{string(byte(i + 'A')), t}

				tt++
				q[tt] = m

				if m == end {
					return
				}
			}

		}
	}
}

func move0(state string) string {
	set(state)

	for i := 0; i < 4; i++ {
		g[0][i], g[1][i] = g[1][i], g[0][i]
	}

	return get()
}
func move1(state string) string {
	set(state)

	tmp0, tmp1 := g[0][3], g[1][3]
	for i := 3; i >= 1; i-- {
		g[0][i] = g[0][i-1]
		g[1][i] = g[1][i-1]
	}
	g[0][0], g[1][0] = tmp0, tmp1

	return get()
}
func move2(state string) string {
	set(state)

	tmp := g[0][1]
	g[0][1] = g[1][1]
	g[1][1] = g[1][2]
	g[1][2] = g[0][2]
	g[0][2] = tmp

	return get()
}

// 把字符串放到 2*4矩阵
func set(state string) {
	for i := 0; i < 4; i++ {
		g[0][i] = state[i]
	}
	j := 4
	for i := 3; i >= 0; i-- {
		g[1][i] = state[j]
		j++

	}
}

// 把2*4矩阵 变成 顺时针的字符串
func get() string {
	res := ""

	for i := 0; i < 4; i++ {
		res += string(g[0][i])
	}
	for i := 3; i >= 0; i-- {
		res += string(g[1][i])

	}

	return res
}

/* ======================================================================== */
//
//
//
//				 _____   _   _   ____
//				| ____| | \ | | |  _ \
//				|  _|   |  \| | | | | |
//				| |___  | |\  | | |_| |
//				|_____| |_| \_| |____/
//
//
//
/* ============================PART 1: I/O ================================== */

var (
	ot = bufio.NewWriterSize(os.Stdout, int(1e6))
	in = bufio.NewScanner(os.Stdin)
)

func init()        { in.Split(bufio.ScanWords); in.Buffer(make([]byte, 4096), int(1e9)) }
func rnS() string  { in.Scan(); return in.Text() }
func rnI() int     { i, _ := strconv.Atoi(rnS()); return i }
func rnF() float64 { f, _ := strconv.ParseFloat(rnS(), 64); return f }

func rsI(l, r int) []int {
	t := make([]int, r)
	for i := l; i < r; i++ {
		t[i] = rnI()
	}
	return t
}

/* ===========================PART 2: Math Func ============================  */
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func memset(a []int, v int) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}
