// 多源
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
	n, m int
	q    [N * N]p
	a    [N]string
	dist [N][N]int
)

func main() {
	defer ot.Flush()
	n, m = rnI(), rnI()
	for i := 0; i < n; i++ {
		a[i] = rnS()
	}

	bfs()

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprintf(ot, "%d ", dist[i][j])
		}
		fmt.Fprintln(ot)
	}
}

func bfs() {
	// 初始化dist
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			dist[i][j] = -1
		}
	}

	// 源点入队
	hh, tt := 0, -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == '1' {
				dist[i][j] = 0
				tt++
				q[tt] = p{i, j}
			}
		}
	}

	dx, dy := [...]int{-1, 0, 1, 0}, [...]int{0, 1, 0, -1}
	for hh <= tt {
		t := q[hh]
		hh++

		for i := 0; i < 4; i++ {
			x, y := t.x+dx[i], t.y+dy[i]
			if x < 0 || x >= n || y < 0 || y >= m {
				continue
			}
			if dist[x][y] != -1 {
				continue
			}

			tt++
			q[tt] = p{x, y}
			dist[x][y] = dist[t.x][t.y] + 1
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
