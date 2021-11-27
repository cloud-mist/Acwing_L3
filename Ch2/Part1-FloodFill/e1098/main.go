package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	N = 60
	M = N * N
)

type p struct {
	x, y int
}

var (
	n, m int
	q    [M]p
	a    [N][N]int
	st   [N][N]bool
)

func main() {
	defer ot.Flush()

	n, m = rnI(), rnI()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			a[i][j] = rnI()
		}
	}

	var cnt, area int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !st[i][j] {
				area = max(area, bfs(i, j))
				cnt++
			}
		}
	}

	fmt.Printf("%d\n%d", cnt, area)
}

func bfs(sx, sy int) int {
	hh, tt := 0, -1
	res := 0
	dx := [...]int{0, -1, 0, 1}
	dy := [...]int{-1, 0, 1, 0}

	tt++
	q[tt] = p{sx, sy}
	st[sx][sy] = true

	for hh <= tt {
		// 取队头&出队&res++
		t := q[hh]
		hh++
		res++

		// 如果四个方向没有越界且没有墙且没有访问过，加入队列
		for i := 0; i < 4; i++ {
			x, y := t.x+dx[i], t.y+dy[i]
			if x < 0 || x >= n || y < 0 || y >= m || st[x][y] {
				continue
			}
			if a[t.x][t.y]>>i&1 == 1 {
				continue
			}
			tt++
			q[tt] = p{x, y}
			st[x][y] = true
		}
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
