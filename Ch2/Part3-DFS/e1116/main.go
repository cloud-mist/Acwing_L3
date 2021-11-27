package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 10

var (
	n, m int
	ans  int
	st   [N][N]bool
	dx   = [...]int{-2, -2, -1, 1, 2, 2, 1, -1}
	dy   = [...]int{-1, 1, 2, 2, 1, -1, -2, -2}
)

func main() {
	defer ot.Flush()
	T := rnI()

	for ; T != 0; T-- {
		ans = 0
		n, m = rnI(), rnI()
		x, y := rnI(), rnI()

		dfs(x, y, 1)
		fmt.Fprintln(ot, ans)
	}
}

func dfs(x, y, cnt int) {
	st[x][y] = true

	if cnt == n*m {
		ans++
	}

	for i := 0; i < 8; i++ {
		a, b := x+dx[i], y+dy[i]
		if a < 0 || a >= n || b < 0 || b >= m {
			continue
		}
		if st[a][b] {
			continue
		}

		dfs(a, b, cnt+1)
	}

	st[x][y] = false // 恢复现场
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
