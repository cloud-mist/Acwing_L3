package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 160

type p struct {
	x, y int
}

var (
	n, m int
	dist [N][N]int
	a    [N]string
	q    [N * N]p
)

func main() {
	defer ot.Flush()

	m, n = rnI(), rnI()
	for i := 0; i < n; i++ {
		a[i] = rnS()
	}

	fmt.Println(bfs())
}

func bfs() int {
	dx := [...]int{-2, -1, 1, 2, 2, 1, -1, -2}
	dy := [...]int{1, 2, 2, 1, -1, -2, -2, -1}

	// 找起点
	var sx, sy int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == 'K' {
				sx, sy = i, j
			}
		}
	}

	// 起点入队
	hh, tt := 0, -1
	tt++
	q[tt] = p{sx, sy}

	// 初始化距离
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			dist[i][j] = -1
		}
	}
	dist[sx][sy] = 0

	for hh <= tt {
		t := q[hh]
		hh++

		for i := 0; i < 8; i++ {
			// 如果不是越界，不是访问过，不是障碍物,不是终点,入队&更新距离
			x, y := t.x+dx[i], t.y+dy[i]
			if x < 0 || x >= n || y < 0 || y >= m {
				continue
			}
			if dist[x][y] != -1 {
				continue
			}
			if a[x][y] == '*' {
				continue
			}
			if a[x][y] == 'H' {
				return dist[t.x][t.y] + 1
			}

			tt++
			q[tt] = p{x, y}
			dist[x][y] = dist[t.x][t.y] + 1

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
