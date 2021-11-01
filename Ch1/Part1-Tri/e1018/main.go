package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	N   = 110
	INF = int(1e9)
)

var w, f [N][N]int

func main() {
	defer ot.Flush()

	// Input
	n := rnI()
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			w[i][j] = rnI()
		}
	}

	// calc
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				f[i][j] = w[i][j]
			} else {
				f[i][j] = INF
				if i > 1 {
					f[i][j] = min(f[i][j], f[i-1][j]+w[i][j])
				}
				if j > 1 {
					f[i][j] = min(f[i][j], f[i][j-1]+w[i][j])
				}
			}
		}
	}

	// output
	fmt.Println(f[n][n])
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
