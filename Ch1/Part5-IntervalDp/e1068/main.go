package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	N   = 210
	INF = int(1e8)
)

var (
	a, s [N]int
	f, g [N][N]int
)

func main() {
	defer ot.Flush()

	// Input & Init f,g & Prefix sum
	n := rnI()
	initArray(&f, INF)
	initArray(&g, -INF)

	for i := 1; i <= n; i++ {
		a[i] = rnI()
		a[i+n] = a[i]
	}

	for i := 1; i <= 2*n; i++ {
		f[i][i], g[i][i] = 0, 0
		s[i] = s[i-1] + a[i]
	}

	// 状态--> 左端点--> 右端点--> 策略
	for len := 2; len <= n; len++ {
		for l := 1; l <= n*2-len+1; l++ {
			r := l + len - 1
			for k := l; k < r; k++ {
				f[l][r] = min(f[l][r], f[l][k]+f[k+1][r]+s[r]-s[l-1])
				g[l][r] = max(g[l][r], g[l][k]+g[k+1][r]+s[r]-s[l-1])
			}
		}
	}

	minV, maxV := INF, -INF
	for i := 1; i <= n; i++ {
		minV = min(minV, f[i][n+i-1])
		maxV = max(maxV, g[i][n+i-1])
	}

	fmt.Printf("%d\n%d", minV, maxV)
}

func initArray(a *[N][N]int, x int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			a[i][j] = x
		}
	}
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
