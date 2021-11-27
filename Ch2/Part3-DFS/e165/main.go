// 剪枝 小猫下山
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const N = 20

var (
	w    []int
	sum  [N]int
	n, m int
	ans  = N
)

func main() {
	defer ot.Flush()

	n, m = rnI(), rnI()
	w = rsI(0, n)
	// 顺序优化
	sort.Slice(w, func(i, j int) bool {
		return w[i] > w[j]
	})

	dfs(0, 0) // 第0只猫，当前车辆0

	fmt.Println(ans)
}

func dfs(u, k int) {

	// 车辆数大于ans,直接return
	if k >= ans {
		return
	}

	// 答案
	if u == n {
		ans = k
		return
	}

	for i := 0; i < k; i++ {
		// 如果可以放当前车,加上，递归，恢复现场
		if sum[i]+w[u] <= m {
			sum[i] += w[u]
			dfs(u+1, k)
			sum[i] -= w[u]
		}
	}

	// 放不到，新车
	sum[k] = w[u]
	dfs(u+1, k+1)
	//	sum[k] = 0
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
