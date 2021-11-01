package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 1010

var (
	v, w [N]int
	f    [N][N]int
)

func main() {
	defer ot.Flush()

	// input
	n, m := rnI(), rnI()
	for i := 1; i <= n; i++ {
		v[i], w[i] = rnI(), rnI()
	}

	// dp
	for i := n; i >= 1; i-- {
		for j := 0; j <= m; j++ {
			f[i][j] = max(f[i][j], f[i+1][j])
			if j >= v[i] {
				f[i][j] = max(f[i][j], f[i+1][j-v[i]]+w[i])
			}
		}
	}

	// 方案数 可以选的话，就必选,因为字典序原因
	j := m
	for i := 1; i <= n; i++ {
		if j >= v[i] && f[i][j] == f[i+1][j-v[i]]+w[i] {
			fmt.Printf("%d ", i)
			j -= v[i]
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
