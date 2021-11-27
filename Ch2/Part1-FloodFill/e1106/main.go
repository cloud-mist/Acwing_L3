package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 1010

type p struct {
	x, y int
}

var (
	n  int
	a  [N][N]int
	st [N][N]bool
	q  [N * N]p
)

func main() {
	defer ot.Flush()
	n = rnI()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = rnI()
		}
	}

	res1, res2 := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !st[i][j] {
				higher, lower := bfs(i, j)
				if !higher {
					res1++
				}
				if !lower {
					res2++
				}
			}
		}
	}
	fmt.Println(res1, res2)
}

func bfs(sx, sy int) (higher, lower bool) {
	hh, tt := 0, -1
	tt++
	q[tt] = p{sx, sy}
	st[sx][sy] = true

	for hh <= tt {
		t := q[hh]
		hh++
		for i := t.x - 1; i <= t.x+1; i++ {
			for j := t.y - 1; j <= t.y+1; j++ {
				if i == t.x && j == t.y {
					continue
				}
				if i < 0 || i >= n || j < 0 || j >= n {
					continue
				}

				if a[i][j] > a[t.x][t.y] {
					higher = true
				} else if a[i][j] == a[t.x][t.y] && !st[i][j] {
					tt++
					q[tt] = p{i, j}
					st[i][j] = true
				} else if a[i][j] < a[t.x][t.y] {
					lower = true
				}
			}
		}
	}
	return
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
