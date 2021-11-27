// 连通性 迷宫
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 110

var (
	n              int
	xa, ya, xb, yb int
	g              [N]string
	st             [N][N]bool
)

func main() {
	defer ot.Flush()
	T := rnI()

	for ; T != 0; T-- {
		// 初始化st
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				st[i][j] = false
			}
		}

		// input
		n = rnI()
		for i := 0; i < n; i++ {
			g[i] = rnS()
		}
		xa, ya, xb, yb = rnI(), rnI(), rnI(), rnI()

		// dfs
		if dfs(xa, ya) {
			fmt.Fprintln(ot, "YES")
		} else {
			fmt.Fprintln(ot, "NO")
		}
	}
}

func dfs(x, y int) bool {
	st[x][y] = true // 标记x,y 已经访问过
	dx, dy := [4]int{-1, 0, 1, 0}, [4]int{0, 1, 0, -1}

	if g[x][y] == '#' { // 如果是障碍，是不能走到的
		return false
	}
	if x == xb && y == yb { // 如果是答案，说明可以走到
		return true
	}

	for i := 0; i < 4; i++ {
		// 遍历4个方向
		a, b := x+dx[i], y+dy[i]
		if a < 0 || a >= n || b < 0 || b >= n {
			continue
		}
		if st[a][b] {
			continue
		}

		if dfs(a, b) {
			return true
		}
	}

	return false
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
