// 友好城市

// 排序--> LIS
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type p struct {
	x, y int
}

func main() {
	defer ot.Flush()
	// input
	n := rnI()
	a := make([]p, n+1)
	f := make([]int, n+10)
	for i := 1; i <= n; i++ {
		a[i].x, a[i].y = rnI(), rnI()
	}

	// sort
	sort.Slice(a, func(i, j int) bool {
		return a[i].x < a[j].x
	})

	// calc
	res := 0
	for i := 1; i <= n; i++ {
		f[i] = 1
		for j := 1; j < i; j++ {
			if a[i].y > a[j].y {
				f[i] = max(f[i], f[j]+1)
			}
		}
		res = max(res, f[i])
	}

	// output
	fmt.Print(res)

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
