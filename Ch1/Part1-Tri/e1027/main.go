package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const N = 20

var (
	w [N][N]int
	f [N * 2][N][N]int
)

func main() {
	defer ot.Flush()

	n := rnI()
	for a, b, v := rnI(), rnI(), rnI(); a != 0 && b != 0 && v != 0; a, b, v = rnI(), rnI(), rnI() {
		w[a][b] = v
	}

	for k := 2; k <= n*2; k++ {
		for i1 := 1; i1 <= n; i1++ {
			for i2 := 1; i2 <= n; i2++ {
				j1, j2 := k-i1, k-i2

				if j1 >= 1 && j1 <= n && j2 >= 1 && j2 <= n {
					// 如果没有重复
					t := w[i1][j1]
					if i1 != i2 {
						t += w[i2][j2]
					}

					// 四种前面的状态
					x := &f[k][i1][i2]
					*x = max(*x, t+f[k-1][i1-1][i2-1])
					*x = max(*x, t+f[k-1][i1-1][i2])
					*x = max(*x, t+f[k-1][i1][i2-1])
					*x = max(*x, t+f[k-1][i1][i2])
				}
			}
		}
	}

	fmt.Println(f[n*2][n][n])
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

/* ==== G0 ===== */
const BUFSIZE = int(1e6)

var (
	ot = bufio.NewWriterSize(os.Stdout, BUFSIZE)
	in = newScanner(os.Stdin)
)

type scanner struct{ sc *bufio.Scanner }

func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e9))
	return &scanner{sc}
}

/* ==== G1 ===== */
func rnS() string  { in.sc.Scan(); return in.sc.Text() }
func rnI() int     { i, _ := strconv.Atoi(rnS()); return i }
func rnF() float64 { f, _ := strconv.ParseFloat(rnS(), 64); return f }

/* ==== G2 ===== */
func rsI(n int) []int {
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = rnI()
	}
	return t
}
func rsF(n int) []float64 {
	t := make([]float64, n)
	for i := 0; i < n; i++ {
		t[i] = rnF()
	}
	return t
}
func rsS(n int) []string {
	t := make([]string, n)
	for i := 0; i < n; i++ {
		t[i] = rnS()
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
