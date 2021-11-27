// 玉米田
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	N   = 14
	M   = 1 << 12
	MOD = int(1e8)
)

var (
	state []int
	n, m  int
	f     [N][M]int
	head  [M][]int
	g     [N]int
)

func main() {
	defer ot.Flush()
	n, m = rnI(), rnI()

	// 一排压缩为一个二进制数,且 1-->0, 0-->1
	for i := 1; i <= n; i++ {
		for j := m - 1; j >= 0; j-- {
			tmp := rnI()
			g[i] += (1 ^ tmp) << j
		}
	}

	// 预处理合法情况 存储
	for i := 0; i < 1<<m; i++ {
		if valid(i) {
			state = append(state, i)
		}
	}

	// 合理的转移方案 存储
	for i, a := range state {
		for j, b := range state {
			if a&b == 0 {
				head[i] = append(head[i], j)
			}
		}
	}

	// dp
	f[0][0] = 1
	for i := 1; i <= n+1; i++ {
		for a := range state {
			for _, b := range head[a] {
				if g[i]&state[a] == 0 {
					f[i][a] += f[i-1][b]
				}
			}
		}
	}

	fmt.Println(f[n+1][0] % MOD)
}

func valid(x int) bool {
	for i := 0; i < m; i++ {
		if (x>>i&1)&(x>>(i+1)&1) == 1 {
			return false
		}
	}

	return true
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
