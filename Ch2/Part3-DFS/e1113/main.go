// 连通性 红与黑
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 25

var (
	g    [N]string
	st   [N][N]bool
	n, m int
)

func main() {
	defer ot.Flush()

	for m, n = rnI(), rnI(); n != 0 && m != 0; m, n = rnI(), rnI() {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				st[i][j] = false
			}
		}

		for i := 0; i < n; i++ {
			g[i] = rnS()
		}

		var x, y int
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if g[i][j] == '@' {
					x, y = i, j
				}
			}
		}

		fmt.Fprintln(ot, bfs(x, y))
	}

}

func bfs(x, y int) int {
	st[x][y] = true
	cnt := 1

	dx, dy := [4]int{-1, 0, 1, 0}, [4]int{0, 1, 0, -1}

	for i := 0; i < 4; i++ {
		a, b := x+dx[i], y+dy[i]
		if a < 0 || a >= n || b < 0 || b >= m {
			continue
		}
		if g[a][b] != '.' {
			continue
		}
		if st[a][b] {
			continue
		}
		cnt += bfs(a, b)
	}

	return cnt
}

/* ======================================================================== */
//
//
//
//						 _____   _   _   ____
//						| ____| | \ | | |  _ \
//						|  _|   |  \| | | | | |
//						| |___  | |\  | | |_| |
//						|_____| |_| \_| |____/
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
