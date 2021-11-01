package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	X   = 25
	Y   = 80
	INF = int(1e10)
)

var f [X][Y]int

func main() {
	defer ot.Flush()

	n, m, k := rnI(), rnI(), rnI()

	// 初始化
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			f[i][j] = INF
		}
	}
	f[0][0] = 0

	// calc
	for i := 1; i <= k; i++ {
		v1, v2, w := rnI(), rnI(), rnI()
		for j := n; j >= 0; j-- {
			for k := m; k >= 0; k-- {
				f[j][k] = min(f[j][k], f[max(j-v1, 0)][max(k-v2, 0)]+w)
			}
		}
	}

	fmt.Print(f[n][m])

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
/* ============================PART1: I/O ================================== */

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
