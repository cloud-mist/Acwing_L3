package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 1010

type p struct {
	x, y int
}

var (
	n   int
	a   [N][N]int
	q   [N * N]p
	pre [N][N]p
)

func main() {
	defer ot.Flush()

	//input
	n = rnI()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = rnI()
			pre[i][j] = p{-1, -1}
		}
	}

	bfs(n-1, n-1) // 从出口搜到入口,这样保存信息的数组，就可以正着遍历答案

	end := p{0, 0}
	for {
		fmt.Fprintln(ot, end.x, end.y)
		if end.x == n-1 && end.y == n-1 {
			break
		}
		end = pre[end.x][end.y]
	}
}

func bfs(sx, sy int) {
	dx, dy := [...]int{-1, 0, 1, 0}, [...]int{0, -1, 0, 1}
	hh, tt := 0, -1
	tt++
	q[tt] = p{sx, sy}

	for hh <= tt {
		t := q[hh]
		hh++

		for i := 0; i < 4; i++ {
			// 如果不是墙 不是越界 没有被遍历过, 那么加入队列,标记前驱
			x, y := t.x+dx[i], t.y+dy[i]
			if x < 0 || x >= n || y < 0 || y >= n {
				continue
			}
			if a[x][y] == 1 {
				continue
			}
			if pre[x][y].x != -1 {
				continue
			}

			tt++
			q[tt] = p{x, y}
			pre[x][y] = t
		}
	}
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
