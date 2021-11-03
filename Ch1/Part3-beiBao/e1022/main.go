package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 1010

var f [N][N]int

func main() {
	defer ot.Flush()
	n, m, k := rnI(), rnI()-1, rnI() // 皮卡丘体力不能等于0,留1手

	for i := 1; i <= k; i++ {
		a, b := rnI(), rnI()
		for j := n; j >= a; j-- {
			for k := m; k >= b; k-- {
				f[j][k] = max(f[j][k], f[j-a][k-b]+1)
			}
		}
	}

	res := int(1e8)
	// 从0开始循环
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if f[i][j] == f[n][m] {
				res = min(res, j)
			}
		}
	}
	fmt.Print(f[n][m], m+1-res)
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
