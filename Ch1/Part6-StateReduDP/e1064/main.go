// 小国王
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	N = 12
	M = 110
	K = 1 << 10
)

var (
	n, m  int
	state []int
	cnt   [K]int
	f     [N][M][K]int
	head  [K][]int
)

func main() {
	defer ot.Flush()

	n, m = rnI(), rnI()

	// 1. 合法状态筛选,存的是数（不存在相邻1的情况）
	for i := 0; i < (1 << n); i++ {
		if valid(i) {
			state = append(state, i)
			cnt[i] = count(i) // 状态i的1的个数
		}
	}

	// 2. 状态转移合法的记录
	// 如果满足两个条件，就把下标为j的状态可以在下标为i的状态前面存入
	for i, a := range state {
		for j, b := range state {
			if (a&b) == 0 && valid(a|b) {
				head[i] = append(head[i], j)
			}
		}
	}

	f[0][0][0] = 1
	for i := 1; i <= n+1; i++ { // n+1为了后面省一个循环答案
		for j := 0; j <= m; j++ { //国王个数
			for a := range state { // 合法状态
				for _, b := range head[a] {
					c := cnt[state[a]]
					if j >= c { // 要预留c个，才能转移，否则不合法
						f[i][j][a] += f[i-1][j-c][b]
					}
				}
			}
		}
	}

	fmt.Println(f[n+1][m][0])
}

func valid(x int) bool {
	for i := 0; i < n; i++ {
		if ((x >> i & 1) & (x >> (i + 1) & 1)) == 1 {
			return false
		}
	}
	return true
}

// 求x的1的个数
func count(x int) int {
	res := 0
	for i := 0; i < n; i++ {
		res += (x >> i & 1)
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
