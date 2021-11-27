// 连通块的数量
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	N = 1010
	M = N * N
)

var (
	n, m int
	st   [N][N]bool
	q    [M]p
	a    []string
)

type p struct {
	x, y int
}

func main() {
	defer ot.Flush()
	n, m = rnI(), rnI()
	a = rsS(0, n)

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 只有是W并且未被访问过才bfs
			if a[i][j] == 'W' && !st[i][j] {
				bfs(i, j)
				res++
			}
		}
	}

	fmt.Println(res)
}

func bfs(sx, sy int) {
	hh, tt := 0, -1

	// 插入当前点&标记
	tt++
	q[0] = p{sx, sy}
	st[sx][sy] = true

	for hh <= tt {
		t := q[hh] //取队头&出队
		hh++

		// 遍历点的八个方向
		for i := t.x - 1; i <= t.x+1; i++ {
			for j := t.y - 1; j <= t.y+1; j++ {
				// 剔除 当前点,不合法的点,是。的点,访问过的点
				if i == t.x && j == t.y {
					continue
				}
				if i < 0 || i >= n || j < 0 || j >= m {
					continue
				}
				if a[i][j] == '.' || st[i][j] {
					continue
				}

				// 入队&标记
				tt++
				q[tt] = p{i, j}
				st[i][j] = true
			}
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
func rsS(l, r int) []string {
	t := make([]string, r)
	for i := l; i < r; i++ {
		t[i] = rnS()
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
