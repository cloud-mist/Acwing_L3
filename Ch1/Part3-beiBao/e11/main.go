package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	MOD = int(1e9) + 7
	INF = int(1e10)
)

func main() {
	defer ot.Flush()

	n, m := rnI(), rnI()
	f, g := make([]int, m+10), make([]int, m+10)
	memset(f, -INF)
	f[0] = 0
	g[0] = 1

	for i := 1; i <= n; i++ {
		v, w := rnI(), rnI()

		for j := m; j >= v; j-- {
			maxn := max(f[j], f[j-v]+w)
			cnt := 0
			if maxn == f[j] {
				cnt += g[j]
			}
			if maxn == f[j-v]+w {
				cnt += g[j-v]
			}
			g[j] = cnt % MOD
			f[j] = maxn
		}
	}

	// 注意最大值不一定是f[m]
	maxx := 0
	for i := 1; i <= m; i++ {
		maxx = max(maxx, f[i])
	}

	res := 0
	for i := 1; i <= m; i++ {
		if f[i] == maxx {
			res += g[i]
		}
	}

	fmt.Print(res % MOD)
}

/* ======================================================================== */
//
//
//
//							 _____   _   _   ____
//							| ____| | \ | | |  _ \
//							|  _|   |  \| | | | | |
//							| |___  | |\  | | |_| |
//							|_____| |_| \_| |____/
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
