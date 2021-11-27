package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 100010

var (
	n, k int
	dist [N]int
	q    [N]int
)

func main() {
	defer ot.Flush()
	n, k = rnI(), rnI()

	fmt.Println(bfs())
}

func bfs() int {

	// 初始化距离
	for i := 0; i < N; i++ {
		dist[i] = -1
	}
	dist[n] = 0

	// 起点入队
	hh, tt := 0, -1
	tt++
	q[tt] = n

	for hh <= tt {
		t := q[hh]
		hh++

		// answer
		if t == k {
			return dist[k]
		}

		if t-1 >= 0 && dist[t-1] == -1 {
			tt++
			q[tt] = t - 1
			dist[t-1] = dist[t] + 1
		}
		if t+1 <= k && dist[t+1] == -1 {
			tt++
			q[tt] = t + 1
			dist[t+1] = dist[t] + 1
		}
		// 只有k为偶数的情况下或者k为奇数
		if t*2-k%2 <= k && dist[t*2] == -1 {
			tt++
			q[tt] = t * 2
			dist[t*2] = dist[t] + 1
		}
	}

	return -1
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
