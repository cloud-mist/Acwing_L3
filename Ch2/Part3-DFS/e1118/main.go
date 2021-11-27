package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 10

var (
	n     int
	p     []int
	st    [N]bool
	group [N][N]int
	ans   = N // 最坏情况
)

func main() {
	defer ot.Flush()

	n = rnI()
	p = rsI(0, n)

	dfs(1, 0, 0, 0)

	fmt.Println(ans)
}

func dfs(g, gc, tc, start int) {
	if g >= ans { // 如果当前答案g 已经比答案大了，说明直接return
		return
	}
	if tc == n { // 如果已经搜了n个元素，说明已经有答案了
		ans = g
	}

	flg := true
	for i := start; i < n; i++ {
		// 如果合法且和当前组的所有元素互质
		if !st[i] && valid(group[g], gc, i) {
			st[i] = true
			group[g][gc] = i
			dfs(g, gc+1, tc+1, i+1)
			st[i] = false

			flg = false
		}

		if g == 1 && start == 0 {
			return
		}
	}

	if flg {
		dfs(g+1, 0, tc, 0)
	}
}

// 判断一个组里的所有数，和当前这个数是不是都互质
func valid(grp [N]int, gc, i int) bool {
	for j := 0; j < gc; j++ {
		if gcd(p[grp[j]], p[i]) > 1 {
			return false
		}
	}
	return true
}

// 最大公约数
func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
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
